package xlog

import (
	"C"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"strings"
	_ "unsafe"
)

type Severity int32
type XLog struct {level Severity; depth int}

const (
	InfoLog Severity = iota
	WarningLog
	ErrorLog
	FatalLog
	NumSeverity
	SInfoLog = "info"
	SWarningLog = "warning"
	SErrorLog = "error"
	SFatalLog = "fatal"
)

var logLevel = int32(InfoLog)

func init() {
	glog.MaxSize = 1024 * 1024 * 1024
}

func Logger(level Severity, depth int) *XLog {
	return &XLog{level: level, depth: depth}
}

func (l *XLog)Write(p []byte) (n int, err error) {
	if len(p) > 0 {l.Printf("%s", string(p))}
	return len(p), nil
}

func (l *XLog)Printf(s string, i ...interface{}) {
	Depth(l.level, s, l.depth, i...)
}

func SetLogDir(dir string) {
	if dir != "" {
		os.MkdirAll(dir, os.ModePerm)
		flag.Set("log_dir", dir)
	}

	flag.Set("alsologtostderr", "true")
	flag.Parse()
}

func SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case SInfoLog: logLevel = int32(InfoLog)
	case SWarningLog: logLevel = int32(WarningLog)
	case SErrorLog: logLevel = int32(ErrorLog)
	case SFatalLog: logLevel = int32(FatalLog)
	}
}

func Flush() {
	glog.Flush()
}

func Info(format string, args ...interface{}) {
	if logLevel <= int32(InfoLog) {
		glog.InfoDepth(1, fmt.Sprintf(format, args...))
	}
}

func Warning(format string, args ...interface{}) {
	if logLevel <= int32(WarningLog) {
		glog.WarningDepth(1, fmt.Sprintf(format, args...))
	}
}

func Error(format string, args ...interface{}) {
	if logLevel <= int32(ErrorLog) {
		glog.ErrorDepth(1, fmt.Sprintf(format, args...))
	}
}

func Panic(info interface{}) {
	glog.FatalDepth(1, info)
}

func InfoDepth(format string, depth int, args ...interface{}) {
	if logLevel <= int32(InfoLog) {
		glog.InfoDepth(depth + 1, fmt.Sprintf(format, args...))
	}
}

func WarningDepth(format string, depth int, args ...interface{}) {
	if logLevel <= int32(WarningLog) {
		glog.WarningDepth(depth + 1, fmt.Sprintf(format, args...))
	}
}

func ErrorDepth(format string, depth int, args ...interface{}) {
	if logLevel <= int32(ErrorLog) {
		glog.ErrorDepth(depth + 1, fmt.Sprintf(format, args...))
	}
}

func PanicDepth(format string, depth int, args ...interface{}) {
	glog.FatalDepth(depth + 1, fmt.Sprintf(format, args...))
}

func Depth(level Severity, format string, depth int, args ...interface{}) {
	switch level {
	case WarningLog: WarningDepth(format, depth, args...)
	case ErrorLog: ErrorDepth(format, depth, args...)
	case FatalLog: PanicDepth(format, depth, args...)
	default: InfoDepth(format, depth, args...)
	}
}

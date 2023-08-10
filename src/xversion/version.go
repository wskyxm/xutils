package xversion

import (
	"fmt"
	"os"
	"runtime"
)

var (
	version      string
	gitBranch    string
	gitTag       string
	gitCommit    string
	gitTreeState string
	buildDate    string
)

type Info struct {
	Version      string `json:"version"`
	GitBranch    string `json:"gitBranch"`
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	return info.GitCommit
}

func GetVersion() Info {
	return Info{
		Version:      version,
		GitBranch:    gitBranch,
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func ShowVersion() bool {
	// 参数检查
	if len(os.Args) < 2 {return false}
	if os.Args[1] != "version" {return false}

	// 显示版本信息
	v := GetVersion()
	f := "Version: %s\nGitBranch: %s\nCommitId: %s\nBuild Date: %s\nGo Version: %s\nOS/Arch: %s\n"
	fmt.Printf(f, v.Version, v.GitBranch, v.GitCommit, v.BuildDate, v.GoVersion, v.Platform)
	return true
}

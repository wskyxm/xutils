package xversion

import (
	"os"
	"strings"
)

const makefileContent = `
BINARY_NAME="${NAME}"
TARGET_FILE="${MAINFILE}"
VERSION="${VERSION}"
TARGET_ARCH="amd64"

# build with verison infos
versionDir="github.com/wskyxm/xutils/xversion"
gitTag=$(shell if [ "${BACKQUOTE}git describe --tags --abbrev=0 2>/dev/null${BACKQUOTE}" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(shell git rev-parse --short HEAD)
gitTreeState=$(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

# init ld flags
ldflags="-s -w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState} -X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch}"

# init default target os and arch
UNAME := $(shell uname)
ifeq ($(UNAME), Linux)
	BINARY=$(BINARY_NAME)
	TARGET_OS="linux"
else
	BINARY=$(BINARY_NAME).exe
	TARGET_OS="windows"
endif

# set target arch
ifneq ($(arch),)
    TARGET_ARCH=$(arch)
endif

# set target os
ifneq ($(os),)
    TARGET_OS=$(os)
endif

help:
	@echo "make os=[linux, windows] arch=[amd64, arm64]"

default:
	@echo "build the ${BINARY_NAME}, ${TARGET_OS}, ${TARGET_ARCH}"
	@GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} go build -ldflags ${ldflags} -o bin/${BINARY} ${TARGET_FILE}
	@echo "build done."
`

func GenerateMakefile(name, mainfile, version, path string) {
	// 初始化Makefile内容
	text := strings.ReplaceAll(makefileContent, "${NAME}", name)
	text = strings.ReplaceAll(text, "${MAINFILE}", mainfile)
	text = strings.ReplaceAll(text, "${VERSION}", version)

	// 保存文件
	os.WriteFile(path, []byte(text), os.ModePerm)
}

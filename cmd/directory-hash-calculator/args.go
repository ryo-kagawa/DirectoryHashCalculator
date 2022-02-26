package main

import (
	"os"
	"strings"

	"github.com/ryo-kagawa/go-utils/conditional"
)

var usage = `
Usage: DirectoryHash [--hashAlgorithm] ["targetDirectory"]
--hashAlgorithm string
    ハッシュアルゴリズムを指定
    対応しているアルゴリズム
    MD5 (default)
targetDirectoryに指定がない場合はカレントディレクトリで実行します
Usage: DirectoryHash --help
    ここを表示します
Usage: DirectoryHash --version
    バージョンを表示します
`

type Args struct {
	help            bool
	version         bool
	hashAlgorithm   string
	targetDirectory string
}

func GetArgs() Args {
	result := Args{}
	args := os.Args[1:]
	if 0 == len(args) {
		return result
	}
	switch args[0] {
	case "--help":
		return Args{
			help: true,
		}
	case "--version":
		return Args{
			version: true,
		}
	}
	hashAlgorithm := conditional.StringFunc(
		1 <= len(args) && strings.HasPrefix(args[0], "--hashAlgorithm="),
		func() string {
			arg := strings.Trim(strings.Split(args[0], "--hashAlgorithm=")[1], "\"")
			args = args[1:]
			return arg
		},
		func() string {
			return "MD5"
		},
	)

	targetDirectory := conditional.StringFunc(
		1 <= len(args),
		func() string {
			return args[0]
		},
		func() string {
			return "."
		},
	)

	return Args{
		hashAlgorithm:   hashAlgorithm,
		targetDirectory: targetDirectory,
	}
}

func (_self Args) IsHelp() bool {
	return _self.help
}

func (_self Args) IsVersion() bool {
	return _self.version
}

func (_self Args) GetHashAlgorithm() string {
	return _self.hashAlgorithm
}

func (_self Args) GetTargetDirectory() string {
	return _self.targetDirectory
}

func (_self Args) GetUsage() string {
	return strings.Trim(usage, "\n")
}

func (_self Args) GetVersion() string {
	return GetVersion()
}

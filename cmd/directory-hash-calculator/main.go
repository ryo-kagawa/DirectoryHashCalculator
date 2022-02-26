package main

import (
	"fmt"

	"github.com/ryo-kagawa/DirectoryHashCalculator/model"
)

func main() {
	result, err := run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func run() (string, error) {
	args := GetArgs()
	if args.IsHelp() {
		return args.GetUsage(), nil
	}
	if args.IsVersion() {
		return args.GetVersion(), nil
	}

	hashAlgorithmCreater, err := model.NewHashAlgorithmCreater(args.GetHashAlgorithm())
	if err != nil {
		return "", err
	}

	directory := model.NewDirectory(args.GetTargetDirectory())
	hash, err := directory.CalculateHash(hashAlgorithmCreater)
	if err != nil {
		return "", err
	}
	return hash.ToHEX(), nil
}

package model

import (
	"io/fs"
	"path/filepath"
)

type Directory struct {
	path string
}

func NewDirectory(path string) Directory {
	return Directory{
		path: path,
	}
}

func (_self Directory) CalculateHash(hashAlgorithmCreater HashAlgorithmCreater) (Binary, error) {
	targetList := TargetList{}

	err := filepath.WalkDir(
		_self.path,
		func(
			path string,
			d fs.DirEntry,
			err error,
		) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				if _self.path != path {
					targetList = append(targetList, NewDirectory(path))
					return filepath.SkipDir
				}
				return nil
			}

			targetList = append(targetList, NewFile(path))
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return targetList.CalculateHash(hashAlgorithmCreater)
}

func (_self Directory) GetName() string {
	return filepath.Base(_self.path)
}

func (_self Directory) IsDirectory() bool {
	return true
}

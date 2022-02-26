package model

import (
	"io"
	"os"
	"path/filepath"
)

type File struct {
	path string
}

func NewFile(path string) File {
	return File{
		path: path,
	}
}

func (_self File) CalculateHash(hashAlgorithmCreater HashAlgorithmCreater) (Binary, error) {
	file, err := os.Open(_self.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hashFunction := hashAlgorithmCreater()
	_, err = io.Copy(hashFunction, file)
	hashValue := hashFunction.Sum(nil)
	return hashValue, nil
}

func (_self File) GetName() string {
	return filepath.Base(_self.path)
}

func (_self File) IsDirectory() bool {
	return false
}

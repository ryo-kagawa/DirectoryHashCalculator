package model

import (
	"crypto/md5"
	"errors"
	"hash"
)

type HashAlgorithmCreater func() hash.Hash

func NewHashAlgorithmCreater(hashAlgorithm string) (HashAlgorithmCreater, error) {
	switch hashAlgorithm {
	case "MD5":
		return func() hash.Hash {
			return md5.New()
		}, nil
	}
	return nil, errors.New("不正なハッシュアルゴリズムが指定されました")
}

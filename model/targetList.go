package model

import (
	"sort"
)

type TargetList []Target

func (_self TargetList) Len() int {
	return len(_self)
}

func (_self TargetList) Swap(i, j int) {
	_self[i], _self[j] = _self[j], _self[i]
}

func (_self TargetList) Less(i, j int) bool {
	return _self[i].GetName() < _self[j].GetName()
}

func (_self TargetList) CalculateHash(hashAlgorithmCreater HashAlgorithmCreater) (Binary, error) {
	sort.Sort(_self)
	hashFunction := hashAlgorithmCreater()
	for i, x := range _self {
		if 0 < i {
			hashFunction.Write([]byte{0x00})
		}
		if x.IsDirectory() {
			hashFunction.Write([]byte{0x00})
		} else {
			hashFunction.Write([]byte{0x01})
		}
		hashFunction.Write([]byte(x.GetName()))
		hash, err := x.CalculateHash(hashAlgorithmCreater)
		if err != nil {
			return nil, err
		}
		hashFunction.Write(hash)
	}
	hashValue := hashFunction.Sum(nil)
	return hashValue, nil
}

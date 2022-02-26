package model

type Target interface {
	CalculateHash(hashAlgorithmCreater HashAlgorithmCreater) (Binary, error)
	GetName() string
	IsDirectory() bool
}

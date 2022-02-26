package model

import "encoding/hex"

type Binary []byte

func (_self Binary) ToHEX() string {
	return hex.EncodeToString(_self)
}

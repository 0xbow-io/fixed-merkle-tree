package fMerkleTree

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"

	"github.com/0xbow-io/go-iden3-crypto/poseidon"
)

func SHA256Hash(left Element, right Element) []byte {
	hash := sha256.New()
	sRight := hex.EncodeToString(right)
	sLeft := hex.EncodeToString(left)
	if sRight[0] == '0' && len(sRight) == 2 {
		sRight = sRight[1:]
	}
	if sLeft[0] == '0' && len(sLeft) == 2 {
		sLeft = sLeft[1:]
	}
	hash.Write([]byte(sLeft))
	hash.Write([]byte(sRight))
	//fmt.Printf("left=%s  right=%s val=%s\n", sLeft, sRight, sLeft+sRight)
	return hash.Sum(nil)
}

func Poseidon(left Element, right Element) []byte {
	result, err := poseidon.Hash([]*big.Int{
		big.NewInt(0).SetBytes(left),
		big.NewInt(0).SetBytes(right),
	})

	if err != nil {
		panic(err.Error())
	}
	return result.Bytes()
}

func Poseidon2(left Element, right Element) []byte {
	result, err := poseidon.Poseidon2([]*big.Int{
		big.NewInt(0).SetBytes(left),
		big.NewInt(0).SetBytes(right),
	})

	if err != nil {
		panic(err.Error())
	}
	return result.Bytes()
}

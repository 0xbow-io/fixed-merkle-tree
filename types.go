package fMerkleTree

import (
	"bytes"
	"encoding/hex"
	"math/big"
)

type Element []byte

func (e Element) Hex() string {
	return hex.EncodeToString(e)
}

func (e Element) BigInt() *big.Int {
	return big.NewInt(0).SetBytes(e)
}

func (e Element) Cmp(x Element) bool {
	return bytes.Compare(e, x) == 0
}

type HashFunction func(left Element, right Element) []byte

type ComparatorFunction func(left Element, right Element) bool

type SerializedTreeState interface {
	GetLevels() int
	GetRoot() Element
	GetLayers() ([][]Element, error)
	GetZeros() ([]Element, error)
}

type serializedTreeState struct {
	Root   Element `db:"root"`
	Levels int     `db:"levels"`
	Layers []byte  `db:"layers"`
	Zeros  []byte  `db:"zeros"`
	ID     int     `db:"id"`
}

func (st *serializedTreeState) GetRoot() Element {
	return st.Root
}

func (st *serializedTreeState) GetLevels() int {
	return st.Levels
}

func (st *serializedTreeState) GetLayers() ([][]Element, error) {
	var out [][]Element
	return out, GobDecode(st.Layers, &out)
}

func (st *serializedTreeState) GetZeros() ([]Element, error) {
	var out []Element
	return out, GobDecode(st.Zeros, &out)
}

func NewSerializedTreeState(tree *MerkleTree) (SerializedTreeState, error) {
	out := &serializedTreeState{Levels: tree.levels, Root: tree.Root()}
	var err error
	out.Layers, err = GobEncode(tree.layers)
	if err != nil {
		return out, err
	}
	out.Zeros, err = GobEncode(tree.zeros)
	if err != nil {
		return out, err
	}
	return out, nil
}

type ProofPath struct {
	PathElements  []Element `json:"pathElements"`
	PathIndices   []int     `json:"pathIndices"`
	PathPositions []int     `json:"pathPositions"`
	PathRoot      Element   `json:"pathRoot"`
}

type TreeEdge struct {
	EdgeElement       Element   `json:"edgeElement"`
	EdgePath          ProofPath `json:"edgePath"`
	EdgeIndex         int       `json:"edgeIndex"`
	EdgeElementsCount int       `json:"edgeElementsCount"`
}

type TreeSlice struct {
	Edge     TreeEdge  `json:"edge"`
	Elements []Element `json:"elements"`
}

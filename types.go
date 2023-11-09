package fMerkleTree

import "encoding/hex"

type Element []byte

func (e Element) Hex() string {
	return hex.EncodeToString(e)
}

type HashFunction func(left Element, right Element) []byte

type ComparatorFunction func(left Element, right Element) bool

type SerializedTreeState struct {
	Levels int    `db:"levels"`
	Layers []byte `db:"layers"`
	Zeros  []byte `db:"zeros"`
	ID     int    `db:"id"`
}

func (st SerializedTreeState) GetLayers() ([][]Element, error) {
	var out [][]Element
	return out, GobDecode(st.Layers, &out)
}

func (st SerializedTreeState) GetZeros() ([]Element, error) {
	var out []Element
	return out, GobDecode(st.Zeros, &out)
}

func NewSerializedTreeState(tree *MerkleTree) (SerializedTreeState, error) {
	var out SerializedTreeState
	out.Levels = tree.levels
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

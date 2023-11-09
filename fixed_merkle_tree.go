package fMerkleTree

import (
	"fmt"
)

type MerkleTree struct {
	*BaseTree
}

func NewMerkleTree(levels int, elements []Element, zeroElement Element, hashFn HashFunction) (*MerkleTree, error) {
	base := &BaseTree{levels: levels}
	if len(elements) > base.Capacity() {
		return nil, fmt.Errorf("tree is full")
	}

	if hashFn == nil {
		return nil, fmt.Errorf("hash function is nil")
	}

	base.hashFn = hashFn
	base.zeroElement = zeroElement
	base.layers = make([][]Element, levels+1)
	base.layers[0] = elements
	out := &MerkleTree{base}
	out.buildZeros()
	out.buildHashes()
	return out, nil
}

func (mt *MerkleTree) buildHashes() {
	for layerIndex := 1; layerIndex <= mt.levels; layerIndex++ {
		nodes := mt.layers[layerIndex-1]

		mt.layers[layerIndex] = mt.processNodes(nodes, layerIndex)
	}
}

/**
* Insert multiple elements into the tree.
* @param elements Elements to insert
 */
func (mt *MerkleTree) BulkInsert(elements []Element) error {
	if len(elements) == 0 {
		return nil
	}
	for _, element := range elements {
		if err := mt.Insert(element); err != nil {
			return err
		}
	}
	return nil
}

func (mt MerkleTree) IndexOf(element Element) int {
	return IndexOfElement(mt.layers[0], element, 0, nil)
}

func (mt MerkleTree) Proof(element Element) (ProofPath, error) {
	index := mt.IndexOf(element)
	return mt.Path(index)
}

func (mt MerkleTree) getTreeEdge(edgeIndex int) (TreeEdge, error) {
	if edgeIndex >= len(mt.layers[0]) {
		return TreeEdge{}, fmt.Errorf("index out of range")
	}
	edgeElement := mt.layers[0][edgeIndex]
	if edgeElement == nil {
		return TreeEdge{}, fmt.Errorf("element not found")
	}
	edgePath, err := mt.Path(edgeIndex)
	if err != nil {
		return TreeEdge{}, err
	}
	return TreeEdge{
		EdgePath:          edgePath,
		EdgeElement:       edgeElement,
		EdgeIndex:         edgeIndex,
		EdgeElementsCount: len(mt.layers[0])}, nil
}

func (mt MerkleTree) GetTreeSlices(count int) ([]TreeSlice, error) {
	length := len(mt.layers[0])
	size := length / count
	if length%count != 0 {
		size++
	}
	if size%2 != 0 {
		size++
	}
	slices := []TreeSlice{}
	for i := 0; i < length; i += size {
		edgeLeft := i
		edgeRight := i + size
		edge, err := mt.getTreeEdge(edgeLeft)
		if err != nil {
			return nil, err
		}
		slices = append(slices, TreeSlice{Edge: edge, Elements: mt.layers[0][edgeLeft:edgeRight]})
	}
	return slices, nil
}

/**
* Serialize entire tree state including intermediate layers into a plain object
* Deserializing it back will not require to recompute any hashes
* Elements are not converted to a plain type, this is responsibility of the caller
 */
func (mt MerkleTree) Serialize() (SerializedTreeState, error) {
	return NewSerializedTreeState(&mt)
}

func DeserializeMerkleTree(data SerializedTreeState, hashFn HashFunction) (*MerkleTree, error) {
	layers, err := data.GetLayers()
	if err != nil {
		fmt.Println("failed to get layers")
		return nil, err
	}
	zeros, err := data.GetZeros()
	if err != nil {
		return nil, err
	}

	out := &MerkleTree{
		BaseTree: &BaseTree{
			levels: data.Levels,
			layers: layers,
			zeros:  zeros,
			hashFn: hashFn,
		},
	}
	out.zeroElement = out.zeros[0]
	return out, nil
}

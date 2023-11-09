package fMerkleTree

import (
	"bytes"
	"fmt"
	"math"
)

type BaseTree struct {
	levels      int
	hashFn      HashFunction
	zeroElement Element
	zeros       []Element
	layers      [][]Element
}

func (bt BaseTree) Capacity() int {
	return int(math.Pow(2, float64(bt.levels)))
}

func (bt BaseTree) Layers() [][]Element {
	return bt.layers
}

func (bt BaseTree) Zeros() []Element {
	return bt.zeros
}

func (bt BaseTree) Elements() []Element {
	return bt.layers[0]
}

func (bt BaseTree) Root() Element {
	if len(bt.layers[bt.levels]) == 0 {
		return bt.zeros[bt.levels]
	}
	return bt.layers[bt.levels][0]
}

/**
* Insert new element into the tree
* @param element Element to insert
 */
func (bt *BaseTree) Insert(element Element) error {
	if len(bt.layers[0]) >= bt.Capacity() {
		return fmt.Errorf("tree is full")
	}
	return bt.Update(len(bt.layers[0]), element)
}

/*
* Insert multiple elements into the tree.
* @param {Array} elements Elements to insert
 */

func (bt *BaseTree) BulkInsert(elements []Element) error {
	if len(elements) == 0 {
		return nil
	}
	if len(bt.layers[0])+len(elements) > bt.Capacity() {
		return fmt.Errorf("tree is full")
	}

	for i := range elements {
		bt.layers[0] = append(bt.layers[0], elements[i])
		level := 0
		index := len(bt.layers[0]) - 1
		for index%2 == 1 {
			level++
			index >>= 1
			left := bt.layers[level-1][index*2]
			right := bt.layers[level-1][index*2+1]
			bt.layers[level][index] = bt.hashFn(left, right)
		}
	}
	return bt.Insert(elements[len(elements)-1])
}

func (bt *BaseTree) SetLayer(i, j int, val Element) {
	if len(bt.layers[i]) <= j {
		tmp := make([]Element, j+1)
		if bt.layers[i] != nil {
			copy(tmp, bt.layers[i])
		}

		bt.layers[i] = tmp
	}
	bt.layers[i][j] = val
}

/**
* Change an element in the tree
* @param {number} index Index of element to change
* @param element Updated element value
 */
func (bt *BaseTree) Update(index int, element Element) error {
	if index < 0 || index > len(bt.layers[0]) || index >= bt.Capacity() {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	bt.SetLayer(0, index, element)
	bt.processUpdate(index)
	return nil
}

/**
* Get merkle path to a leaf
* @param {number} index Leaf index to generate path for
* @returns {{pathElements: Object[], pathIndex: number[]}} An object containing adjacent elements and left-right index
 */
func (bt *BaseTree) Path(index int) (ProofPath, error) {
	if index < 0 || index >= len(bt.layers[0]) {
		return ProofPath{}, fmt.Errorf("index out of bounds: %d", index)

	}

	var (
		elIndex                 = index
		pathElements  []Element = make([]Element, bt.levels)
		pathIndices   []int     = make([]int, bt.levels)
		pathPositions []int     = make([]int, bt.levels)
	)

	for level := 0; level < bt.levels; level++ {
		pathIndices[level] = elIndex % 2
		leafIndex := elIndex ^ 1
		if leafIndex < len(bt.layers[level]) {
			pathElements[level] = bt.layers[level][leafIndex]
			pathPositions[level] = leafIndex
		} else {
			pathElements[level] = bt.zeros[level]
			pathPositions[level] = 0
		}
		elIndex >>= 1
	}
	return ProofPath{
		PathElements:  pathElements,
		PathIndices:   pathIndices,
		PathPositions: pathPositions,
		PathRoot:      bt.layers[bt.levels][0]}, nil
}

func (bt *BaseTree) buildZeros() {
	bt.zeros = make([]Element, bt.levels+1)
	bt.zeros[0] = bt.zeroElement
	for i := 1; i <= bt.levels; i++ {
		bt.zeros[i] = bt.hashFn(bt.zeros[i-1], bt.zeros[i-1])
	}
}

func (bt *BaseTree) processNodes(nodes []Element, layerIndex int) []Element {
	length := len(nodes)
	currentLength := int(math.Ceil(float64(length) / 2))
	currentLayer := make([]Element, currentLength)
	currentLength--
	starFrom := length - ((length % 2) ^ 1)
	j := 0
	for i := starFrom; i >= 0; i -= 2 {
		if nodes[i-1] == nil {
			break
		}
		left := nodes[i-1]
		var right Element
		if i == starFrom && length%2 == 1 {
			right = bt.zeros[layerIndex-1]
		} else {
			right = nodes[i]
		}
		currentLayer[currentLength-j] = bt.hashFn(left, right)
		j++
	}
	return currentLayer
}

func (bt *BaseTree) processUpdate(index int) {
	for level := 1; level <= bt.levels; level++ {
		index >>= 1
		left := bt.layers[level-1][index*2]
		var right Element
		if index*2+1 < len(bt.layers[level-1]) {
			right = bt.layers[level-1][index*2+1]
		} else {
			right = bt.zeros[level-1]
		}
		bt.SetLayer(level, index, bt.hashFn(left, right))
	}
}

/**
* Find an element in the tree
* @param elements elements of tree
* @param element An element to find
* @param comparator A function that checks leaf value equality
* @param fromIndex The index to start the search at. If the index is greater than or equal to the array's length, -1 is returned
* @returns {number} Index if element is found, otherwise -1
 */
func IndexOfElement(elements []Element, element Element, fromIndex int, comparator ComparatorFunction) int {
	for i, ele := range elements {
		if comparator != nil {
			if comparator(element, ele) {
				return i
			}
		} else {
			if bytes.Equal(element, ele) {
				return i
			}
		}
	}
	return -1
}

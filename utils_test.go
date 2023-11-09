package fMerkleTree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GobDecode(t *testing.T) {
	tArray := [][]Element{{{1}, {2}}, {{3}, {4}}}
	tBytes, err := GobEncode(tArray)
	fmt.Printf("t: %v\n", tBytes)
	require.NoError(t, err)
	out := [][]Element{}
	err = GobDecode(tBytes, &out)
	require.NoError(t, err)
	require.ElementsMatch(t, tArray, out)
}

package fMerkleTree

import (
	"bytes"
	"encoding/gob"
)

func GobEncode(e any) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(e)
	return buf.Bytes(), err
}

func GobDecode(data []byte, e any) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(e)
}

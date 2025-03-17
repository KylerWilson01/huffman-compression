// Package encoder is the package for handling the encoding part of the huffman compression
package encoder

import (
	"os"

	"github.com/KylerWilson01/huffman-compression/internal/header"
)

// EncodeFile encodes the given file and creates a new file
func EncodeFile(f *os.File, t []header.PrefixNode, nfn string) {
	nf := createFile(nfn)
	content := replaceCharWithBitString(f, t)

	_, err := nf.Write([]byte{})
	if err != nil {
		panic("could not write header to file")
	}

	_, err = nf.Write(content)
	if err != nil {
		panic("could not write content to file")
	}
}

func replaceCharWithBitString(f *os.File, t []header.PrefixNode) []byte {
	return nil
}

func createFile(nf string) *os.File {
	return nil
}

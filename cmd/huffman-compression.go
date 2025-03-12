package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	frequencymap "github.com/KylerWilson01/huffman-compression/internal/frequency-map"
)

var (
	fp         string
	compress   bool
	decompress bool
)

func main() {
	var err error
	var f *os.File
	var fileString string

	// parse the flags
	flag.StringVar(&fp, "file", "", "The file path that needs to be compressed or decompressed")
	flag.BoolVar(
		&compress,
		"compress",
		false,
		"True if the file needs to be compressed",
	)
	flag.BoolVar(
		&decompress,
		"decompress",
		false,
		"True if the file needs to be decompressed",
	)
	flag.Parse()

	if compress && decompress {
		err := fmt.Errorf("Cannot compress and decompress at the same time")
		panic(err)
	}

	// open the file
	if fp == "" {
		err := fmt.Errorf("No file path given")
		panic(err)
	}

	f, err = os.Open(fp)
	if err != nil {
		panic(errors.New("Could not open file"))
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err.Error())
		}

		fileString += line

	}

	fm := frequencymap.NewFrequencyMap(fileString)
	fm.FindFrequencyOfChars()

	os.Exit(0)
}

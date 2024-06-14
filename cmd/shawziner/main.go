package main

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gucio321/shawziner/pkg/mscoreNotationSchema"
	"github.com/gucio321/shawziner/pkg/shawzin"
)

func main() {
	// load file from input -f
	var filename string
	flag.StringVar(&filename, "f", "", "input file")
	flag.Parse()

	// read file
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// decopress file
	fileZip, err := zip.NewReader(bytes.NewReader(file), int64(len(file)))
	if err != nil {
		log.Fatal(err)
	}

	fileName := ""
	for _, file := range fileZip.File {
		name := file.FileHeader.FileInfo().Name()
		if name[len(name)-5:] == ".mscx" {
			fileName = name
			break
		}
	}

	if fileName == "" {
		log.Fatal("No mscx file found")
	}

	fileReader, err := fileZip.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// now read file from the reader
	newDataBuf := new(bytes.Buffer)
	_, err = newDataBuf.ReadFrom(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	newData := newDataBuf.Bytes()

	// marshal file
	result := new(mscoreNotationSchema.MUEFile)
	err = xml.Unmarshal(newData, result)
	if err != nil {
		log.Fatal(err)
	}

	// Parse our file into something golang-readable
	resultU, err := result.AsGoMusic()
	if err != nil {
		log.Fatal(err)
	}

	// now convert to Shawzin version
	resultS, err := shawzin.GetShawzin(resultU, shawzin.Chromatic)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultS)
}

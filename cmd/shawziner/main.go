package main

import (
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

	// marshal file
	result := new(mscoreNotationSchema.MUEFile)
	err = xml.Unmarshal(file, result)
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

	/*
		fmt.Println(result.Measures[2])
		fmt.Println(resultU)
	*/
	fmt.Println(resultS)
}

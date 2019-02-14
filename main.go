package main

import (
	"flag"
	"fmt"
	"os"
)

// @todo add mapping to custom types
// @todo discover receiver name if set has any method implemented
// @todo if two sets share the same item type a collision of names happens
func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "Name of the file to search for slice types")
	flag.Parse()

	fileTypes, err := parseFile(filename)
	if err != nil {
		fmt.Printf("Error parsing file %s %s\n", filename, err)
		os.Exit(1)
	}

	err = generateFile(fileTypes)
	if err != nil {
		fmt.Printf("Error generating file %s %s\n", filename, err)
		os.Exit(1)
	}
}

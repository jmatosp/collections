package main

import (
	"flag"
	"fmt"
	"os"
)

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

package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

type fileTypes struct {
	filename    string
	packageName string
	sliceTypes  []sliceType
}
type sliceType struct {
	name     string
	itemName string
}

func parseFile(filename string) (fileTypes, error) {
	file, err := os.Open(filename)
	if err != nil {
		return fileTypes{}, nil
	}

	return parseGo(filename, file)
}

// parseFile parses a Go file and returns all found
// slice types and their item type
func parseGo(filename string, r io.Reader) (fileTypes, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, r, 0)
	if err != nil {
		return fileTypes{}, err
	}

	var result []sliceType
	var packageName string

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.File:
			packageName = x.Name.Name
		case *ast.TypeSpec:
			arrayType, ok := x.Type.(*ast.ArrayType)
			if !ok {
				return true
			}
			sliceTypeName := x.Name.Name
			elt, ok := arrayType.Elt.(*ast.Ident)
			if !ok {
				return true
			}
			result = append(result, sliceType{
				name:     sliceTypeName,
				itemName: elt.Name,
			})
		}
		return true
	})

	return fileTypes{
		filename:    filename,
		packageName: packageName,
		sliceTypes:  result,
	}, nil
}

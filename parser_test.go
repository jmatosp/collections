package main

import (
	"strings"
	"testing"
)

func TestParseGo(t *testing.T) {
	source := `
package main
type MyType int
type MyTypeSlice []MyType
`

	typesFound, err := parseGo("source.go", strings.NewReader(source))

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if len(typesFound.sliceTypes) != 1 {
		t.Errorf("expecting 1 slice but found %d", len(typesFound.sliceTypes))
	}

	if typesFound.filename != "source.go" {
		t.Errorf("expecting source filename to be source.go but found %s", typesFound.filename)
	}

	if typesFound.packageName != "main" {
		t.Errorf("expecting package name to be main but found %s", typesFound.packageName)
	}

	if typesFound.sliceTypes[0].name != "MyTypeSlice" {
		t.Errorf("expecting slice type name to be MyTypeSlice but found %s", typesFound.sliceTypes[0].name)
	}

	if typesFound.sliceTypes[0].itemName != "MyType" {
		t.Errorf("expecting slice item type name to be MyType but found %s", typesFound.sliceTypes[0].itemName)
	}
}

func TestParseGo_MultipleSliceTypes(t *testing.T) {
	source := `
package main
type MyType1 int
type MyTypeSlice1 []MyType1
type MyType2 int
type MyTypeSlice2 []MyType2
`

	typesFound, err := parseGo("source.go", strings.NewReader(source))

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if len(typesFound.sliceTypes) != 2 {
		t.Errorf("expecting 2 slice but found %d", len(typesFound.sliceTypes))
	}

	if typesFound.sliceTypes[0].name != "MyTypeSlice1" {
		t.Errorf("expecting slice type name to be MyTypeSlice1 but found %s", typesFound.sliceTypes[0].name)
	}

	if typesFound.sliceTypes[0].itemName != "MyType1" {
		t.Errorf("expecting slice item type name to be MyType1 but found %s", typesFound.sliceTypes[0].itemName)
	}

	if typesFound.sliceTypes[1].name != "MyTypeSlice2" {
		t.Errorf("expecting slice type name to be MyTypeSlice2 but found %s", typesFound.sliceTypes[0].name)
	}

	if typesFound.sliceTypes[1].itemName != "MyType2" {
		t.Errorf("expecting slice item type name to be MyType2 but found %s", typesFound.sliceTypes[0].itemName)
	}
}

func TestParseGo_NoTypesFound(t *testing.T) {
	source := `package main`

	typesFound, err := parseGo("source.go", strings.NewReader(source))

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if len(typesFound.sliceTypes) != 0 {
		t.Errorf("not expecting to find any types but found %d type(s)", len(typesFound.sliceTypes))
	}
}

func TestParseGo_HandlesParsingError(t *testing.T) {
	source := `invalid source code`

	_, err := parseGo("source.go", strings.NewReader(source))

	if err == nil {
		t.Error("expecting an error got none")
	}
}

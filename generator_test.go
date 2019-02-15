package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

func TestGenerateSource(t *testing.T) {
	types := fileTypes{
		filename:    "source.go",
		packageName: "main",
		sliceTypes: []sliceType{
			{
				name:     "MyTypeList",
				itemName: "MyType",
			},
		},
	}

	source, err := generateSource(types)

	if err != nil {
		t.Errorf("unexpected error generating source got %s", err.Error())
	}

	goldenFile := filepath.Join("fixtures", "collection.go.golden")
	if *update {
		t.Log("update golden file")
		err = ioutil.WriteFile(goldenFile, []byte(source), 0644)
		if err != nil {
			t.Errorf("error generating golden file: %s", err.Error())
		}
	}

	goldenSource, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Errorf("error reading golden file: %s", err.Error())
	}

	if string(goldenSource) != source {
		t.Error("source generated and golden file don't match")
	}
}

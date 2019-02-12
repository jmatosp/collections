package main

import (
	"fmt"
	"strings"
	"testing"
)

var persons = Persons{
	{"John", 21},
	{"Hannah", 17},
	{"Hannibal", 57},
	{"Jose", 40},
	{"Joana", 20},
}

func ExampleAges_Filter() {
	filtered := persons.Filter(func(item Person) bool {
		return strings.HasPrefix(item.Name, "J")
	})

	fmt.Println(filtered)
	// Output:
	// {Hannah 17}
	// {Hannibal 57}
}

func TestNames_Filter(t *testing.T) {
	adultsFunc := func(item Person) bool { return item.Age >= 18 }

	underAge := persons.Filter(adultsFunc)

	if len(underAge) != 1 {
		t.Fail()
	}
}

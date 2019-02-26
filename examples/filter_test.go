package main

import (
	"fmt"
	"strings"
	"testing"
)

func ExamplePersons_Filter() {
	var persons = Persons{
		{"John", 21},
		{"Hannah", 17},
		{"Hannibal", 57},
		{"Jose", 40},
		{"Joana", 20},
	}

	filtered := persons.Filter(func(item Person) bool {
		return strings.HasPrefix(item.Name, "J")
	})

	fmt.Println(filtered)
	// Output:
	// {John 21}
	// {Jose 40}
	// {Joana 20}
}

func TestNames_Filter(t *testing.T) {
	var persons = Persons{
		{"John", 21},
		{"Hannah", 17},
		{"Hannibal", 57},
		{"Jose", 40},
		{"Joana", 20},
	}

	adultsFunc := func(item Person) bool { return item.Age >= 18 }

	underAge := persons.Filter(adultsFunc)

	if len(underAge) != 0 {
		t.Fail()
	}

	if underAge[0].Name != "John" {
		t.Fail()
	}
}

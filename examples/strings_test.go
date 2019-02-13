package main

import (
	"fmt"
)

func ExamplePerson_Strings() {
	var persons = Persons{
		{"John", 21},
		{"Hannah", 17},
		{"Hannibal", 57},
		{"Jose", 40},
		{"Joana", 20},
	}

	names := persons.MapToStrings(func(item Person) string {
		return item.Name
	})

	fmt.Println(names)
	// Output:
	// [John Hannah Hannibal Jose Joana]
}

package main

import "fmt"

func ExamplePersons_MapToStrings() {
	var persons = Persons{
		{"John", 21},
		{"Hannah", 17},
	}

	names := persons.MapToStrings(func(item Person) string {
		return item.Name
	})

	fmt.Println(names)
	// Output:
	// [John Hannah]
}

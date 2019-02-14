package main

import "fmt"

func ExamplePersons_MapToInts() {
	var persons = Persons{
		{"John", 21},
		{"Hannah", 17},
	}

	ages := persons.MapToInts(func(item Person) int {
		return item.Age
	})

	fmt.Println(ages)
	// Output:
	// [21 17]
}


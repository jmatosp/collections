package main

import "fmt"

func ExamplePersons_All() {
	var persons = Persons{
		{"John", 21},
		{"Hannah", 17},
		{"John", 51},
	}

	johns := persons.All(func(item Person) bool {
		return item.Name == "John"
	})

	fmt.Println(johns)
	// Output:
	// {John 21}
	// {John 51}
}

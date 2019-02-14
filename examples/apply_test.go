package main

import "fmt"

func ExamplePersons_Apply() {
	var persons = Persons{
		{"John", 21},
		{"Hannah", 17},
	}

	persons.Apply(func(item Person) {
		fmt.Println(item)
	})

	// Output:
	// {John 21}
	// {Hannah 17}
}

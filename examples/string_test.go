package main

import (
	"fmt"
)

func ExampleString() {
	var persons = Persons{
		{"John", 21},
	}

	fmt.Println(persons.String())
	// Output:
	// {John 21}
}

package main

import (
	"fmt"
)

func ExamplePersons_First() {
	var persons = Persons{
		{"John", 21},
		{"Hannah", 17},
		{"Hannibal", 57},
		{"Jose", 40},
		{"Joana", 20},
	}

	oldestPerson := persons.Sort(func(item, other Person) bool {
		return item.Age > other.Age
	}).First()

	fmt.Println(oldestPerson)
	// Output:
	// {Hannibal 57}
}

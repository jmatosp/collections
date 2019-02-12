package main

//go:generate collections -name filter.go

type Person struct {
	Name string
	Age  int
}

type Persons []Person

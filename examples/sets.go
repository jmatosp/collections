package main

//go:generate collections -file $GOFILE

type Person struct {
	Name string
	Age  int
}

type Persons []Person

type Order struct {
	Number int
	Units  int
	Price  int
}

type Orders []Order

func (o Order) Amount() int {
	return o.Units * o.Price
}

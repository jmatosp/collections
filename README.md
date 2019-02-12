# Collection

*** Experimental - first prototype ***

Dynamically generate custom type slice handling

Generics look a like

## Quick example

```go
package main

//go:generate collections -name=orders.go

type Order struct {
    Shipped  bool
    Customer string
}

type Orders []Order 

var orders = Orders{
	{true, "Phantom Softwares"},
	{true, "Ieworks"},
	{false, "Seawares"},
	{true, "Pixystems"},
}

func main() {
	// print all shipped orders sorted by customer name
	orders.
		All(func (item Order) { return order.Shipped } ).
		Sort(func(item, other Order) { return strings.Compare(item.Customer, other.Customer) } ). 
		Println()

    // Output:
    // {true Ieworks} 
    // {true Phantom Softwares} 
    // {true Pixystems} 
}

``` 

## Install

```
go get -u github.com/jmatosp/collections
``` 

## Usage

Add to you go source file the `go:generate`

```go
package main 

//go:generate collections -name=person.go

type Person struct {}

type Persons []Person

```

Or by command line

```
collections -name=mytype.go
```

## API

###  collection._Filter()_

Filters out elements from a collection, returns a new collection

```go
package main

import "strings"

//go:generate collections -name=filename.go

type Person struct {
	Name string
	Age  int
}

type Persons []Person

var persons = Persons{
	{"John", 21},
	{"Hannah", 17},
	{"Hannibal", 57},
	{"Jose", 40},
	{"Joana", 20},
}

func main() {
	nameStartsWithJ := func(item Person) bool { return strings.HasPrefix(item.Name, "J") }
	
	filtered := persons.Filter(nameStartsWithJ)

	fmt.Println(filtered)
	// Output:
	// {Hannah 17}
	// {Hannibal 57}
}

```

More examples in `example/` subfolder

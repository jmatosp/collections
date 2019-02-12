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

### collection._Filter( func(item {Type}) bool )_

Filters out elements using a function on each element, returns a new collection

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

###  collection._All( func(item {type}) bool )_

Similar to [Filter]() but element that pass the functions are returned in the new collection

```go
filtered := persons.Filter(func(item Person) { return item.Age >= 18 } )
```

###  collection._First() {type}_

Returns the first item in the collection

```go
oldestPerson := persons.
	Sort(func(item, other Person) { return item.Age > other.Age } ).
	First()
fmt.Println(oldestPerson)
// Output:
// {"Hannibal", 57},
```

###  collection._Sort(func(item, other {type}) bool)_

Sorts a collection by a provided function, returns a new collection

```go
persons.
	Sort(func(item, other Person) { return strings.Compare(item.Name, other.Name) } ).
	Println()
// Output:
// {"Hannah", 17},
// {"Hannibal", 57},
// {"Joana", 20},
// {"John", 21},
// {"Jose", 40},
```

###  collection._Apply(func(item {type}))_

Runs a function on every element of the collection.

Example sending email to every person:
```go
sendEmail := func(person Person) { service.SendWelcome(person) }
persons.Apply(sendEmail)
```

Example mapping between object types:
```go
var ages []int
persons.Apply(func(person Person) { ages = append(ages, person.Age) } )
fmt.Println(ages)
// Output:
// {21 17 57 40 20}
```

###  collection._Println()_

Prints all elements of collection one by line

###  collection._String()_

Returns a collection list representation


# Collections

Dynamically generate custom type slice handling.

This tool will generate methods to manipulate your slices in an functional sort of way.

The API includes: `All`, `Filter`, `Apply`, `Map`, `MapToInts`, `MapToStrings`, `Println`, `Print`, `String`

## Quick example

Example how to manipulate one of your custom slice type: 

```go
package main

import "strings"

//go:generate collections -file $GOFILE

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
		All(func(item Order) bool { return item.Shipped }).
		Sort(func(item, other Order) bool { return strings.Compare(item.Customer, other.Customer) < 0 }).
		Println()

	// Output:
	// {true Ieworks}
	// {true Phantom Softwares}
	// {true Pixystems}
}
``` 

More examples in [examples](examples) folder.

## Install

```
go get -u github.com/jmatosp/collections
``` 

## Usage

Add to you go source file the `go:generate`:

```go
package main 

//go:generate collections -file $GOFILE

type Person struct {}

type Persons []Person

```

Or by command line:

```
collections -file source.go
```

## API

### collection._Filter( func(item {Type}) bool )_

Filters out elements using a function on each element, returns a new collection.

[Example](examples/filter_test.go)

###  collection._All( func(item {type}) bool )_

Similar to Filter but element that pass the functions are returned in the new collection.

[Example](examples/all_test.go)

###  collection._Map() []{type}_

Maps the slice to a new collection of same type, doesnt mutate original slice.

[Example](examples/map_test.go)

###  collection._MapToInts() []int_

Maps the slice objects into a slice of ints.

[Example](examples/maptoints_test.go)

###  collection._MapToString() []string_

Maps the slice objects into a slice of strings.

[Example](examples/maptoints_test.go)

###  collection._First() {type}_

Returns the first item in the collection

[Example](examples/filter_test.go)

###  collection._Sort(func(item, other {type}) bool)_

Sorts a collection by a provided function, returns a new collection.

[Example](examples/sort_test.go)

###  collection._Apply(func(item {type}))_

Runs a function on every element of the collection.

[Example](examples/apply_test.go)

###  collection._Println()_

Prints all elements of collection one by line.

###  collection._String()_

Returns a collection list representation.

[Example](examples/string_test.go)

# Collection

*** Experimental - first prototype ***

Dynamically generate custom type slice handling

Generics look a like

Library done without any external libraries, only go standard library 

## Install

```
go get -u github.com/jmatos/collections
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
collection
```

## API

###  collection._Filter()_

Filters out elements from a collection

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

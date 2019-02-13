package main

func ExamplePersons_Sort() {
	orders := Orders{
		{Number: 1, Units: 1, Price: 10},
		{Number: 2, Units: 1, Price: 20},
		{Number: 3, Units: 1, Price: 30},
	}
	
	byPriceDesc := orders.Sort(func(item, other Order) bool {
		return item.Price > other.Price
	})

	byPriceDesc.Println()
	// Output:
	// {3 1 30}
	// {2 1 20}
	// {1 1 10}
}

package main

func ExampleOrders_Map() {
	orders := Orders{
		{Number: 1, Units: 1, Price: 10},
		{Number: 2, Units: 1, Price: 20},
		{Number: 3, Units: 1, Price: 30},
	}

	// multiply price by 2
	orders = orders.Map(func(item Order) Order {
		return Order{
			Number: item.Number,
			Units:  item.Units,
			Price:  item.Price * 2,
		}
	})

	orders.Println()
	// Output:
	// {1 1 20}
	// {2 1 40}
	// {3 1 60}
}

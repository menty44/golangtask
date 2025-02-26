package main

import "fmt"

// Definition of the product sctruct
type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    string
	SalesCount int
	ViewsCount int
}

func main() {
	fmt.Println("Scnip task")

	// Create sample data
	products := []Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: "2019-01-04", SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: "2012-01-04", SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: "2014-05-28", SalesCount: 1048, ViewsCount: 20123},
	}

	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}

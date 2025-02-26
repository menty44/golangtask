package main

import (
	"fmt"
	"sort"
)

// Definition of the product sctruct
type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    string
	SalesCount int
	ViewsCount int
}

// Sorter is the interface that all sorting strategies must implement.
type Sorter interface {
	Sort([]Product) []Product
}

// ByPrice sorts products by price.
type ByPrice struct{}

func (bp ByPrice) Sort(products []Product) []Product {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	return products
}

// BySalesPerView sorts products by the ratio of sales_count to views_count.
type BySalesPerView struct{}

func (bsv BySalesPerView) Sort(products []Product) []Product {
	sort.Slice(products, func(i, j int) bool {
		ratioI := float64(products[i].SalesCount) / float64(products[i].ViewsCount)
		ratioJ := float64(products[j].SalesCount) / float64(products[j].ViewsCount)
		return ratioI > ratioJ // Higher ratio comes first
	})
	return products
}

// ProductSorter sorts products using a specified sorting strategy.
type ProductSorter struct {
	Sorter Sorter
}

func (ps *ProductSorter) Sort(products []Product) []Product {
	return ps.Sorter.Sort(products)
}

func main() {

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Scnip task")
	fmt.Println("")

	// Create sample data
	products := []Product{
		{ID: 1, Name: "Alabaster Table", Price: 12.99, Created: "2019-01-04", SalesCount: 32, ViewsCount: 730},
		{ID: 2, Name: "Zebra Table", Price: 44.49, Created: "2012-01-04", SalesCount: 301, ViewsCount: 3279},
		{ID: 3, Name: "Coffee Table", Price: 10.00, Created: "2014-05-28", SalesCount: 1048, ViewsCount: 20123},
	}

	// Sort by price
	priceSorter := &ProductSorter{Sorter: ByPrice{}}
	sortedByPrice := priceSorter.Sort(products)
	fmt.Println("Sorted by Price:")
	for _, p := range sortedByPrice {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

	// Sort by sales per view ratio
	salesPerViewSorter := &ProductSorter{Sorter: BySalesPerView{}}
	sortedBySalesPerView := salesPerViewSorter.Sort(products)
	fmt.Println("\nSorted by Sales per View Ratio:")
	for _, p := range sortedBySalesPerView {
		ratio := float64(p.SalesCount) / float64(p.ViewsCount)
		fmt.Printf("ID: %d, Name: %s, Ratio: %.4f\n", p.ID, p.Name, ratio)
	}
}

package main

import "fmt"

type PricingStrategy interface {
	Calculate(Order) float64
	getName() string
}

type Order struct {
	Name  string
	Price float64
	Count int
}

type RegularPricing struct {
}

type SaleType struct {
	Event    string
	Discount int
}

type SalePricing struct {
	Sale SaleType
}

func (r RegularPricing) Calculate(o Order) float64 {
	return o.Price * float64(o.Count)
}

func (s SalePricing) Calculate(o Order) float64 {
	return o.Price * float64(o.Count) / 100 * (100 - float64(s.Sale.Discount))
}

func (r RegularPricing) getName() string {
	return "Regular"
}

func (s SalePricing) getName() string {
	return s.Sale.Event
}

func main() {

	strategies := []PricingStrategy{RegularPricing{},
		SalePricing{Sale: SaleType{Event: "BlackFriday", Discount: 15}},
		SalePricing{Sale: SaleType{Event: "NewYear'sDay", Discount: 10}},
		SalePricing{Sale: SaleType{Event: "Christmas", Discount: 5}},
		SalePricing{Sale: SaleType{Event: "RussianBlackFriday", Discount: -30}},
	}
	orders := []Order{{Name: "test1", Price: 100, Count: 6},
		{Name: "test2", Price: 346.67, Count: 22},
		{Name: "test3", Price: 23590, Count: 3},
		{Name: "test4", Price: 54.99, Count: 124}}

	for k := range orders {
		fmt.Printf("Order %#v \n", orders[k])
		for _, v := range strategies {
			fmt.Printf("  Total cost with %s price %.2f$ \n", v.getName(), v.Calculate(orders[k]))
		}

	}

}

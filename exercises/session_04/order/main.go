package main

import "fmt"

func main() {
	orders := []Order{
		{
			ID: 1,
			Items: []OrderItem{
				{ProductID: 1, Quantity: 2, Price: 50},
				{ProductID: 2, Quantity: 1, Price: 30},
			},
			Status: "pending",
		},
		{
			ID: 2,
			Items: []OrderItem{
				{ProductID: 3, Quantity: 15, Price: 100},
			},
			Status: "pending",
		},
	}

	rules := []DiscountRule{
		{MinAmount: 100, DiscountPercent: 5.0, Description: "5% off for small orders"},
		{MinAmount: 500, DiscountPercent: 10.0, Description: "10% off for medium orders"},
		{MinAmount: 1000, DiscountPercent: 15.0, Description: "15% off for bulk orders"},
	}

	processedOrders := ProcessOrders(orders, rules)
	stats := CalculateOrderStatistics(processedOrders)

	fmt.Println("--- Orders ---")
	for _, o := range processedOrders {
		fmt.Printf("ID: %d | Total: %.2f | Status: %s\n", o.ID, o.Total, o.Status)
	}

	fmt.Println("--- Statistics ---")
	fmt.Println("Total Orders:", stats["totalOrders"])
	fmt.Println("Total Revenue:", stats["totalRevenue"])
	fmt.Println("Total Discount:", stats["totalDiscount"])
	fmt.Println("Average Value:", stats["averageOrderValue"])
	fmt.Println("Completed:", stats["completedOrders"])
}

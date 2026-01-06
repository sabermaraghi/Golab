package main

import (
"errors"
"fmt"
)

type Product struct {
ID       int
Name     string
Quantity int
Price    float64
}

func AddProduct(products map[int]Product, product Product) error {
if product.ID == 0 {
  return errors.New("product ID cannot be zero")
}

if existing, ok := products[product.ID]; ok {
  existing.Quantity += product.Quantity
  products[product.ID] = existing
  return nil
}
products[product.ID] = product
return nil
}

func RemoveProduct(products map[int]Product, productID int) error {
if _, exists := products[productID]; !exists {
  return errors.New("product not found")
}
delete(products, productID)
return nil
}

// Check inventory
func CheckStock(inventory map[int]Product, productID int) (int, bool) {
product, exists := inventory[productID]
if !exists {
  return 0, false
}
return product.Quantity, true
}

func CalculateTotalValue(inventory map[int]Product) float64 {
var total float64 = 0.0
for _, product := range inventory {
  total += product.Price * float64(product.Quantity)
}
return total
}

func main() {
inventory := make(map[int]Product)

// Add Products
AddProduct(inventory, Product{ID: 1, Name: "Laptop", Quantity: 10, Price: 1500.0})
AddProduct(inventory, Product{ID: 2, Name: "Mouse", Quantity: 50, Price: 25.0})

// To Increase Quantity
AddProduct(inventory, Product{ID: 1, Name: "Laptop", Quantity: 5, Price: 1500.0})

// Inventory calculation
qty, exists := CheckStock(inventory, 1)
fmt.Printf("Laptop -> quantity: %d, exists: %v\n", qty, exists) // 15, true

// Calculate Total Value
total := CalculateTotalValue(inventory)
fmt.Printf("Total inventory value: %.2f\n", total) // 16250.00

// Remove Product
err := RemoveProduct(inventory, 2)
if err != nil {
  fmt.Println("Error:", err)
}

fmt.Printf("After remove -> total: %.2f\n", total) // 22500.00
}

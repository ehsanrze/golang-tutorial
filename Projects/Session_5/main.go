// main.go
package main

import (
	"ecommerce/inventory"
	"fmt"
)

// Generic function to apply a discount to any type that implements Item
func ApplyDiscount[T inventory.Item](item T, discount float64) T {
	fmt.Printf("Before discount: [%s] %s - $%.2f\n", item.GetID(), item.GetName(), item.GetPrice())

	item.ApplyDiscount(discount)

	fmt.Printf("After discount: [%s] %s - $%.2f\n", item.GetID(), item.GetName(), item.GetPrice())

	return item
}

// Function that prints any type of data
func PrintAny(data any) {
	fmt.Printf("Data: %#v\n", data)
}

// Additional: Using Generics with slices
func GetItemNames[T inventory.Item](items []T) []string {
	names := make([]string, len(items))
	for i, item := range items {
		names[i] = item.GetName()
	}
	return names
}

// Display inventory using array
func displayInventoryArray(items []inventory.Item) {
	fmt.Println("Inventory (Array):")
	for index, item := range items {
		fmt.Printf("%d. [%s] %s - $%.2f\n", index+1, item.GetID(), item.GetName(), item.GetPrice())
	}
	fmt.Println()
}

// Display inventory using map
func displayInventoryMap(items map[string]inventory.Item) {
	fmt.Println("Inventory (Map):")
	for id, item := range items {
		fmt.Printf("ID: %s | %s - $%.2f\n", id, item.GetName(), item.GetPrice())
	}
	fmt.Println()
}

func main() {
	// Initialize inventory items
	fixedInventory := []inventory.Item{
		&inventory.Book{ID: "B001", Name: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 39.99},
		&inventory.Electronics{ID: "E001", Name: "Smartphone", Brand: "TechBrand", Price: 299.99, Warranty: 24},
		&inventory.Book{ID: "B002", Name: "Introducing Go", Author: "Caleb Doxsey", Price: 29.99},
		&inventory.Electronics{ID: "E002", Name: "Laptop", Brand: "ComputeX", Price: 899.99, Warranty: 12},
	}

	// Initialize the inventoryMap
	inventoryMap := make(map[string]inventory.Item)
	for _, item := range fixedInventory {
		inventoryMap[item.GetID()] = item
	}

	// Display inventory using array
	displayInventoryArray(fixedInventory)

	// Display inventory using map
	displayInventoryMap(inventoryMap)

	// Using generics to get item names from array
	names := GetItemNames(fixedInventory)
	fmt.Println("Item Names:", names)
	fmt.Println()

	// Using interface{} (any) to print different types of data
	PrintAny("Hello, World!")
	PrintAny(12345)
	PrintAny(3.14159)
	PrintAny(fixedInventory[0])
	fmt.Println()

	// Iterate over the map and apply discount
	for _, item := range inventoryMap {
		// Apply a 10% discount
		ApplyDiscount(item, 10.0)
	}
}

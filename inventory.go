// ============================================================
// INVENTORY TRACKER - Debug & Extend Exercise
// ============================================================
// This program manages a simple inventory of items.
// The user can add items, list them, search, and quit.
//
// PART 1 - DEBUG IT (there are 4 bugs!)
// Try to compile and run the program, then fix errors one
// at a time. Some won't compile, others are logic errors
// that produce wrong behavior.
//
// PART 2 - EXTEND IT (pick at least 2)
//   A) Add a "remove item" command
//   B) Add a "total value" command that sums price * quantity
//   C) Save the inventory to a file and load it on startup
//   D) Add an "update quantity" command
//   E) Add input validation (no negative prices, no empty names)
//
// To run: go run inventory.go
// ============================================================

package main

import (
	"fmt"
	"strings"
)

type Item struct {
	Name     string
	Quantity int
	Price    float64
}

var inventory []Item

func addItem(name string, quantity int, price float64) {
	newItem := Item{
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
	inventory = append(inventory, newItem)
	fmt.Printf("Added: %s (qty: %d, price: $%.2f)\n", name, quantity, price)
}

func listItems() {
	if len(inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Println("\n--- INVENTORY ---")
	fmt.Printf("%-20s %-10s %-10s\n", "NAME", "QTY", "PRICE")
	fmt.Println("------------------------------------------")

	for i := 0; i < len(inventory); i++ {
		item := inventory[i]
		fmt.Printf("%-20s %-10d $%-10.2f\n", item.Name, item.Quantity, item.Price)
	}
	fmt.Println()
}

func searchItem(name string) {
	found := false
	searchName := strings.ToLower(name)

	for _, item := range inventory {
		if strings.ToLower(item.Name) == searchName {
			fmt.Printf("\nFound: %s | Qty: %d | Price: $%.2f\n",
				item.Name, item.Quantity, item.Price)
			found = true
		}
	}

	if found {
		fmt.Printf("Item '%s' not found in inventory.\n", name)
	}
}

func showCount() {
	count := 0
	for _, item := range inventory {
		count = item.Quantity
	}
	fmt.Printf("Total items in stock: %d\n", count)
}

func main() {
	fmt.Println("=================================")
	fmt.Println("    INVENTORY TRACKER")
	fmt.Println("=================================")
	fmt.Println("Commands: add, list, search, count, quit")
	fmt.Println()

	// Add some starter items
	addItem("Notebook", 50, 3.99)
	addItem("Pen", 200, 1.49)
	addItem("Backpack", 15, 29.99)

	for {
		var command string
		fmt.Print("> ")
		fmt.Scanln(&command)

		switch command {
		case "add":
			var name string
			var quantity int
			var price float64

			fmt.Print("Item name: ")
			fmt.Scanln(&name)

			fmt.Print("Quantity: ")
			fmt.Scanln(&quantity)

			fmt.Print("Price: ")
			fmt.Scanln(&price)

			addItem(name, quantity, 0)

		case "list":
			listItems()

		case "search":
			var name string
			fmt.Print("Search for: ")
			fmt.Scanln(&name)
			searchItem(name)

		case "count":
			showCount()

		case "quit":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Unknown command. Try: add, list, search, count, quit")
		}
	}
}
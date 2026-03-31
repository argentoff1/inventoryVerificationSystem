package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Price    float64
	Quantity int
}

func main() {
	var inventoryData string
	var operations string
	var params string
	fmt.Scanln(&inventoryData)
	fmt.Scanln(&operations)
	fmt.Scanln(&params)

	// Parse inventory data
	inventory := make(map[string]Product)
	if inventoryData != "" {
		products := strings.Split(inventoryData, ",")
		for _, productStr := range products {
			parts := strings.Split(productStr, ":")
			name := parts[0]
			price, _ := strconv.ParseFloat(parts[1], 64)
			quantity, _ := strconv.Atoi(parts[2])
			inventory[name] = Product{Price: price, Quantity: quantity}
		}
	}

	fmt.Println("=== INVENTORY MANAGEMENT SYSTEM ===")
	fmt.Printf("System initialized with %d products\n", len(inventory))
	fmt.Println("Starting interactive session...")
	// Parse operations
	operationParts := strings.Split(operations, ",")

	if params != "" {
		paramsList := strings.Split(params, "|")
		for i, params := range paramsList {
			switch operationParts[i] {
			case "check":
				parts := strings.Split(params, ":")
				productName := parts[0]
				stock, checkErr := checkStock(inventory, productName)
				fmt.Println("--- STOCK CHECK ---")
				fmt.Printf("Checking stock for: %s\n", productName)
				if checkErr == nil {
					fmt.Printf("Stock level: %d units\n", stock)
				} else {
					fmt.Println("Product not found in inventory")
				}
				fmt.Println("Operation completed. Continuing to next operation...")

			case "add":
				parts := strings.Split(params, ":")
				productName := parts[0]
				price, _ := strconv.ParseFloat(parts[1], 64)
				quantity, _ := strconv.Atoi(parts[2])
				addItemErr := addNewItem(inventory, productName, price, quantity)
				fmt.Println("--- ADD ITEM ---")
				fmt.Printf("Adding new product: %s\n", productName)
				if addItemErr == nil {
					fmt.Println("Product added successfully")
				} else {
					fmt.Printf("Failed to add product: %s\n", addItemErr)
				}
				fmt.Println("Operation completed. Continuing to next operation...")

			case "update":
				parts := strings.Split(params, ":")
				productName := parts[0]
				change, _ := strconv.Atoi(parts[1])
				updateErr := updateStock(inventory, productName, change)
				fmt.Println("--- UPDATE STOCK ---")
				fmt.Printf("Updating stock for: %s\n", productName)
				if updateErr == nil {
					if change > 0 {
						fmt.Printf("Added %d units. New stock: %d\n", change, inventory[productName].Quantity)
					} else if change < 0 {
						absChange := int(math.Abs(float64(change)))
						fmt.Printf("Removed %d units. New stock: %d\n", absChange, inventory[productName].Quantity)
					}
				} else {
					fmt.Printf("Update failed: %s\n", updateErr)
				}
				fmt.Println("Operation completed. Continuing to next operation...")

			case "report":
				parts := strings.Split(params, ",")
				reportType := parts[0]
				threshold, _ := strconv.ParseFloat(parts[1], 64)
				fmt.Println("--- GENERATE REPORT ---")
				fmt.Printf("Generating %s report with threshold %g\n", reportType, threshold)
				generateReport(inventory, reportType, threshold)
				fmt.Println("Operation completed. Continuing to next operation...")

			case "exit":
				fmt.Println("--- SYSTEM EXIT ---")
				totalProducts := len(inventory)
				totalItems := calcTotalQuantity(inventory)
				totalValue := calcTotalValue(inventory)
				fmt.Println("Final inventory status:")
				fmt.Printf("Total products: %d\n", totalProducts)
				fmt.Printf("Total items: %d\n", totalItems)
				fmt.Printf("Total value: $%.2f\n", totalValue)
				fmt.Println("Session completed successfully")
				fmt.Println("Thank you for using the Inventory Management System")
			}
		}
	}
}

func sortInventory(inventory map[string]Product) []string {
	keys := make([]string, 0, len(inventory))
	for key := range inventory {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func calcTotalQuantity(inventory map[string]Product) int {
	var totalQuantity int
	for key := range inventory {
		totalQuantity += inventory[key].Quantity
	}
	return totalQuantity
}

func calcTotalValue(inventory map[string]Product) float64 {
	var totalValue float64
	for key := range inventory {
		totalValue += inventory[key].Price * float64(inventory[key].Quantity)
	}
	return totalValue
}

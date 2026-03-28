package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Price    float64
	Quantity int
}

func main() {
	inventory := make(map[string]Product)

	// Read input using bufio.Scanner to handle spaces properly
	scanner := bufio.NewScanner(os.Stdin)

	// Данные в формате: "store_name,location"
	scanner.Scan()
	storeInfo := scanner.Text()

	// Данные в формате: "product1:price1:quantity1,product2:price2:quantity2
	scanner.Scan()
	productData := scanner.Text()

	// TODO: Write your code below
	// 1. Parse store information (split by comma, check length)
	storeInfoParts := strings.Split(storeInfo, ",")
	// Добавить обработку ошибки
	if len(storeInfoParts) != 2 {
		return
	}
	storeName := storeInfoParts[0]
	location := storeInfoParts[1]

	// 2. Parse product data (split by comma, then by colon for each product)
	entries := strings.Split(productData, ",")
	for _, entry := range entries {
		parts := strings.SplitN(strings.TrimSpace(entry), ":", 3)
		// Добавить обработку ошибки
		if len(parts) != 3 {
			continue
		}

		product := strings.TrimSpace(parts[0])
		priceStr := strings.TrimSpace(parts[1])
		quantityStr := strings.TrimSpace(parts[2])

		// Errors
		price, parseFloatErr := strconv.ParseFloat(priceStr, 64)
		if parseFloatErr != nil {
			continue
		}
		quantity, parseIntErr := strconv.Atoi(quantityStr)
		if parseIntErr != nil {
			continue
		}

		inventory[product] = Product{Price: price, Quantity: quantity}
	}
	// 3. Display store information
	fmt.Printf("=== %s Inventory System ===\n", storeName)
	fmt.Printf("Location: %s\n", location)
	fmt.Printf("Inventory initialized with %d products\n", len(inventory))
	// 4. Display current inventory (sorted alphabetically)
	sortedKeys := sortInventory(inventory)
	fmt.Println("Current Inventory:")
	for _, key := range sortedKeys {
		fmt.Printf("- %s: $%.2f (Stock: %d)\n", key, inventory[key].Price, inventory[key].Quantity)
	}
	// 5. Calculate and display inventory statistics
	fmt.Println("Inventory Statistics:")
	fmt.Printf("Total Products: %d\n", len(inventory))
	fmt.Printf("Total Items in Stock: %d\n", calcTotalQuantity(inventory))
	fmt.Printf("Total Inventory Value: $%.2f\n", calcTotalValue(inventory))
	// 6. Display system status
	fmt.Println("System Status: Ready")
	fmt.Println("Inventory management system initialized successfully")
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

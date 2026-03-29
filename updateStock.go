package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func updateStock(inventory map[string]Product, productName string, change int) error {
	if _, exists := inventory[productName]; exists {
		currentQty := inventory[productName].Quantity
		newQty := currentQty + change
		if newQty < 0 {
			absChange := int(math.Abs(float64(change)))
			return errors.New(fmt.Sprintf("insufficient stock: cannot reduce %s by %d, only %d available", productName, absChange, currentQty))
		}
		inventory[productName] = Product{
			Price:    inventory[productName].Price,
			Quantity: newQty,
		}
	} else {
		return errors.New(fmt.Sprintf("product not found: %s", productName))
	}
	return nil
}

func UpdateStockMain() {
	inventory := make(map[string]Product)

	// Read input using bufio.Scanner to handle spaces properly
	scanner := bufio.NewScanner(os.Stdin)

	// Строка, содержащая существующие данные инвентаря в формате: "product1:price1:quantity1,product2:price2:quantity2"
	scanner.Scan()
	productData := scanner.Text()

	// Строка, содержащая запросы на обновление запасов в формате: "product1:change1,product2:change2"
	scanner.Scan()
	changesRequest := scanner.Text()

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
		inventory[product] = Product{
			Price:    price,
			Quantity: quantity,
		}
	}

	fmt.Println("Processing Stock Updates:")
	entries = strings.Split(changesRequest, ",")
	var successUpdateCount, failUpdateCount int
	var itemsRejectedList []string
	for _, entry := range entries {
		parts := strings.SplitN(strings.TrimSpace(entry), ":", 2)
		// Добавить обработку ошибки
		if len(parts) != 2 {
			continue
		}
		product := strings.TrimSpace(parts[0])
		changesStr := strings.TrimSpace(parts[1])

		changes, parseIntErr := strconv.Atoi(changesStr)
		if parseIntErr != nil {
			continue
		}

		err := updateStock(inventory, product, changes)
		if err == nil {
			absChange := int(math.Abs(float64(changes)))
			if changes == 0 {
				fmt.Printf("%s: No change - Current stock: %d\n", product, inventory[product].Quantity)
			} else if changes > 0 {
				fmt.Printf("%s: Added %d units - New stock: %d\n", product, absChange, inventory[product].Quantity)
			} else if changes < 0 {
				fmt.Printf("%s: Removed %d units - New stock: %d\n", product, absChange, inventory[product].Quantity)
			}
			successUpdateCount++
		} else {
			fmt.Printf("%s: Update failed - %s\n", product, err)
			failUpdateCount++
			itemsRejectedList = append(itemsRejectedList, product)
		}
	}
	fmt.Println("Update Summary:")
	fmt.Printf("Updates processed: %d\n", len(entries))
	fmt.Printf("Updates successful: %d\n", successUpdateCount)
	fmt.Printf("Updates failed: %d\n", failUpdateCount)

	fmt.Println("Updated Inventory:")
	keys := make([]string, 0, len(inventory))
	for key := range inventory {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("- %s: $%.2f (Stock: %d)\n", key, inventory[key].Price, inventory[key].Quantity)
	}

	fmt.Println("Final Inventory Statistics:")
	totalProducts := len(inventory)
	itemsInStock := calcTotalQuantity(inventory)
	totalItemsValue := calcTotalValue(inventory)
	fmt.Printf("Total Products: %d\n", totalProducts)
	fmt.Printf("Total Items in Stock: %d\n", itemsInStock)
	fmt.Printf("Total Inventory Value: $%.2f\n", totalItemsValue)

	if len(itemsRejectedList) != 0 {
		fmt.Printf("Failed updates: %s\n", strings.Join(itemsRejectedList, ","))
	} else {
		fmt.Println("All stock updates were processed successfully")
	}
}

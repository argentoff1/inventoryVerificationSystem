package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func addNewItem(inventory map[string]Product, productName string, price float64, quantity int) error {
	if _, exists := inventory[productName]; exists {
		return errors.New("product already exists: " + productName)
	}
	inventory[productName] = Product{
		Price:    price,
		Quantity: quantity,
	}
	return nil
}

func AddItemMain() {
	inventory := make(map[string]Product)

	// Read input using bufio.Scanner to handle spaces properly
	scanner := bufio.NewScanner(os.Stdin)

	// Строка, содержащая существующие данные инвентаря в формате: "product1:price1:quantity1,product2:price2:quantity2"
	scanner.Scan()
	productData := scanner.Text()

	// Строка, содержащая новые продукты для добавления в формате: "product1:price1:quantity1,product2:price2:quantity2"
	scanner.Scan()
	productsToAdd := scanner.Text()

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

	fmt.Println("Adding New Items:")
	var itemsAddedCount, itemsRejectedCount int
	var itemsRejectedList []string
	entries = strings.Split(productsToAdd, ",")
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

		err := addNewItem(inventory, product, price, quantity)
		if err == nil {
			fmt.Printf("%s: Successfully added - $%.2f (Stock: %d)\n", product, price, quantity)
			itemsAddedCount++
		} else {
			fmt.Printf("%s: Failed to add - %s\n", product, err)
			itemsRejectedCount++
			itemsRejectedList = append(itemsRejectedList, product)
		}
	}

	fmt.Println("Addition Summary:")
	fmt.Printf("Items processed: %d\n", len(entries))
	fmt.Printf("Items added: %d\n", itemsAddedCount)
	fmt.Printf("Items rejected: %d\n", itemsRejectedCount)

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

	if itemsRejectedCount > 0 {
		fmt.Printf("Rejected products: %s\n", strings.Join(itemsRejectedList, ","))
	} else {
		fmt.Println("All new products were added successfully")
	}
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkStock(inventory map[string]Product, productName string) (int, error) {
	for key := range inventory {
		if inventory[key] == inventory[productName] {
			return inventory[productName].Quantity, nil
		}
	}

	err := errors.New(fmt.Sprintf("product not found: %s", productName))

	return 0, err
}

func CheckStockMain() {
	inventory := make(map[string]Product)

	// Read input using bufio.Scanner to handle spaces properly
	scanner := bufio.NewScanner(os.Stdin)

	// Строка, содержащая существующие данные инвентаря в формате: "product1:price1:quantity1,product2:price2:quantity2
	scanner.Scan()
	productData := scanner.Text()

	// Строка, содержащая запросы на проверку запасов в формате: "product1,product2,product3"
	scanner.Scan()
	productsToCheck := scanner.Text()

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
	productsToCheckParts := strings.Split(productsToCheck, ",")

	fmt.Println("Stock Check Results:")
	prodCheckedCount := len(productsToCheckParts)
	var prodFoundCount, prodNotFoundCount, totalStock int
	var prodNotFoundList []string
	for _, product := range productsToCheckParts {
		stock, err := checkStock(inventory, product)
		if err == nil {
			fmt.Printf("%s: %d units in stock\n", product, stock)
			prodFoundCount++
			totalStock += stock
		} else {
			fmt.Printf("%s: Error - %s\n", product, err)
			prodNotFoundCount++
			prodNotFoundList = append(prodNotFoundList, product)
		}
	}

	fmt.Println("Check Summary:")
	fmt.Printf("Products checked: %d\n", prodCheckedCount)
	fmt.Printf("Products found: %d\n", prodFoundCount)
	fmt.Printf("Products not found: %d\n", prodNotFoundCount)
	fmt.Printf("Total stock for found products: %d units\n", totalStock)

	if len(prodNotFoundList) != 0 {
		result := strings.Join(prodNotFoundList, ",")

		fmt.Printf("Missing products: %s\n", result)
	} else {
		fmt.Println("All requested products are available")
	}
}

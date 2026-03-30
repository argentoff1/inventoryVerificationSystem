package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func generateReport(inventory map[string]Product, reportType string, threshold float64) {
	// Display report header
	switch reportType {
	case "full":
		fmt.Println("=== FULL INVENTORY REPORT ===")
	case "low_stock":
		fmt.Println("=== LOW STOCK REPORT ===")
	case "high_value":
		fmt.Println("=== HIGH VALUE REPORT ===")
	}

	// Calculate general statistics
	totalProducts := len(inventory)
	totalItems := 0
	totalValue := 0.0

	for _, product := range inventory {
		totalItems += product.Quantity
		totalValue += product.Price * float64(product.Quantity)
	}

	averagePrice := totalValue / float64(totalItems)

	// Display general statistics
	fmt.Printf("Total Products: %d\n", totalProducts)
	fmt.Printf("Total Items in Stock: %d\n", totalItems)
	fmt.Printf("Total Inventory Value: $%.2f\n", totalValue)
	fmt.Printf("Average Product Price: $%.2f\n", averagePrice)

	// Filter products based on report type
	var filteredProducts []string
	for name, product := range inventory {
		switch reportType {
		case "full":
			filteredProducts = append(filteredProducts, name)
		case "low_stock":
			if product.Quantity <= int(threshold) {
				filteredProducts = append(filteredProducts, name)
			}
		case "high_value":
			productValue := product.Price * float64(product.Quantity)
			if productValue >= threshold {
				filteredProducts = append(filteredProducts, name)
			}
		}
	}

	// Sort filtered products alphabetically
	sort.Strings(filteredProducts)

	// Display filtered products section header
	switch reportType {
	case "full":
		fmt.Println("All Products:")
	case "low_stock":
		fmt.Printf("Products with stock ≤ %.0f:\n", threshold)
	case "high_value":
		fmt.Printf("Products with value ≥ $%.2f:\n", threshold)
	}

	// Display filtered products
	filteredItems := 0
	filteredValue := 0.0
	for _, name := range filteredProducts {
		product := inventory[name]
		productValue := product.Price * float64(product.Quantity)
		fmt.Printf("- %s: $%.2f × %d = $%.2f\n", name, product.Price, product.Quantity, productValue)
		filteredItems += product.Quantity
		filteredValue += productValue
	}

	// Display filtered statistics
	fmt.Println("Filtered Results:")
	fmt.Printf("Products shown: %d\n", len(filteredProducts))
	fmt.Printf("Items in filtered products: %d\n", filteredItems)
	fmt.Printf("Value of filtered products: $%.2f\n", filteredValue)

	// Find most and least expensive products
	var mostExpensiveName, leastExpensiveName string
	var mostExpensivePrice, leastExpensivePrice float64
	first := true

	var sortedNames []string
	for name := range inventory {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)

	for _, name := range sortedNames {
		product := inventory[name]
		if first {
			mostExpensiveName = name
			mostExpensivePrice = product.Price
			leastExpensiveName = name
			leastExpensivePrice = product.Price
			first = false
		} else {
			if product.Price > mostExpensivePrice {
				mostExpensiveName = name
				mostExpensivePrice = product.Price
			}
			if product.Price < leastExpensivePrice {
				leastExpensiveName = name
				leastExpensivePrice = product.Price
			}
		}
	}

	fmt.Println("Price Analysis:")
	fmt.Printf("Most expensive: %s at $%.2f\n", mostExpensiveName, mostExpensivePrice)
	fmt.Printf("Least expensive: %s at $%.2f\n", leastExpensiveName, leastExpensivePrice)

	// Find highest and lowest stock products
	var highestStockName, lowestStockName string
	var highestStock, lowestStock int
	first = true

	for _, name := range sortedNames {
		product := inventory[name]
		if first {
			highestStockName = name
			highestStock = product.Quantity
			lowestStockName = name
			lowestStock = product.Quantity
			first = false
		} else {
			if product.Quantity > highestStock {
				highestStockName = name
				highestStock = product.Quantity
			}
			if product.Quantity < lowestStock {
				lowestStockName = name
				lowestStock = product.Quantity
			}
		}
	}

	fmt.Println("Stock Analysis:")
	fmt.Printf("Highest stock: %s with %d units\n", highestStockName, highestStock)
	fmt.Printf("Lowest stock: %s with %d units\n", lowestStockName, lowestStock)

	// Inventory health assessment
	lowStockCount := 0
	highStockCount := 0
	highValueCount := 0

	for _, product := range inventory {
		if product.Quantity <= 5 {
			lowStockCount++
		}
		if product.Quantity > 20 {
			highStockCount++
		}
		productValue := product.Price * float64(product.Quantity)
		if productValue >= 500.0 {
			highValueCount++
		}
	}

	fmt.Printf("Low stock items (≤5): %d\n", lowStockCount)
	fmt.Printf("High stock items (>20): %d\n", highStockCount)
	fmt.Printf("High value items (≥$500): %d\n", highValueCount)

	// Report completion
	fmt.Println("Report generated successfully")
	fmt.Printf("Threshold applied: %.2f\n", threshold)
}

func GenerateReportMain() {
	var inventoryData string
	var reportConfig string
	fmt.Scanln(&inventoryData)
	fmt.Scanln(&reportConfig)

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

	// Parse report configuration
	configParts := strings.Split(reportConfig, ",")
	reportType := configParts[0]
	threshold, _ := strconv.ParseFloat(configParts[1], 64)

	// Generate report
	generateReport(inventory, reportType, threshold)
}

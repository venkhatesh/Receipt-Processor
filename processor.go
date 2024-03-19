package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"strings"
	"time"
	"unicode"
)

// this method is to count alphabet and numeric values in a string
func countAlphanumericChars(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}

// this method calculates points based on the given rules for a receipt
func calculatePoints(receipt Receipt) int {
	points := 0
	total, err := processDigits(receipt.Total)
	if err != nil {
		log.Printf("Error converting total to float64: %V", err)
		return 0
	}

	// to calculate Points for retailer name
	points += countAlphanumericChars(receipt.Retailer)

	// to calculate Points for round dollar total
	if total == math.Floor(total) {
		points += 50
	}

	// to calculate Points for total being a multiple of 0.25
	if math.Mod(total*100, 25) == 0 {
		points += 25
	}

	// to calculate Points for items
	points += (len(receipt.Items) / 2) * 5

	// to calculate Item description length and price calculation
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			fmt.Println(reflect.TypeOf((item.Price)))
			price, err := processDigits(item.Price)
			if err != nil {
				log.Printf("Error converting price to float: %v", err)
				return 0
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	// to calculate points for odd day
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// to calculate points for time between 2:00pm and 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}
	return points
}

package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

/**
* calculatePoints calculates the points for a given receipt based on specific rules.
* The rules include points for alphanumeric characters in the retailer name, round dollar amounts,
* multiples of 0.25, number of items, item description length, odd purchase date, and time of purchase.
*
* @param receipt Receipt - The receipt for which the points are being calculated.
*
* @return int - The total points calculated for the receipt.
*/
func calculatePoints(receipt Receipt) int {
	points := 0
	total, err := strconv.ParseFloat(receipt.Total, 64)

	if err != nil {
		log.Printf("Error parsing receipt total: %v", err)
		return points
	}

	// One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}

	// 50 points if the total is a round dollar amount with no cents
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Points based on item description and price
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
		}
	}

	// 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

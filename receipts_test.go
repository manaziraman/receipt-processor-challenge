package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* TestCalculatePoints tests the calculatePoints function with various test cases.
* It asserts that the calculated points match the expected points for different receipts.
*
* @param t *testing.T - The testing object provided by the Go testing framework.
*/
func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		receipt Receipt
		points  int
	}{
		{
			receipt: Receipt{
				Retailer:     "Test Store",
				PurchaseDate: "2021-08-06",
				PurchaseTime: "14:00",
				Items: []Item{
					{ShortDescription: "Item 1", Price: "10.00"},
					{ShortDescription: "Item 2", Price: "20.00"},
				},
				Total: "30.00",
			},
			points: 105,
		},
		{
			receipt: Receipt{
				Retailer:     "Another Store",
				PurchaseDate: "2021-08-05",
				PurchaseTime: "15:30",
				Items: []Item{
					{ShortDescription: "Item A", Price: "5.00"},
					{ShortDescription: "Item B", Price: "5.25"},
					{ShortDescription: "Item C", Price: "5.50"},
				},
				Total: "15.75",
			},
			points: 63,
		},
		{
			receipt: Receipt{
				Retailer:     "",
				PurchaseDate: "2021-08-07",
				PurchaseTime: "13:59",
				Items:        []Item{},
				Total:        "0.00",
			},
			points: 81,
		},
		{
			receipt: Receipt{
				Retailer:     "Odd Date Store",
				PurchaseDate: "2021-08-07",
				PurchaseTime: "14:00",
				Items: []Item{
					{ShortDescription: "Item X", Price: "3.00"},
					{ShortDescription: "Item Y", Price: "3.00"},
				},
				Total: "6.00",
			},
			points: 110,
		},
	}

	for _, test := range tests {
		points := calculatePoints(test.receipt)
		assert.Equal(t, test.points, points, "calculatePoints(%v) got %v, want %v", test.receipt, points, test.points)
	}
}

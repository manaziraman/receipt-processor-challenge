package main

/**
* Receipt represents the structure of a receipt, including details such as the retailer,
* purchase date and time, items, and total amount.
*
* Fields:
*  Retailer     string - The name of the retailer.
*  PurchaseDate string - The date of purchase in the format "YYYY-MM-DD".
*  PurchaseTime string - The time of purchase in the format "HH:MM".
*  Items        []Item - An array of items purchased.
*  Total        string - The total amount of the purchase.
*/
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

/**
* Item represents an individual item in a receipt, including details such as the short description
* and price of the item.
*
* Fields:
*  ShortDescription string - A brief description of the item.
*  Price            string - The price of the item.
*/
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

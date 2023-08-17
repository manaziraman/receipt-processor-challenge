package main

import (
	"github.com/google/uuid"
)

// In-memory database to store receipts and points
var receiptsDB = make(map[string]int)

/**
* generateUUID generates a new UUID.
*
* @return string - The generated UUID.
* @return error - An error if generating the UUID fails.
*/
func generateUUID() (string, error) {
	id, err := uuid.NewRandom() // Negligible chance of duplicate IDs
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

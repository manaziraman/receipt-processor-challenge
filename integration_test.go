package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

/**
* TestIntegration tests the integration of the ProcessReceipts and GetPoints endpoints.
* It loads example receipts from the examples directory, processes them, and retrieves
* the points.
*
* We load json files from the examples directory, convert them to receipts, and then
* back to json to isolate the json formatting from the endpoints.
*
* @param t *testing.T - The testing object provided by the Go testing framework.
*/
func TestIntegration(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", ProcessReceipts).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

	// Get all JSON files in the examples directory
	exampleFiles, err := filepath.Glob("examples/*.json")
	if err != nil {
		t.Fatalf("Failed to list example files: %v", err)
	}

	for _, exampleFile := range exampleFiles {
		t.Run(filepath.Base(exampleFile), func(t *testing.T) {
			// Load the example receipt
			receipt, err := loadExampleReceipt(exampleFile)
			if err != nil {
				t.Fatalf("Failed to load example receipt: %v", err)
			}

			// Convert the receipt to JSON
			input, err := json.Marshal(receipt)
			if err != nil {
				t.Fatalf("Failed to marshal receipt: %v", err)
			}

			// Test processing the receipt
			req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(input))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)

			var response map[string]string
			json.Unmarshal(rr.Body.Bytes(), &response)
			receiptID := response["id"]

			// Test retrieving points for the processed receipt
			req, _ = http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
			rr = httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)

			var pointsResponse map[string]int
			json.Unmarshal(rr.Body.Bytes(), &pointsResponse)
			assert.Equal(t, calculatePoints(*receipt), pointsResponse["points"])
		})
	}
}

/**
* loadExampleReceipt loads a receipt from a JSON file.
*
* @param filename string - The path to the JSON file containing the receipt.
*
* @return *Receipt - A pointer to the loaded receipt.
* @return error - An error if loading the receipt fails.
*/
func loadExampleReceipt(filename string) (*Receipt, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var receipt Receipt
	if err := json.NewDecoder(file).Decode(&receipt); err != nil {
		return nil, err
	}

	return &receipt, nil
}

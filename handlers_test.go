package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

/**
* TestProcessReceipts tests the ProcessReceipts function with various test cases.
*
* @param t *testing.T - The testing object provided by the Go testing framework.
*/
func TestProcessReceipts(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		method         string
		input          string
		expectedStatus int
	}{
		{"Valid Receipt", "POST", `{"retailer": "Test", "purchaseDate": "2021-01-01", "items": [{"shortDescription": "item1", "price": "10"}], "total": "10"}`, http.StatusOK},
		{"Missing Retailer", "POST", `{"purchaseDate": "2021-01-01", "items": [{"shortDescription": "item1", "price": "10"}], "total": "10"}`, http.StatusBadRequest},
		{"Invalid Total Value", "POST", `{"retailer": "Test", "purchaseDate": "2021-01-01", "items": [{"shortDescription": "item1", "price": "10"}], "total": "-10"}`, http.StatusBadRequest},
		{"Malformed JSON", "POST", `{"retailer": "Test", "purchaseDate": "2021-01-01", "items": [{"shortDescription": "item1", "price": "10"}, "total": "10"}`, http.StatusBadRequest},
		{"Unsupported Method GET", "GET", "", http.StatusMethodNotAllowed},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest(test.method, "/receipts/process", bytes.NewBufferString(test.input))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ProcessReceipts)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatus, rr.Code)
		})
	}
}

/**
* TestGetPoints tests the GetPoints function with test cases for valid and
* non-existent receipt IDs.
*
* @param t *testing.T - The testing object provided by the Go testing framework.
*/
func TestGetPoints(t *testing.T) {
	// Test case for valid receipt ID
	{
		// Setup a test receipt
		receiptID, _ := generateUUID()
		receiptsDB[receiptID] = 100

		req, _ := http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
		rr := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/receipts/{id}/points", GetPoints)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response map[string]int
		json.Unmarshal(rr.Body.Bytes(), &response)
		assert.Equal(t, 100, response["points"])
	}

	// Test case for non-existent receipt ID
	{
		nonExistentID := "non-existent-id"

		req, _ := http.NewRequest("GET", "/receipts/"+nonExistentID+"/points", nil)
		rr := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/receipts/{id}/points", GetPoints)
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	}
}

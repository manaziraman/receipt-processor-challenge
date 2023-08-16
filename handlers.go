package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
ProcessReceipts handles the processing of receipts. It validates the receipt and
calculates the points.

@param w http.ResponseWriter - The HTTP response writer.
@param r *http.Request - The HTTP request.
*/
func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	// Check for unsupported HTTP method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check for request size limit (e.g., 1MB) and throw error if exceeded
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		// Handle malformed JSON
		log.Printf("Error decoding JSON: %v", err)
		log.Printf("Request body: %v", r.Body)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || len(receipt.Items) == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		// Handle invalid total
		log.Printf("Error parsing total: %v", err)
		http.Error(w, "Invalid total value", http.StatusBadRequest)
		return
	}

	// Validate constraints (e.g., positive total)
	if total <= 0 {
		http.Error(w, "Invalid total value", http.StatusBadRequest)
		return
	}

	id, err := generateUUID()
	if err != nil {
		log.Printf("Error generating UUID: %v", err)

		// Handle error, e.g., log it and return an internal server error response
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	receiptsDB[id] = calculatePoints(receipt)

	// Successful response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

/*
GetPoints handles the retrieval of points for a given receipt ID.

@param w http.ResponseWriter - The HTTP response writer.
@param r *http.Request - The HTTP request.
*/
func GetPoints(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt ID from the URL path
	vars := mux.Vars(r)
	receiptID := vars["id"]

	// Retrieve the points for the given ID from memory
	points, exists := receiptsDB[receiptID]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Return the points as a JSON object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}

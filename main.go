package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
main initializes the HTTP server and sets up the routes for processing receipts and retrieving points.
It starts the server at http://localhost:8080 and listens for incoming HTTP requests.
*/
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", ProcessReceipts).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
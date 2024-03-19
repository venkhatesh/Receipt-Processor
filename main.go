package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Receipt Structure
type Receipt struct {
	ID            string      `json:"id"`
	Retailer      string      `json:"retailer"`
	PurchaseDate  string      `json:"purchaseDate"`
	PurchaseTime  string      `json:"purchaseTime"`
	Items         []Item      `json:"items"`
	Total         interface{} `json:"total"` // used interface so that the service can handle both float and string inputs
	PointsAwarded int         `json:"-"`
}

type Item struct {
	ShortDescription string      `json:"shortDescription"`
	Price            interface{} `json:"price"` // used interface so that the service can handle both float and string inputs
}

var receipts = make(map[string]*Receipt)

// this method checks if the provided json body has string or float for points and total params
func processDigits(digit interface{}) (float64, error) {
	switch v := digit.(type) {
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("unsupported type for digit: %T", v)
	}
}

func processReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	receipt.ID = uuid.NewString()
	receipt.PointsAwarded = calculatePoints(receipt)
	receipts[receipt.ID] = &receipt

	resp := map[string]string{"id": receipt.ID}
	json.NewEncoder(w).Encode(resp)
}

func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	normalizedPath := strings.Trim(r.URL.Path, "/")
	pathSegments := strings.Split(normalizedPath, "/")

	// Check if the URL matches the expected format: "receipts/{id}/points"
	if len(pathSegments) != 3 || pathSegments[0] != "receipts" || pathSegments[2] != "points" {
		errMsg := fmt.Sprintf("Invalid URL format. Expected /receipts/{id}/points but got /%s", normalizedPath)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	id := pathSegments[1]

	// Check if the {id} segment is empty, which would happen with a URL like /receipts//points
	if id == "" {
		http.Error(w, "Missing receipt ID in the URL", http.StatusBadRequest)
		return
	}
	// printReceiptIDs() // this method is for debugging
	receipt, exists := receipts[id]
	if !exists {
		http.Error(w, "Receipts not found", http.StatusNotFound)
		return
	}
	resp := map[string]int{"points": receipt.PointsAwarded}
	json.NewEncoder(w).Encode(resp)
}

// this method is for debugging
func printReceiptIDs() {
	for id := range receipts {
		fmt.Println(id)
	}
}

func main() {
	http.HandleFunc("/receipts/process", processReceiptsHandler)
	http.HandleFunc("/receipts/", getPointsHandler)
	fmt.Println("Startring server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

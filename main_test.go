package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func populateSampleData() {
	// test data
	sampleReceipts := []Receipt{
		{
			ID:           "123",
			Retailer:     "Supermarket",
			PurchaseDate: "2023-03-15",
			PurchaseTime: "14:30",
			Items: []Item{
				{ShortDescription: "Apples", Price: 3.50},
				{ShortDescription: "Bread", Price: 2.25},
			},
			Total: "5.75",
		},
		{
			ID:           "124",
			Retailer:     "Electronics Store",
			PurchaseDate: "2023-03-16",
			PurchaseTime: "16:45",
			Items: []Item{
				{ShortDescription: "USB-C Cable", Price: 19.99},
				{ShortDescription: "Headphones", Price: 89.99},
			},
			Total: "109.98",
		},
	}

	for _, receipt := range sampleReceipts {
		receipts[receipt.ID] = &receipt
	}
}

func TestGetPointsHandler(t *testing.T) {

	populateSampleData()
	tests := []struct {
		name       string
		url        string
		wantStatus int
		wantBody   string
	}{
		//test if the ID is present in the local database and if gives OK, 200 response
		{"Valid Receipt ID & URL", "/receipts/123/points", http.StatusOK, ""},
		//test if the service gives bad request response when the url path is not proper
		{"Invalid URL - missing points", "/receipts/123/", http.StatusBadRequest, ""},
		//test if the service give not found status if the id is not found in the local database
		{"ID not found - extra segment", "/receipts/400/points", http.StatusNotFound, ""},
		//test if the service gives bad request when if is not found in the url path
		{"Invalid URL - no ID", "/receipts//points", http.StatusBadRequest, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(getPointsHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.wantStatus)
			}

			if tc.wantBody != "" && !strings.Contains(rr.Body.String(), tc.wantBody) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tc.wantBody)
			}
		})
	}
}

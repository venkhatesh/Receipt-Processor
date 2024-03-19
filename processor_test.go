package main

import (
	"testing"
)

// test calculate points method
func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name       string
		receipt    Receipt
		wantPoints int
	}{
		{
			name: "Valid receipt",
			receipt: Receipt{
				Retailer:     "Test Store",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "15:00",
				Items: []Item{
					{ShortDescription: "Item 1", Price: 10.01},
				},
				Total: "20.00",
			},
			wantPoints: 103,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPoints := calculatePoints(tc.receipt)
			if gotPoints != tc.wantPoints {
				t.Errorf("calculatePoints() = %v, want %v", gotPoints, tc.wantPoints)
			}
		})
	}
}

package entity

import (
	"testing"
	
)
func TestBookPrice(t *testing.T) {
	books := Books{

		Title: "Phitchatuy",
		Price: 25,
		Code:  "B31234",
	}

	err := books.Validate()
	if err != nil {
		t.Errorf("Price must be between 50 and 5000")
	}
}

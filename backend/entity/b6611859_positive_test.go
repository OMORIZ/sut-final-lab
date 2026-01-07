package entity

import (
	"testing"
	
)
func TestBook(t *testing.T) {
	books := Books{

		Title: "Phitchatuy",
		Price: 250,
		Code:  "B31234",
	}

	err := books.Validate()
	if err != nil {
		t.Errorf("err ,%v", err)
	}
}

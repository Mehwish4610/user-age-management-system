package models

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	dob := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	age := CalculateAge(dob)

	currentYear := time.Now().Year()
	expectedAge := currentYear - 2000

	if age != expectedAge {
		t.Errorf("Expected %d, got %d", expectedAge, age)
	}
}
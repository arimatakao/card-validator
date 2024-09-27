package validator

import (
	"strconv"
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	currentTime := time.Now()
	currentMonth := strconv.Itoa(int(currentTime.Month()))
	currentYear := strconv.Itoa(currentTime.Year())
	futureTime := currentTime.AddDate(1, 1, 1)
	futureMonth := strconv.Itoa(int(futureTime.Month()))
	futureYear := strconv.Itoa(futureTime.Year())

	tests := []struct {
		testName        string
		cardNumber      string
		expirationMonth string
		expirationYear  string
		expectedValid   bool
	}{
		{
			testName:        "Valid card",
			cardNumber:      "1234567812345670",
			expirationMonth: futureMonth,
			expirationYear:  futureYear,
			expectedValid:   true,
		},
		{
			testName:        "Short card number",
			cardNumber:      "123",
			expirationMonth: futureMonth,
			expirationYear:  futureYear,
			expectedValid:   false,
		},
		{
			testName:        "Long card number",
			cardNumber:      "123456781234567890",
			expirationMonth: futureMonth,
			expirationYear:  futureYear,
			expectedValid:   false,
		},
		{
			testName:        "Empty card number",
			cardNumber:      "",
			expirationMonth: futureMonth,
			expirationYear:  futureYear,
			expectedValid:   false,
		},
		{
			testName:        "Empty expiration month",
			cardNumber:      "1234567812345670",
			expirationMonth: "",
			expirationYear:  futureYear,
			expectedValid:   false,
		},
		{
			testName:        "Non-numeric expiration month",
			cardNumber:      "1234567812345670",
			expirationMonth: "abc",
			expirationYear:  futureYear,
			expectedValid:   false,
		},
		{
			testName:        "Empty expiration year",
			cardNumber:      "1234567812345670",
			expirationMonth: futureMonth,
			expirationYear:  "",
			expectedValid:   false,
		},
		{
			testName:        "Non-numeric expiration year",
			cardNumber:      "1234567812345670",
			expirationMonth: futureMonth,
			expirationYear:  "xyz",
			expectedValid:   false,
		},
		{
			testName:        "Expired card",
			cardNumber:      "1234567812345670",
			expirationMonth: currentMonth,
			expirationYear:  currentYear,
			expectedValid:   false,
		},
		{
			testName:        "Invalid card number (Luhn check)",
			cardNumber:      "1234567812345671",
			expirationMonth: futureMonth,
			expirationYear:  futureYear,
			expectedValid:   false,
		},
		{
			testName:        "Malformed card number",
			cardNumber:      "1234abcd1234567",
			expirationMonth: futureMonth,
			expirationYear:  futureYear,
			expectedValid:   false,
		},
	}

	for _, test := range tests {
		valid, err := IsValid(test.cardNumber, test.expirationMonth, test.expirationYear)
		if valid != test.expectedValid {
			t.Errorf("Test: %s\nExpected valid: %v, got: %v error: %d, %s",
				test.testName,
				test.expectedValid, valid, err.GetCode(), err.GetMessage())
		}
	}
}

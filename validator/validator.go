package validator

import (
	"strconv"
	"time"
)

func IsValid(cardNumber string, expirationMonth string, expirationYear string) (bool, ValidErr) {
	// Check data format
	if cardNumber == "" || len(cardNumber) < 12 || len(cardNumber) > 19 {
		return false, newErr(ErrNumberLenght, "Card number must be between 12 and 19 digits")
	}

	if expirationMonth == "" {
		return false, newErr(ErrMonthEmpty, "Expiration month field is required")
	}
	monthNumber, err := strconv.Atoi(expirationMonth)
	if err != nil {
		return false, newErr(ErrMonthNotNumber, "Expiration month must be a number")
	}

	if expirationYear == "" {
		return false, newErr(ErrYearEmpty, "Expiration year field is required")
	}
	yearNumber, err := strconv.Atoi(expirationYear)
	if err != nil {
		return false, newErr(ErrYearNotNumber, "Expiration year must be a number")
	}

	// Check expiration date
	currentYear, currentMonth, _ := time.Now().Date()
	if yearNumber < currentYear ||
		(yearNumber == currentYear && monthNumber <= int(currentMonth)) {
		return false, newErr(ErrDateExpired, "The card has expired")
	}

	// Luhn card number validation
	sum := 0
	shouldDouble := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false, newErr(ErrNumberMalformated,
				"Card number contains invalid characters")
		}

		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble
	}

	if sum%10 != 0 {
		return false, newErr(ErrNumberNotValid, "Card number is not valid")
	}

	return true, nil
}

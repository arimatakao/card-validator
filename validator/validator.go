package validator

func IsValid(cardNumber string, expirationMonth string, expirationYear string) (bool, ValidErr) {
	return false, validationError{
		Code:    ErrUnknown,
		Message: "Unknown error",
	}
}

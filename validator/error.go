package validator

type ValidErr interface {
	GetCode() int
	GetMessage() string
}

const (
	ErrNumberLenght = iota + 1
	ErrNumberMalformated
	ErrNumberNotValid
	ErrMonthEmpty
	ErrMonthNotNumber
	ErrYearEmpty
	ErrYearNotNumber
	ErrDateExpired
)

type validationError struct {
	Code    int
	Message string
}

func newErr(code int, message string) validationError {
	return validationError{
		Code:    code,
		Message: message,
	}
}

func (e validationError) GetCode() int {
	return e.Code
}

func (e validationError) GetMessage() string {
	return e.Message
}

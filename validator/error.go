package validator

type ValidErr interface {
	GetCode() int
	GetMessage() string
}

const (
	ErrUnknown int = iota + 1
	ErrNumber
	ErrMonth
	ErrYear
)

type validationError struct {
	Code    int
	Message string
}

func (e validationError) GetCode() int {
	return e.Code
}

func (e validationError) GetMessage() string {
	return e.Message
}

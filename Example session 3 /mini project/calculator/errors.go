package calculator

import "errors"

var (
	ErrDivideByZero    = errors.New("division by zero is not allowed")
	ErrInvalidInput    = errors.New("invalid input")
	ErrInvalidOperator = errors.New("invalid operator")
)

package calculator

import "fmt"

func PerformOperation(a, b *float64, operator string) (*float64, error) {
	if a == nil || b == nil {
		return nil, fmt.Errorf("inputs cannot be nil")
	}

	var result float64
	switch operator {
	case "+":
		result = *a + *b
	case "-":
		result = *a - *b
	case "*":
		result = *a * *b
	case "/":
		if *b == 0 {
			return nil, fmt.Errorf("division by zero is not allowed")
		}
		result = *a / *b
	default:
		return nil, fmt.Errorf("invalid operator: %s", operator)
	}

	return &result, nil
}

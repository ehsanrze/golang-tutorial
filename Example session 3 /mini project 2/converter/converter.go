package converter

import "fmt"

func ConvertTemperature(value *float64, unit string) (*float64, error) {
	if value == nil {
		return nil, fmt.Errorf("temperature value cannot be nil")
	}

	var result float64
	switch unit {
	case "C":
		result = (*value * 9 / 5) + 32
	case "F":
		result = (*value - 32) * 5 / 9
	default:
		return nil, fmt.Errorf("invalid unit: %s (must be 'C' or 'F')", unit)
	}

	return &result, nil
}

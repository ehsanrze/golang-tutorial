package converter

import "errors"

var (
	ErrInvalidUnit    = errors.New("invalid unit: must be 'C' or 'F'")
	ErrNilTemperature = errors.New("temperature value cannot be nil")
)

package main

import (
	"fmt"
	"github.com/ehsanrze/golang-tutorial/mini-project2/converter"
)

func main() {
	celsius := 25.0
	fahrenheit := 77.0

	result, err := converter.ConvertTemperature(&celsius, "C")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%.2f°C is %.2f°F\n", celsius, *result)
	}

	result, err = converter.ConvertTemperature(&fahrenheit, "F")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%.2f°F is %.2f°C\n", fahrenheit, *result)
	}

	invalidUnit := "K"
	result, err = converter.ConvertTemperature(&celsius, invalidUnit)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", *result)

	}

	var nilValue *float64
	result, err = converter.ConvertTemperature(nilValue, "C")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

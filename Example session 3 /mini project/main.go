package main

import (
	"fmt"
	"github.com/ehsanrze/golang-tutorial/mini-project/calculator"
)

func main() {
	a, b := 10.0, 5.0

	result, err := calculator.PerformOperation(&a, &b, "+")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Addition Result: %.2f\n", *result)
	}

	result, err = calculator.PerformOperation(&a, &b, "/")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Division Result: %.2f\n", *result)
	}

	result, err = calculator.PerformOperation(&a, &b, "%")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", *result)
	}

	zero := 0.0
	result, err = calculator.PerformOperation(&a, &zero, "/")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", *result)
	}
}

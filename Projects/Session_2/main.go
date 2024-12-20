package main

import (
	"fmt"
)

// Function to display the menu options
func displayMenu() {
	fmt.Println("ATM Menu:")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

// Function to simulate balance check
func checkBalance(balance float64) {
	fmt.Printf("Your current balance is: $%.2f\n", balance)
}

// Function to simulate depositing money
func depositMoney(balance *float64, amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid deposit amount. Please enter a positive value.")
	} else {
		*balance += amount
		fmt.Printf("You deposited $%.2f. Your new balance is: $%.2f\n", amount, *balance)
	}
}

// Function to simulate withdrawing money
func withdrawMoney(balance *float64, amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid withdrawal amount. Please enter a positive value.")
	} else if amount > *balance {
		fmt.Println("Insufficient funds.")
	} else {
		*balance -= amount
		fmt.Printf("You withdrew $%.2f. Your new balance is: $%.2f\n", amount, *balance)
	}
}

func main() {
	var balance float64 = 1000.00 // Initial balance
	var choice int

	for {
		displayMenu()
		fmt.Print("Enter your choice (1-4): ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			continue
		}

		switch choice {
		case 1:
			checkBalance(balance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			_, err := fmt.Scan(&depositAmount)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid amount.")
				continue
			}
			depositMoney(&balance, depositAmount)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			_, err := fmt.Scan(&withdrawAmount)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid amount.")
				continue
			}
			withdrawMoney(&balance, withdrawAmount)
		case 4:
			fmt.Println("Thank you for using the ATM. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option from the menu.")
		}
	}
}

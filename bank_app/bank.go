package main

import (
	"fmt"
)

func main() {

	var balance float64 = 1000
	var money float64

	fmt.Println("Welcome to GO bank")
	for {

		fmt.Println("What do you want to do ??")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// switch choice {
		// case 1:
		// 	fmt.Printf("Your balance is: %.2f\n", balance)
		// case 2:
		// 	fmt.Print("Enter amount you want to deposite: ")
		// 	fmt.Scan(&money)
		// 	balance = balance + money
		// default:
		// 	println("Good bye !!!!")
		// 	break
		// }

		if choice == 1 {
			fmt.Printf("Your balance is: %.2f\n", balance)
		} else if choice == 2 {
			fmt.Print("Enter amount you want to deposite: ")
			fmt.Scan(&money)
			balance = balance + money
		} else if choice == 3 {
			fmt.Print("Enter amount you want to withdraw: ")
			fmt.Scan(&money)
			balance = balance - money
		} else {
			println("Good bye !!!!")
			break
		}

	}

}

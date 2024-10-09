package main

import (
	"fmt"
)

func main() {
	var amount, tip float64

	fmt.Print("Enter bill amount: ")
	fmt.Scanln(&amount)
	fmt.Print("Enter the tip percentage: ")
	fmt.Scanln(&tip)

	tip_calc := (tip / 100) * amount
	fmt.Printf("Your tip in amount: %.2f\n", tip_calc)

	total_amount := amount + tip_calc
	fmt.Printf("Your total amount: %.2f\n", total_amount)
}
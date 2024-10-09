package main

import (
	"fmt"
)

// Celcius to Fahrenheit: (celcius * 9/5) + 32
// Fahrenheit to Celcius: (fahrenheit - 32) * 5/9


func celciusToFahrenheit(celcius float64) float64 {
	return (celcius * 9/5) + 32
}

func fahrenheitToCelcius(fahrenheit float64) float64{
	return (fahrenheit - 32) * 5/9
}

func main() {
	fmt.Println("Temperature Convertor")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celcius")
	fmt.Print("Enter the conversion type (1 or 2): ")

	var choice int

	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if choice != 1 && choice != 2 {
		fmt.Println("Provide either 1 or 2")
		return
	}

	var temp float64
	fmt.Print("Enter the temperature value: ")
	_, err = fmt.Scanln(&temp)
	if err != nil {
		fmt.Println("Invalid temperature input")
		return
	}

	switch choice{
	case 1:
		result := celciusToFahrenheit(temp)
		fmt.Printf("%.2f degrees celcius = %.2f fahrenheit\n", temp, result)
	case 2:
		result := fahrenheitToCelcius(temp)
		fmt.Printf("%.2f fahrenheit = %.2f degrees celcius\n", temp, result)
	}
}
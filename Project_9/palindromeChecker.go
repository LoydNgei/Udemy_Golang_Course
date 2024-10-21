package main

import "fmt"

func reverseWord(wordToCheck string) string {
	reversed := ""
	for i := len(wordToCheck) - 1; i >= 0; i-- {
		reversed += string(wordToCheck[i])
	}
	return reversed
}

func main() {
	var wordToCheck string
	fmt.Print("Enter the word to check: ")
	fmt.Scanln(&wordToCheck)

	reversed := reverseWord(wordToCheck)
	if reversed == wordToCheck {
		fmt.Printf("%s is a palindrome\n", wordToCheck)
	} else {
		fmt.Printf("%s is not a palindrome\n", wordToCheck)
	}
}

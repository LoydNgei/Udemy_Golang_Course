package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findLongestWord(text string) string {
	longestWord := ""
	maxLength := 0

	words := strings.Fields(text)
	for _, word := range words {
		if len(word) > maxLength {
			longestWord = word
			maxLength = len(word)
		}
	}
	return longestWord
}

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence or a paragraph: ")
	text, _ = reader.ReadString('\n')

	longestWord := findLongestWord(text)
	fmt.Printf("The longest word is: %s\n", longestWord)

}

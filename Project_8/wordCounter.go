package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence or paragraph: ")
	word, _ = reader.ReadString('\n')

	words := strings.Fields(word)
	count := len(words)
	fmt.Printf("The number of words is: %d\n", count)
}

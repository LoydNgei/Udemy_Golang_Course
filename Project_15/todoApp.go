package main

import (
	"fmt"
)

func showTodo() {

}

func addTodo() {

}

func main() {
	for {
		var choice int
		fmt.Println("Welcome to ToDo application")

		fmt.Println("1. Add Todos")
		fmt.Println("2. Show Todos")
		fmt.Println("3. Exit")
		fmt.Print("Choose one: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case 1:
			showTodo()
		case 2:
			addTodo()
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Wrong input! Try again")
		}
		fmt.Println()
	}
}

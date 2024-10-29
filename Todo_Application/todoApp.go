package main

import (
	"fmt"
	"bufio"
	"os"
)

var todos []string


func addTodo() {
	reader := bufio.NewReader(os.Stdin)
	var task string
	fmt.Print("Enter the task description: ")
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	todos = append(todos, task)
	fmt.Println("Task added successfully")
}

func listToDos() {
	if len(todos) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for i, todo := range todos {
		fmt.Printf("%d. %s", i+1, todo)
	}
}

func main() {
	
	fmt.Println("Welcome to ToDo application")

	for {
		var choice int
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
			addTodo()
		case 2:
			listToDos()
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Wrong input! Try again")
		}
		fmt.Println()
	}
}

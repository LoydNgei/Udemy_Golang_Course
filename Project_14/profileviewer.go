package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func displayInfo(p Person) {
	fmt.Println("Profile Information")
	fmt.Printf("Name:   %s\n", p.Name)
	fmt.Printf("Age:    %d\n", p.Age)
	fmt.Printf("Email:  %s\n", p.Email)
	fmt.Println()
}

func main() {
	person1 := Person{Name: "Loyd", Age: 24, Email: "Loyd@example.com"}
	person2 := Person{Name: "Jane", Age: 20, Email: "Jane@example.com"}

	displayInfo(person1)
	displayInfo(person2)
}

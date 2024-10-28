package main

import (
	"fmt"
)

// Struct to hold the Name and phone number

type Contact struct {
	Name string
	Phone int
}

var contacts []Contact

// Add contact function

func addContact() {
	var name string
	var phone int

	fmt.Print("Enter the Contact Name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Enter the Contact Phone Number: ")
	_, err = fmt.Scanln(&phone)
	if err != nil {
		fmt.Println(err)
		return
	}

	contact := Contact{Name: name, Phone: phone}
	contacts = append(contacts, contact)
	fmt.Println("Contact added successfully")

}

// The function that lists the contacts

func listContacts() {
	if len(contacts) == 0 {
		fmt.Println("No contacts found")
		return
	}
	for i, contact := range contacts {
		fmt.Printf("%d. %s - %d\n", i+1, contact.Name, contact.Phone)
	}
}

func main() {

	fmt.Println("Contact list app")

	for {
		var choice int
		fmt.Println("1. Add contact")
		fmt.Println("2. List contact")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch choice {
		case 1:
			addContact()
		case 2:
			listContacts()
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Wrong choice! Try again")
		}
		fmt.Println()
	}


}
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var fname, lname string
// 	var YOB int
// 	fmt.Print("First Name: ")
// 	fmt.Scanln(&fname)
// 	fmt.Print("Last Name: ")
// 	fmt.Scanln(&lname)
// 	fmt.Print("Year of birth: ")
// 	fmt.Scanln(&YOB)

// 	ThisYear := time.Now().Year()

// 	ageNow := ThisYear - YOB
// 	fmt.Printf("Your name is %v %v and you are %d years old\n", fname, lname, ageNow)
// }
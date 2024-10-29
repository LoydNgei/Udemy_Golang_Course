package main

import (
    "fmt"
)

func main() {
    var height, weight float64

    fmt.Print("Weight in KGs: ")
    fmt.Scanln(&weight)
    fmt.Print("Height in M: ")
    fmt.Scanln(&height)

    bmi := weight / (height * height)
    fmt.Printf("Your BMI is %.2f\n", bmi)

    if bmi < 18.5 {
        fmt.Println("Category: Under weight")
    } else if bmi >= 18.5 && bmi <= 24.9 {
        fmt.Println("Category: Normal weight")
    } else if bmi >= 25 && bmi <= 29.9 {
        fmt.Println("Category: Over Weight")
    } else {
        fmt.Println("Category: Obese")
    }
}
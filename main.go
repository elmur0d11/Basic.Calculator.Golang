package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to GoCalculator")
	fmt.Println("Supported operations: +, -, *, /")
	
	var num1 float64
	var num2 float64
	var operator string

	//Get data from user
	fmt.Println("Enter number 1:")
	fmt.Scan(&num1)
	if num1 <= 0{
		fmt.Print("Number should bigger than 0")
		os.Exit(1109)
	}

	fmt.Println("Enter number 2:")
	fmt.Scan(&num2)
	if num2 <= 0{
		fmt.Print("Number should bigger than 0")
		os.Exit(1109)
	}

	fmt.Print("Enter operation:")
	fmt.Scan(&operator)
	
	//Checking
	operation(num1, num2, operator)
}

func operation(num1 float64, num2 float64, op string){
	switch(op){
	case "+":
		plus := num1 + num2
		fmt.Printf("%v + %v = %v", num1, num2, plus)

	case "-":
		minus := num1 - num2
		fmt.Printf("%v + %v = %v", num1, num2, minus)

	case "*":
		multiple := num1 * num2
		fmt.Printf("%v * %v = %v", num1, num2, multiple)

	case "/":
		subtract := num1 / num2
		fmt.Printf("%v / %v = %v", num1, num2, subtract)

	default:
		fmt.Printf("App can't run <%s> this operation, enter valid operations \n", op)
		fmt.Println("by elmur0d11 - v.1")
		os.Exit(9011)
	}
}
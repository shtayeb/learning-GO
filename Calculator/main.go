package main

import "fmt"

func main() {
	var operation string
	var number1 int
	var number2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("======================")

	fmt.Print("Which operation do you want to perform?: ")
	fmt.Scanf("%s",&operation)
	
	fmt.Println("Enter first number: ")
	fmt.Scanf("%d",&number1)

	fmt.Println("Enter second number:")
	fmt.Scanf("%d",&number2)

	switch operation {
		case "+":
			fmt.Println(number1 + number2)
		case "-":
			fmt.Println(number1 - number2)
		case "*":
			fmt.Println(number1 * number2)
		case "/":
			fmt.Println(number1 / number2)
	}
}
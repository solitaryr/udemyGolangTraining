/*
Create a program that prints to the terminal asking for a
user to enter a small number and a larger number. Print the
remainder of the larger number divided by the smaller number.
*/
package main

import "fmt"

func main() {
	var num1, num2, remainder int

	fmt.Print("Enter a small number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter a larger number: ")
	fmt.Scanln(&num2)

	remainder = num2 % num1
	fmt.Println("The remainder is", remainder)
}

/*
Write a function with one variadic parameter that finds the greatest number
in a list of positive numbers.
*/
package main

import "fmt"

func main() {
	numList := []int{1, 5, 53, 3, 10, 27, 12}
	fmt.Println("The largest number is", max(numList...))
}

func max(nums ...int) int {
	var largest int
	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}
	return largest
}

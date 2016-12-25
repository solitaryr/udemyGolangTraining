/*
Find a problem at projecteuler.net then create the solution. Add a
comment beneath your solution that includes a description of the
problem. Upload your solution to github. Tweet me a link to your
solution.

n! means n * (n − 1) * ... * 3 * 2 * 1

For example, 10! = 10 * 9 * ... * 3 * 2 * 1 = 3628800,
and the sum of the digits in the number 10! is

 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import "fmt"

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumDigits(n int64) int {
	// loop through each number in n and add
	var total int
	// num := strconv.Itoa(n)
	num := fmt.Sprintf("%d", n)
	fmt.Printf("%s - %T\n", num, num)
	for i := 0; i < len(num); i++ {
		total += int(num[i]) - '0'
	}
	return total
}

func main() {
	fmt.Println(factorial(25))
	fmt.Println(sumDigits(factorial(25)))
}

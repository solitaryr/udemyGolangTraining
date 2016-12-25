/*
Write a function which takes an integer. The function will have two returns.
The first return should be the argument divided by 2. The second return should
be a bool that let’s us know whether or not the argument was even.

For example:
  half(1) returns (0, false)
  half(2) returns (1, true)
*/
package main

import "fmt"

func half(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	h, even := half(5)
	fmt.Println(h, even)
}

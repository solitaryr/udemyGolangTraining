/*
Modify the previous program to use a func expression.
*/
package main

import "fmt"

func main() {
	half := func(x int) (float64, bool) {
		return float64(x) / 2, x%2 == 0
	}
	fmt.Printf("%T\n", half)
	fmt.Println(half(5))
}

package main

import "fmt"

func main() {
	// declare single
	var message string
	message = "Hello World!"

	// declare multiple
	var a, b, c int
	a = 1

	// initialize many at once
	var d, e, f int = 1, 2, 3

	// infer type
	var g, h, i = 3, false, '4'

	fmt.Println(message)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Printf("%v %T\n", g, g)
	fmt.Printf("%v %T\n", h, h)
	fmt.Printf("%v %T\n", i, i)

}

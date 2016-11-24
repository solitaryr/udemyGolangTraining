package main

import "fmt"

var o = "this is stored in the variable o"     // package scope
var p, q string = "stored in p", "stored in q" // package scope
var r string                                   // package scope

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

	// init shorthand
	// you can only do this inside a func
	message1 := "Hello World!"
	j, k, l := 1, false, 3
	m := 4
	n := true

	// All together now
	o = "stored in d" // declaration above; assignment here; package scope
	var s = 42        // function scope - subsequent variables have func scope:
	t := 43
	u := "stored in g"
	v, w := "stored in h", "stored in i"
	x, y, z, aa := 44.7, true, false, 'm' // single quotes
	ab := "n"                             // double quotes
	ac := `o`                             // back ticks

	fmt.Println(message)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Printf("%v %T\n", g, g)
	fmt.Printf("%v %T\n", h, h)
	fmt.Printf("%v %T\n", i, i)
	fmt.Println(message1)
	fmt.Println(j, k, l, m, n)

	fmt.Println("All together now!")
	fmt.Println("o - ", o)
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("r - ", r)
	fmt.Println("s - ", s)
	fmt.Println("t - ", t)
	fmt.Println("u - ", u)
	fmt.Println("v - ", v)
	fmt.Println("w - ", w)
	fmt.Println("x - ", x)
	fmt.Println("y - ", y)
	fmt.Println("z - ", z)
	fmt.Println("aa - ", aa)
	fmt.Println("ab - ", ab)
	fmt.Println("ac - ", ac)
}

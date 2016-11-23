package main

import "fmt"

func main() {
	fmt.Print("decimal: ")
	fmt.Println(42)

	fmt.Print("binary ")
	fmt.Printf("%b \n", 42)

	fmt.Print("hexadecimal: ")
	fmt.Printf("%x %#x %#X \n", 42, 42, 42)

	fmt.Println("loop: ")
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d\t%b\t%x\n", i, i, i)
	}
}

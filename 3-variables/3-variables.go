package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true // type inference
	fmt.Println(d)

	var e int // initialised with 0
	fmt.Println(e)

	f := "short" // := used to declare and initialise simutaneously
	fmt.Println(f)
}

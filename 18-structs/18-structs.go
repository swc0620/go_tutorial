package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Ann", age: 40}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // pointer is automatically dereferenced

	sp.age = 51
	fmt.Println(sp.age)
}

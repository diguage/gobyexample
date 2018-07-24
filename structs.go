package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"D瓜哥", 18})
	fmt.Println(person{name: "D瓜哥", age: 19})
	fmt.Println(person{name: "D瓜哥"})
	fmt.Println(&person{name: "D瓜哥", age: 20})

	s := person{name: "D瓜哥", age: 21}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 22
	fmt.Println(sp.age)
	fmt.Println(s.age)
}

package main

import "fmt"

func main() {
	if 119%2 == 0 {
		fmt.Println("119 is even")
	} else {
		fmt.Println("119 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has many digit")
	}
}

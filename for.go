package main

import "fmt"

func main() {
	fmt.Println("Basic type:")
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\nTristage")
	for j := 7; j < 10; j++ {
		if j%3 == 0 {
			fmt.Println(j)
		}
	}

	fmt.Println("\nSimplest")
	for {
		fmt.Println("Welcome to https://www.diguage.com/")
		break
	}

	fmt.Println("\nCommon")
	for n := 4; n < 10; n++ {
		if n%2 == 0 {
			fmt.Println(n)
		}
	}
}

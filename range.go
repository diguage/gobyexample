package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for i, num := range nums {
		fmt.Printf("%T: %d\n", i, i)
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("idx:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Printf("key: %s\n", k)
	}

	for i, c := range "diguage" {
		fmt.Println(i, c)
	}
	for i, c := range "D瓜哥" {
		fmt.Println(i, c)
	}
}

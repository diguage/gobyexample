package main

import "fmt"
import "sort"

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}

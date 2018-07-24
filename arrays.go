package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
	// fmt.Println("len", a.Length()) ❌

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("del:", b)

    fmt.Println("sl1:", b[:])
    fmt.Println("sl2:", b[2:])
    fmt.Println("sl3:", b[:2])

    // https://stackoverflow.com/a/27160765/951836
    fmt.Printf("art:%T\n", b)
    fmt.Printf("slt:%T\n", b[:2])
	// c := append(b, 6, 7)
	// fmt.Println("apd:", b, c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(" 2d:", twoD)
}

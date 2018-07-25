package main

import "fmt"
import "math/rand"
import "time"

func main() {
	i := 2
	ri := rand.Intn(5)
	r31i := rand.Int31n(5)
	r63i := rand.Int63n(5)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", ri)
	fmt.Printf("%T\n", r31i)
	fmt.Printf("%T\n", r63i)
	fmt.Printf("%T\n", 2*time.Second)
	//fmt.Printf("%T\n", i * time.Second)
	//fmt.Printf("%T\n", ri * time.Second)
	//fmt.Printf("%T\n", r31i * time.Second)
	//fmt.Printf("%T\n", r63i * time.Second)
}

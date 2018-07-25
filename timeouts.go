package main

import "time"
import "fmt"
import "math/rand"

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		c1 <- "result 1"
	}()

	select {
	case re := <-c1:
		fmt.Println(re)
	case <-time.After(time.Duration(rand.Intn(2)) * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		c2 <- "result 2"
	}()

	select {
	case re := <-c2:
		fmt.Println(re)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

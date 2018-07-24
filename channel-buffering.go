package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "D瓜哥"
	messages <- "https://www.diguage.com/"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

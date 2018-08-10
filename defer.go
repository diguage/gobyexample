package main

import "fmt"
import "os"

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing…")
	fmt.Fprintln(f, "Hello, D瓜哥！")
}

func closeFile(f *os.File) {
	fmt.Println("Closing…")
	f.Close()
}

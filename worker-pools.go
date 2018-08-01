package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker id", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker id", id, "finished job", job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 10; i <= 50; i += 10 {
		jobs <- i
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Println("result", result)
	}

	// // fatal error: all goroutines are asleep - deadlock!
	// for result := range results {
	//     fmt.Println("result", result)
	// }
}

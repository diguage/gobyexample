package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(5 * time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

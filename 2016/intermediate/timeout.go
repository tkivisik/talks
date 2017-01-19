package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {

	result := make(chan int)

	go func() {
		time.Sleep(500 * time.Millisecond)
		result <- 42
	}()

	for {
		select {
		case value := <-result:
			fmt.Printf("Computed the value in time :%v", value)
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Printf("Timed out!")
			return
		}
	}
}

// END OMIT

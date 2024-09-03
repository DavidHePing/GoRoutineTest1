package main

import (
	"fmt"
	"sync"
	"time"
)

func Deadlock_test1_deadlock() {
	var wg sync.WaitGroup
	ch := make(chan int) // Create a channel to share data

	// Goroutine to increment counter
	go func() {
		defer wg.Done()
		counter := 0
		for i := 0; i < 5; i++ {
			counter++
			ch <- counter // Send the updated counter value through the channel
			time.Sleep(1 * time.Second)
		}
		close(ch) // Close the channel when done
	}()

	wg.Add(1)
	wg.Wait()

	fmt.Println("Finished")
}

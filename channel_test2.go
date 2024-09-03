package main

import (
	"fmt"
	"sync"
	"time"
)

func channel_test2_test1() {
	ch := make(chan int) // Create a channel to share data

	// Goroutine to increment counter
	go func() {
		counter := 0
		for i := 0; i < 5; i++ {
			counter++
			ch <- counter // Send the updated counter value through the channel
			fmt.Println("put in", counter)
			time.Sleep(1 * time.Second)
		}
		close(ch) // Close the channel when done
	}()

	for value := range ch {
		fmt.Println("Counter value:", value)
	}

	fmt.Println("Finished")
}

func channel_test2_test2() {
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

	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println("Counter value:", value)
		}
	}()

	wg.Add(2)
	wg.Wait()

	fmt.Println("Finished")
}

func channel_test2_test3() {
	ch := make(chan int) // Create a channel to share data

	// Goroutine to increment counter
	go func() {
		counter := 0
		for i := 0; i < 5; i++ {
			counter++
			ch <- counter // Send the updated counter value through the channel
			fmt.Println("put in", counter)
		}
		close(ch) // Close the channel when done
	}()

	for value := range ch {
		fmt.Println("Counter value:", value)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Finished")
}

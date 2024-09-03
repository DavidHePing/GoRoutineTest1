package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Atomic_test1() {
	var counter int32  // Use an int32 for atomic operations
	var counter2 int32 // Use an int32 for atomic operations
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			atomic.AddInt32(&counter, 1)
			// fmt.Println("Goroutine1 counter:", counter)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			atomic.AddInt32(&counter, 1)
			// fmt.Println("Goroutine2 counter:", counter)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			counter2++
			fmt.Println("Goroutine1 counter2:", counter2)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			counter2++
			fmt.Println("Goroutine2 counter2:", counter2)
		}
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter)
	fmt.Println("Final counter2 value:", counter2)

}

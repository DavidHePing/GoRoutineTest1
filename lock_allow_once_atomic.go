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
			fmt.Println("Goroutine1 counter:", counter)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			counter++
			// atomic.AddInt32(&counter, 1)
			fmt.Println("Goroutine2 counter:", counter)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			counter2++
			// fmt.Println("Goroutine1 counter2:", counter2)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			counter2++
			// fmt.Println("Goroutine2 counter2:", counter2)
		}
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter)
	fmt.Println("Final counter2 value:", counter2)

}

func Atomic_test2() {
	var val1 int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&val1, 1)
		}()
	}
	wg.Wait()
	fmt.Println("val1:", atomic.LoadInt32(&val1)) //must be 1000

	var val2 int32
	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			val2++
		}()
	}
	wg2.Wait()
	fmt.Println("val2:", atomic.LoadInt32(&val2)) //may be not 1000
}

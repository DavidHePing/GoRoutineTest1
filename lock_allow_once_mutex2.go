package main

import (
	"fmt"
	"sync"
	"time"
)

func Lock_by_mutex_lock_2() {
	var lock sync.Mutex
	val := 0
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()

		for i := 0; i < 5; i++ {
			val++
			fmt.Println("Goroutine 1:", val)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			val++
			fmt.Println("Goroutine 2:", val)
			time.Sleep(time.Millisecond * 100)
		}
	}()
	wg.Add(2)
	wg.Wait()

	fmt.Printf("final val = %d\n", val)
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func MutexLockTest1_without_lock() {
	var wg sync.WaitGroup
	val := 0

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			val++
			fmt.Printf("First gorutine val++ and val = %d\n", val)
			time.Sleep(3000)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			val++
			fmt.Printf("Sec gorutine val++ and val = %d\n", val)
			time.Sleep(1000)
		}
	}()
	wg.Add(2)
	wg.Wait()
}

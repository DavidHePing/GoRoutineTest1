package main

import (
	"fmt"
	"sync"
	"time"
)

func Lock_Rw_Test2() {
	var val int32
	var wg sync.WaitGroup
	var lock sync.RWMutex

	wg.Add(2)
	go func() {
		defer wg.Done()
		lock.Lock()
		for i := 0; i < 10; i++ {
			val++
			time.Sleep(1 * time.Second)
		}
		lock.Unlock()
	}()

	go func() {
		time.Sleep(10 * time.Millisecond)
		defer wg.Done()
		lock.Lock()
		for i := 0; i < 10; i++ {
			fmt.Println("i:", val)
		}
		val++
		fmt.Println("after:", val)
		lock.Unlock()
	}()

	wg.Wait()
	fmt.Println("val:", val)
}

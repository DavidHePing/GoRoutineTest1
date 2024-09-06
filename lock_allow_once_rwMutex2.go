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
		lock.RLock()
		for i := 0; i < 10; i++ {
			val++
			time.Sleep(1 * time.Second)
		}
		lock.RUnlock()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		defer wg.Done()
		lock.RLock()
		for i := 0; i < 10; i++ {
			fmt.Println("i:", val)
		}
		val++
		fmt.Println("after:", val)
		lock.RUnlock()
	}()

	wg.Wait()
	fmt.Println("final val:", val)
}

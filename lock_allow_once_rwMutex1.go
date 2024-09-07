package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Lock_Rw_Test1() {
	var val1 int32
	var wg sync.WaitGroup
	var lock sync.RWMutex
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			lock.RLock()
			val1++
			lock.RUnlock()
		}()
	}
	wg.Wait()
	fmt.Println("val1:", atomic.LoadInt32(&val1)) //may be not 1000

	var val2 int32
	var wg2 sync.WaitGroup
	wg2.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg2.Done()
			val2++
		}()
	}
	wg2.Wait()
	fmt.Println("val2:", atomic.LoadInt32(&val2)) //may be not 1000
}

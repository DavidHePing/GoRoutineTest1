package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Lock_by_mutex_lock_3() {
	var val1 int32
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			val1++
			lock.Unlock()
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

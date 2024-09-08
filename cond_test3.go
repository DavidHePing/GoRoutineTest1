package main

import (
	"sync"
)

func Cond_Test3_Broadcast() {
	var wg sync.WaitGroup
	wg.Add(2)
	var wgWaitConditionLock sync.WaitGroup
	wgWaitConditionLock.Add(2)

	var cond = sync.NewCond(&sync.Mutex{})
	go func() {
		defer wg.Done()
		cond.L.Lock()
		wgWaitConditionLock.Done() //after lock, make sure it would entry to wait
		cond.Wait()
		println("Goroutine1!")
		cond.L.Unlock()
	}()

	go func() {
		defer wg.Done()
		cond.L.Lock()
		wgWaitConditionLock.Done() //after lock, make sure it would entry to wait
		cond.Wait()
		println("Goroutine2!")
		cond.L.Unlock()
	}()

	wgWaitConditionLock.Wait()

	cond.L.Lock()
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
}

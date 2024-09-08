package main

import (
	"sync"
	"time"
)

func Cond_Test1(isWaitForAllGoroutineSetup bool) {
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

	if isWaitForAllGoroutineSetup {
		//Ensure that both goroutines to start waiting, if not, would deadlock
		wgWaitConditionLock.Wait()
	}

	cond.L.Lock()
	//might be step 1 if no wait, if signal -> lock -> lock -> signal, deadlock would happen
	cond.Signal()
	cond.L.Unlock()
	time.Sleep(1 * time.Second)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()
	wg.Wait()
}

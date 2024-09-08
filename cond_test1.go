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
		wgWaitConditionLock.Done()
		cond.Wait()
		println("Goroutine1!")
		cond.L.Unlock()
	}()

	go func() {
		defer wg.Done()
		cond.L.Lock()
		wgWaitConditionLock.Done()
		cond.Wait()
		println("Goroutine2!")
		cond.L.Unlock()
	}()

	//still deadlock if signal is executed before any goroutine's cond.Wait
	wgWaitConditionLock.Wait()

	if isWaitForAllGoroutineSetup {
		//Ensure that both goroutines to start waiting, if not, would deadlock
		time.Sleep(10 * time.Millisecond)
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

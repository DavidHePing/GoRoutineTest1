package main

import (
	"sync"
	"time"
)

func Cond_Test3_Broadcast() {
	var wg sync.WaitGroup
	wg.Add(2)

	var cond = sync.NewCond(&sync.Mutex{})
	go func() {
		defer wg.Done()
		cond.L.Lock()
		cond.Wait()
		println("Goroutine1!")
		cond.L.Unlock()
	}()

	go func() {
		defer wg.Done()
		cond.L.Lock()
		cond.Wait()
		println("Goroutine2!")
		cond.L.Unlock()
	}()

	time.Sleep(1 * time.Second)

	cond.L.Lock()
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
}

package main

import (
	"sync"
	"time"
)

func Cond_Test2_Fetal() {
	var wg sync.WaitGroup
	wg.Add(2)
	var cond = sync.NewCond(&sync.Mutex{})
	go func() {
		defer wg.Done()
		//need lock before wait
		cond.Wait()
		println("Goroutine1!")
	}()

	go func() {
		defer wg.Done()
		//need lock before wait
		cond.Wait()
		println("Goroutine2!")
	}()

	//Ensure that both goroutines have time to start waiting, if not, would deadlock
	time.Sleep(1 * time.Second)

	//need lock before wait
	cond.Signal()
	time.Sleep(1 * time.Second)
	//need lock before wait
	cond.Signal()
	wg.Wait()
}

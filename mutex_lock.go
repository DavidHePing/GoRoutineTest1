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

func MutexLockTest1_lock_1of_useless() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	val := 0

	go func() {
		defer wg.Done()

		lock.Lock()
		for i := 0; i < 5; i++ {
			val++
			fmt.Printf("First gorutine val++ and val = %d\n", val)
			time.Sleep(3000)
		}
		lock.Unlock()
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

func MutexLockTest1_with_lock() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	val := 0

	go func() {
		defer wg.Done()

		lock.Lock()
		for i := 0; i < 5; i++ {
			val++
			fmt.Printf("First gorutine val++ and val = %d\n", val)
			time.Sleep(3000)
		}
		lock.Unlock()
	}()

	go func() {
		defer wg.Done()

		lock.Lock()
		for i := 0; i < 5; i++ {
			val++
			fmt.Printf("Sec gorutine val++ and val = %d\n", val)
			time.Sleep(1000)
		}
		lock.Unlock()
	}()
	wg.Add(2)
	wg.Wait()
}

func MutexLockTest1_with_2lock_useless() {
	var lock sync.Mutex
	var lock2 sync.Mutex
	var wg sync.WaitGroup
	val := 0

	go func() {
		defer wg.Done()

		lock.Lock() // Acquire the lock
		for i := 0; i < 5; i++ {
			val++
			fmt.Printf("First gorutine val++ and val = %d\n", val)
			time.Sleep(3000)
		}
		lock.Unlock()
	}()

	go func() {
		defer wg.Done()

		lock2.Lock() // Acquire the lock
		for i := 0; i < 5; i++ {
			val++
			fmt.Printf("Sec gorutine val++ and val = %d\n", val)
			time.Sleep(1000)
		}
		lock2.Unlock()
	}()
	wg.Add(2)
	wg.Wait()
}

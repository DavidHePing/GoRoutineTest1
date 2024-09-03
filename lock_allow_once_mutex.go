package main

import (
	"fmt"
	"sync"
	"time"
)

func show1(str string, lock *sync.Mutex, val *int) {
	lock.Lock()
	defer lock.Unlock()

	for i := 0; i < 5; i++ {
		*val++
		fmt.Println(str, "gorutine val++ and val =", *val)
		time.Sleep(1000)
	}
}

func Lock_by_mutex_lock_1() {
	var lock sync.Mutex
	val := 0
	var wg sync.WaitGroup
	defer wg.Wait()

	go func() {
		defer wg.Done()
		show1("First", &lock, &val)
	}()

	go func() {
		defer wg.Done()

		show1("Sec", &lock, &val)
	}()
	wg.Add(2)
	fmt.Printf("final val = %d\n", val)
}

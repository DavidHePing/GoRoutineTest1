package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock sync.Mutex
	val  = 0
)

func show(str string) {
	lock.Lock()
	defer lock.Unlock()

	for i := 0; i < 5; i++ {
		val++
		fmt.Println(str, "gorutine val++ and val =", val)
		time.Sleep(1000)
	}

}

func Lock_by_mutex_lock() {
	var wg sync.WaitGroup
	defer wg.Wait()
	val := 0

	go func() {
		defer wg.Done()
		show("First")
	}()

	go func() {
		defer wg.Done()

		show("Sec")
	}()
	wg.Add(2)
	fmt.Printf("final val = %d\n", val)
}

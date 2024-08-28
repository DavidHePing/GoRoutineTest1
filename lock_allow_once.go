package main

import (
	"fmt"
	"sync"
	"time"
)

func Lock_by_once() {
	addVal := func(val *int, num int, goroutine string) {
		var onceLock sync.Once
		onceLock.Do(func() {
			*val += num
			fmt.Println(goroutine, "gorutine val++ and val =", *val)
		})
	}

	var wg sync.WaitGroup
	val := 0

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			addVal(&val, 1, "First")
			time.Sleep(3000)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			addVal(&val, 1, "Sec")
			time.Sleep(1000)
		}
	}()
	wg.Add(2)
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func show2(str string, once *sync.Once, val *int) {
	//only do once, if First did, Second will not do
	once.Do(func() {
		for i := 0; i < 5; i++ {
			*val++
			fmt.Println(str, "gorutine val++ and val =", *val)
			time.Sleep(1000)
		}
	})

}

func Lock_by_syncOnce() {
	var once sync.Once
	val := 0
	var wg sync.WaitGroup
	defer wg.Wait()

	go func() {
		defer wg.Done()
		show2("First", &once, &val)
	}()

	go func() {
		defer wg.Done()

		show2("Sec", &once, &val)
	}()
	wg.Add(2)
	fmt.Printf("final val = %d\n", val)
}

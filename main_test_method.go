package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ReleaseChildWhenParentRelase() {
	go func() {
		time.Sleep(100000000)
		fmt.Println("goroutine Done!")
	}()
	fmt.Println("Done!")

}

func WaitGoRoutine() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer fmt.Println("goroutine1 drop out")
		fmt.Println("start a goroutine1")
		time.Sleep(time.Second)
	}()

	go func() {
		defer wg.Done() //wg counter -1
		defer fmt.Println("goroutine2 drop out")
		fmt.Println("start a goroutine2")
		time.Sleep(time.Second)
	}()

	fmt.Println("wait a goroutine")
	wg.Wait()
}

func SelectTest() {

	firstRoutine := make(chan string)
	secRoutine := make(chan string)

	go func() {
		r := rand.Intn(100)
		time.Sleep(time.Microsecond * time.Duration(r))
		firstRoutine <- "first goroutine"
	}()

	go func() {
		r := rand.Intn(100)
		time.Sleep(time.Microsecond * time.Duration(r))
		secRoutine <- "Second goroutine"
	}()

	select {
	case f := <-firstRoutine:
		fmt.Println("first: ", f)
	case s := <-secRoutine:
		fmt.Println("second: ", s)
	}
}

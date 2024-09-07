package main

import (
	"fmt"
)

func Deadlock_test1_deadlock() {
	ch1 := make(chan int)
	//deadlock cause no sender
	a := <-ch1
	fmt.Println(a)
}

func Deadlock_test2_deadlock(isRun bool) {
	defer fmt.Println("End!!!")

	fmt.Println("Start!!! isRun:", isRun)
	ch1 := make(chan int)

	if isRun {
		go func() {
			result := 1 + 2
			ch1 <- result
		}()
	} else {
		//dead lock cause non bufferd channel should set up sender and reciever in first time
		ch1 <- 1 + 2
	}

	result := <-ch1
	fmt.Println(result)
}

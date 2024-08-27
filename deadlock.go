package main

import (
	"fmt"
)

func DeadLockTest1() {
	ch1 := make(chan int)
	a := <-ch1
	fmt.Println(a)
}

func DeadLockTest2(isRun bool) {
	defer fmt.Println("End!!!")

	fmt.Println("Start!!! isRun:", isRun)
	ch1 := make(chan int)

	if isRun {
		go func() {
			result := 1 + 2
			ch1 <- result
		}()
	} else {
		// would dead lock
		ch1 <- 1 + 2
	}

	result := <-ch1
	fmt.Println(result)
}

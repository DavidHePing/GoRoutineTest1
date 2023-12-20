package main

import "fmt"

func DeadLockTest1() {
	ch1 := make(chan int)
	a := <-ch1
	fmt.Println(a)
}

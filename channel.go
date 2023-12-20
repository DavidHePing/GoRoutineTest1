package main

import (
	"fmt"
	"time"
)

func MockApi() {
	ch1 := DoneAsync(1, 2)
	ch2 := DoneAsync(2, 1)

	println(<-ch1)
	println(<-ch2)
}

func DoneAsync(val int, waitTime int64) chan int {
	r := make(chan int)
	fmt.Println("Warming up ...", val)
	go func() {
		time.Sleep(time.Duration(waitTime) * time.Second)
		fmt.Println("Done ...", val)
		r <- val
	}()
	return r
}

func UseChannelGetValue() {
	val := make(chan int)

	go func() {
		fmt.Println("intput val 1")
		val <- 1
	}()

	//First do
	go func() {
		fmt.Println("intput val")
		val <- 4
		val <- 3
		val <- 2
		time.Sleep(time.Second * 2)
	}()

	ans := []int{}
	for {
		ans = append(ans, <-val)
		fmt.Println(ans)
		if len(ans) == 4 {
			fmt.Println("Done")
			break
		}
	}
}

func WaitGoRoutineUseChannel() {
	forever := make(chan int)

	go func() {
		defer fmt.Println("goroutine channel 1 drop out")
		fmt.Println("start a goroutine channel 1")
		time.Sleep(time.Second)
		forever <- 1
	}()

	go func() {
		defer fmt.Println("goroutine channel 2 drop out")
		fmt.Println("start a goroutine channel 2")
		time.Sleep(time.Second)
		forever <- 2
	}()

	fmt.Println("wait a goroutine")

	<-forever
	<-forever

	// a := <-forever
	// fmt.Println(a)
	// b := <-forever
	// fmt.Println(b)

}

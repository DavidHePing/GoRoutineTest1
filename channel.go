package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MockApi() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go getValueAsync(ch1, 1)
	go getValueAsync(ch2, 2)

	println(<-ch1)
	println(<-ch2)
}

func getValueAsync(channel chan int, val int) {
	go func() {
		r := rand.Intn(100)
		time.Sleep(time.Microsecond * time.Duration(r))
		channel <- val
	}()
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

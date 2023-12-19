package main

import (
	"fmt"
	"time"
)

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

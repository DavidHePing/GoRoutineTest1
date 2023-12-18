package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// ReleaseChildWhenParentRelase()
	// UseChannelGetValue()
	// WaitGoRoutine()
	// WaitGoRoutineUseChannel()
	// LockTest()
	SelectTest()
}

func ReleaseChildWhenParentRelase() {
	go func() {
		time.Sleep(100000000)
		fmt.Println("goroutine Done!")
	}()
	fmt.Println("Done!")

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

func LockTest() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	val := 0

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			lock.Lock()
			val++
			fmt.Printf("First gorutine val++ and val = %d\n", val)
			lock.Unlock()
			time.Sleep(3000)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			lock.Lock()
			val++
			fmt.Printf("Sec gorutine val++ and val = %d\n", val)
			lock.Unlock()
			time.Sleep(1000)
		}
	}()
	wg.Add(2)
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

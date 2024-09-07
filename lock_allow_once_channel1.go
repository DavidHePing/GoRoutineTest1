package main

import "sync"

func Lock_by_channel_1() {
	var ch = make(chan int, 1)
	ch <- 0
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			val := <-ch
			ch <- val + 1
		}()
	}

	wg.Wait()
	println(<-ch) // must be 1000

	var ch2 = make(chan int, 1)
	ch2 <- 0

	for i := 0; i < 1000; i++ {
		go func() {
			val := <-ch2
			ch2 <- val + 1
		}()
	}

	println(<-ch2) // may not be 1000, just cause no wait
}

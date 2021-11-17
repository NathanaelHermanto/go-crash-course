package main

import (
	"fmt"
	"time"
)

func main() {
	/*var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("cats")
		wg.Done()
	}()

	wg.Wait()*/

	//go count("sheep")
	//count("cats") // go count("cats") => program exited when go routine finishes
	///////////////////////////////////////////////////////
	//channel
	c := make(chan string)
	go count("cats", c)

	/*for {
		msg, open := <-c // *receive message from the channel
		if open {
			fmt.Println(msg)
		} else {
			break
		}
	}*/
	//better shorter version of the above
	for msg := range c {
		fmt.Println(msg)
	}
	//////////////////////////////////////////////////////////
	//setting cap for the chanel (buffered channe)
	c = make(chan string, 2) // 2 cap for the channel
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
	/////////////////////////////////////////////////////////
	// select statement
	/*c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "every 2000ms"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}*/
	///////////////////////////////////////////////////////
	// workerpool pattern
	/*jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}*/
	// fibo nums will probably not come out in order
}

func count(s string, c chan string) {
	//endless loop
	for i := 0; i <= 5; i++ {
		c <- s // send the message through the channel*
		time.Sleep(time.Millisecond * 500)
	}
	close(c) // close channel when finished
}

/*
func fibo(n int) int {
	if n <= 1 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}

func worker(jobs <-chan int, results chan<- int) { // receive jobs send results
	for n := range jobs {
		results <- fibo(n)
	}
}*/

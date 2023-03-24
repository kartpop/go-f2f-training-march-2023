// Wait for Result Pattern and fanout
// Usecase --> function fanoutSemaphore: when producing messages, keep count of goroutines in sync with logical/actual no. of processors
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	waitForResult()
	// /fanOut()
	//fanOutSemaphore()

}

func waitForResult() {

	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // just for latency
		//ch <- "I completed my task"
	}()

	fmt.Println("Wait for task completion message: ", <-ch)
}

func fanOut() {

	noOftasks := 5
	task := make(chan string, noOftasks)

	for i := 0; i < noOftasks; i++ {
		// goroutine count = no of task 1:1
		// if count more then block as buffer size full
		// if less deadlock on reciever
		go func(n int) {
			time.Sleep(time.Second)
			task <- fmt.Sprintf("%d task completed", n)
		}(i)
	}

	for noOftasks > 0 {
		fmt.Println(<-task)
		noOftasks--
	}
}

func fanOutSemaphore() {

	noOftasks := 10
	task := make(chan string, noOftasks)

	//g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, 5)
	for i := 0; i < noOftasks; i++ {
		sem <- true
		go func(n int) {
			{
				time.Sleep(time.Second)
				task <- fmt.Sprintf("%d task completed", n)
				fmt.Printf("No of running goroutines are: %d\n", runtime.NumGoroutine())
			}
			<-sem
		}(i)
	}

	for noOftasks > 0 {
		fmt.Println(<-task)
		noOftasks--
	}
}

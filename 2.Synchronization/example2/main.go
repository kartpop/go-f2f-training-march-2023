// Wait for Task Pattern and pooling

// Parent Goroutine - Sender which is sending task to child
// Child Goroutie - Reciever which is recieving/ waiting for task

// Retrive data from Channel for with range and without range

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

	// waitForTask()
	// pooling()
	// poolingWithoutRange()
	boundedWorkPooling()
}

func waitForTask() {

	ch := make(chan string)

	go func() {
		fmt.Println("Recived task", <-ch)
	}()
	ch <- "pass bricks"
	time.Sleep(time.Duration(rand.Intn(500)))
}

// uses wait for task pattern
// important for efficiency in resource usage
func pooling() {

	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for i := 0; i < g; i++ {
		go func(n int) {
			for t := range ch {
				fmt.Printf("Gohper: %d Task Recieved : %s\n", n, t)
			}
		}(i)
	}

	const task = 10
	for t := 0; t < task; t++ {
		ch <- fmt.Sprint("task ", t)
		// fmt.Println("Task sent")
	}
	close(ch)
	time.Sleep(time.Second)

}

// uses wait for task pattern
// important for efficiency in resource usage
func poolingWithoutRange() {

	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for i := 0; i < g; i++ {
		go func(n int) {
			for {
				t, withData := <-ch
				if !withData {
					break
				}
				fmt.Printf("Gohper: %d Task Recieved : %s\n", n, t)
			}
		}(i)
	}

	const task = 10
	for t := 0; t < task; t++ {
		ch <- fmt.Sprint("task ", t)
		// fmt.Println("Task sent")
	}
	close(ch)
	time.Sleep(time.Second)

}

// Bounded work pooling
func boundedWorkPooling() {

	ch := make(chan string)

	const task = 10

	var wg sync.WaitGroup

	g := runtime.GOMAXPROCS(0)
	wg.Add(g)
	for i := 0; i < g; i++ {
		go func(n int) {
			defer wg.Done()
			for t := range ch {
				fmt.Printf("Gohper: %d Task Recieved : %s\n", n, t)
			}
		}(i)
	}

	for t := 0; t < task; t++ {
		ch <- fmt.Sprint("task ", t)
	}
	close(ch)
	wg.Wait()
	//	time.Sleep(time.Second) replace time.Sleep with Waitgroup

}

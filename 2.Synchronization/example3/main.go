// Drop pattern
// useful for services which may experience heavy load at times.
// So using this pattern for services can drop requests when the services reaches a capacity of pending requests.
// Example . DNS service- useful in case of DDOS attack.

//Buffered channel

// Select statement :
// Is blocking call that allows the parent goroutine to handle mutliple channel operations at the same time.
// Each case in select represents send or a recive .

// Also default keyword which turns select into non-blocking call.
// When the channel buffer is full , that will cause case statement to block since the send cant complete
// When every case select case is blocked then default statement is executed.
// we can write handling for those request for which our system is not able to handle.

package main

import "fmt"

func main() {

	drop()

}

func drop() {

	// server capacity (we need load testing/perfomance testing)
	const cap = 10

	ch := make(chan string, cap)

	// just for demo single gorotine will create back pressure on system
	go func() {
		for r := range ch {
			fmt.Printf("Server: recieved %s request\n", r)
		}
	}()

	const workload = 50
	for w := 0; w < workload; w++ {
		select {
		case ch <- "requestData":
			fmt.Printf("Parent: sending %d request\n", w)
		default:
			fmt.Printf("Parent: dropped request %d request\n", w)
		}
	}

	close(ch)
}

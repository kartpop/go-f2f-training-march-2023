// Cancellation Pattern

// used to tell a function performing some I/O / any operation how long am willing to wait for the operation to complete
// if timeout then cancel operation or just walk away

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	cancellation()
}

func cancellation() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // imp if forget then memory leak

	ch := make(chan string, 1) // why buffer 1????
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "data"
	}()

	select {
	case d := <-ch:
		fmt.Println("Parent: processed request", d)
	case <-ctx.Done():
		fmt.Println("Parent: bye bye ")
	}

}

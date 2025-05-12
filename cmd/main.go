package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create two channels, one for even numbers and one for odd numbers
	evenCh := make(chan int)
	oddCh := make(chan int)

	var wg sync.WaitGroup

	// Go-Routine to send even numbers up to 30
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				evenCh <- i
			}
		}
		close(evenCh)
	}()

	// Go-Routine to send odd numbers up to 30
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				oddCh <- i
			}
		}
		close(oddCh)
	}()

	// Alternately print even and odd numbers until both channels are closed
	go func() {
		for {
			// This loop keeps running indefinitely (for {}), alternating between printing even and odd numbers, until both channels are closed and drained.
			odd, okOdd := <-oddCh
			if okOdd {
				fmt.Println("Odd:", odd)
			}
			even, okEven := <-evenCh
			if okEven {
				fmt.Println("Even:", even)
			}
			if !okOdd && !okEven {
				break
			}
		}
	}()

	// Wait for all Go-Routines to finish
	wg.Wait()
}

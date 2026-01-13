// package main

// import (
// 	"fmt"
// 	"os"
// 	"sync"
// 	"time"
// )

// // Writer goroutine: sends log messages into channel
// func writer(id int, ch chan<- string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		msg := fmt.Sprintf("Writer %d: log entry %d\n", id, i)
// 		ch <- msg
// 		time.Sleep(100 * time.Millisecond) // simulate delay
// 	}
// }

// // File writer goroutine: consumes channel and writes to file
// func fileWriter(file *os.File, ch <-chan string, done chan<- struct{}) {
// 	for msg := range ch {
// 		_, err := file.WriteString(msg)
// 		if err != nil {
// 			fmt.Println("Write error:", err)
// 		}
// 	}
// 	// signal completion
// 	done <- struct{}{}
// }

// func main() {
// 	// Open file for append
// 	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// Channels
// 	logCh := make(chan string, 100) // buffered channel
// 	done := make(chan struct{})

// 	// Start file writer goroutine
// 	go fileWriter(file, logCh, done)

// 	var wg sync.WaitGroup

// 	// Start multiple writers
// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go writer(i, logCh, &wg)
// 	}

// 	// Wait for writers to finish
// 	wg.Wait()
// 	close(logCh) // close channel so fileWriter exits

// 	// Wait for fileWriter to finish
// 	<-done

// 	fmt.Println("All logs written to file.")
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Producer: reads file and sends lines into channel
func fileReader(path string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Open error:", err)
		close(ch)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text() // push line into channel
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scan error:", err)
	}
	close(ch) // signal no more data
}

// Consumer: pulls lines from channel and processes them
func consumer(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range ch {
		fmt.Printf("Consumer %d got: %s\n", id, line)
	}
}

func main() {
	linesCh := make(chan string, 100) // buffered channel
	var wg sync.WaitGroup

	// Start reader
	wg.Add(1)
	go fileReader("server.log", linesCh, &wg)

	// Start consumers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go consumer(i, linesCh, &wg)
	}

	wg.Wait()
	fmt.Println("Reading complete.")
}

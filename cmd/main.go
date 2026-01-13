package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Writer: Appends logs with a simulated delay
func writer(file *os.File, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 200; i++ { // Reduced count for demonstration
		_, err := file.WriteString(fmt.Sprintf("Writer %d: log entry %d\n", id, i))
		if err != nil {
			fmt.Printf("Writer %d error: %v\n", id, err)
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// Unlike the tailing reader, which continuously monitors new log entries, the snapshot reader captures the entire file state at a single point in time.
// Snapshot Reader: Loads the current state of the file into memory
func snapshotReader(name string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)

	// os.ReadFile opens, reads, and closes the file automatically
	data, err := os.ReadFile("server.log") // After the delay, it uses os.ReadFile to load the entire contents of server.log into memory.
	if err != nil {
		fmt.Printf("Reader %s error: %v\n", name, err)
		return
	}
	fmt.Printf("--- Snapshot %s ---\n%s\n-------------------\n", name, string(data))
}

func main() {
	// Initialize the file
	logFile := "server.log"
	writeFile, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer writeFile.Close()

	var wg sync.WaitGroup
	/*
		Three writer goroutines are launched:

		    Each is added to the WaitGroup.

		    They run concurrently, appending logs to the same file.

		This simulates multiple processes writing logs simultaneously.

	*/
	// Start 3 Writers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go writer(writeFile, i, &wg)
	}
	/*
		Four snapshot readers are launched:

		    Two short-term readers with delays of 100ms and 300ms.

		    Two long-term readers with delays of 5s and 15s.

		    Each is added to the WaitGroup.

		This means snapshots are taken at different points in time, showing how the file grows as writers continue to append logs.
	*/
	// Start 2 Snapshot Readers at different intervals
	wg.Add(1)
	go snapshotReader("Short-Term", 100*time.Millisecond, &wg)
	wg.Add(1)
	go snapshotReader("Long-Term", 300*time.Millisecond, &wg)
	// Start 2 Snapshot Readers at different intervals
	wg.Add(1)
	go snapshotReader("Short-Term", 5*time.Second, &wg)
	wg.Add(1)
	go snapshotReader("Long-Term", 15*time.Second, &wg)
	/*
		    The program waits for all goroutines (wg.Wait()).

		    Once all writers and readers finish, it prints a final message: "All writers and snapshot readers complete."

		This ensures a clean exit after all concurrent work is done.
	*/
	wg.Wait()
	fmt.Println("All writers and snapshot readers complete.")
}

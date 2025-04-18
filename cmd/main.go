package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func commandListener(commands chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter commands (on/off/exit):")

	for scanner.Scan() {
		cmd := strings.TrimSpace(scanner.Text())
		commands <- cmd
		if cmd == "exit" {
			close(commands)
			return
		}
	}
}

func startOnPrinter(stopChan <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("It is ON")
		case <-stopChan:
			return
		}
	}
}

func commandHandler(commands <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var stopChan chan bool
	var printWG sync.WaitGroup
	isOn := false
	for cmd := range commands {
		switch cmd {
		case "on":
			if !isOn {
				log.Println("Light is ON")
				stopChan = make(chan bool)
				printWG.Add(1)
				go startOnPrinter(stopChan, &printWG)
				isOn = true
			}
		case "off":
			if isOn {
				log.Println("Light is OFF")
				stopChan <- true
				printWG.Wait()
				isOn = false
			}
		case "exit":
			log.Println("Exiting...")
			if isOn {
				stopChan <- true
				printWG.Wait()
			}
			return
		default:
			log.Println("Unknown command:", cmd)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	commands := make(chan string)

	wg.Add(2)
	go commandListener(commands, &wg)
	go commandHandler(commands, &wg)

	wg.Wait()
}


# Introduce Yourself:

I have Btech in IT, Started my Career at Sensibol as Golang Backend Developer. There my Primary role was Developing REST APIs and Microservices, Business logics, Bug Fixes etc. We use Golang Aws MySql Redis etc. After I joined Calsoft and worked with Extreme Network CLient where my primary role was Bug Fixes, Crate APIs, Write Unit and Functional Test cases, We use Same Golang RabbitMQ, SQL gRPC etc.

# Projects:
I worked on PDL: Music distribution and royalty management platform. Basically I was the main work force on the Project.

Also worked on Singshala: Which is similar to Learn to Sing app.

Worked on XCO which is data center automation project, where we manage Multiple Switches and their PORTs and binding them according Profiles Fabric and Tenants.  

# Microservice vs Monolith:
- **Microservices** are better for large projects where scaling and almost zero downtime are required. Bug fixing and maintaining the codebase are easier. A disadvantage of microservices can be inter-service network calls.
- **Monolith** is suitable for smaller projects but can become difficult to scale or maintain as the application grows.

# Authentication vs Authorization
Authentication is about validating the credentials like username password biometrics etc.

Authorization is about granting permission of access or modify specific resource.

# Golang Garbage Collection
Golang uses automatic garbage collection to manage memory. Developers do not need to allocate or deallocate memory manually, which reduces memory-related errors.

# Pointer
A pointer store the memory address of a variable, struct, function, or any data type.

By & we can get the address of the varibale is called reference.

By * we can access the value of the pointer pointing. 
```go
func main() {
    a := 9           // Declare an integer variable 'a'
    p := &a          // 'p' is a pointer that stores the address of 'a' (referencing)

    fmt.Println("Address of a:", &a)    
    // Outputs the memory address of 'a'
    fmt.Println("Value of p:", p)       
    // Outputs the memory address stored in p
    fmt.Println("Value at address p:", *p) 
    // Dereferencing p to get the value (outputs: 9)

    *p = 10          // Change the value of 'a' using the pointer
    fmt.Println("New value of a:", a)   // Outputs: New value of a: 10
}
```
# Goroutine vs Thread:
Goroutines are designed for concurrency, meaning multiple tasks can run using context switching. Threads are designed for parallelism, meaning multiple tasks can run simultaneously on multiple CPU cores.

In context switching, imagine multiple tasks need to run at the same time. The Go runtime starts Task 1, and after some time, it pauses it and starts Task 2. After some time, it pauses Task 2 and can either start Task 3 or resume Task 1, and so on.

Goroutines have a dynamic stack size and are managed by the Go runtime. Threads have a fixed-size stack and are managed by the OS kernel.

The Go scheduler is responsible for deciding when and on which OS thread each goroutine runs. This allows Go to efficiently handle thousands (or more) of goroutines running concurrently.

In Go, when we run multiple goroutines (like mini-tasks), the Go runtime decides when each goroutine should run. By default, Go will let a goroutine keep running until it finishes or does something that allows other tasks to take over (like waiting for I/O or sleeping).

However, sometimes you might want to force Go to give other goroutines a chance to run even if the current goroutine is still doing something. This is where runtime.Gosched() comes in. When you call runtime.Gosched(), it tells Go to stop the current goroutine for a moment and allow other goroutines to run if they need to. It doesn't stop your goroutine permanently, but just gives other tasks a chance to run before your goroutine continues.

runtime.GOMAXPROCS(n): This function sets the maximum number of CPUs that Go will use for running goroutines. By default, Go will use all available CPUs. This function allows you to limit it to n CPUs. runtime.NumCPU(): This function returns the number of CPUs available on the current machine.

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// FibRecursive calculates Fibonacci numbers recursively
func FibRecursive(x int) int {
	if x <= 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return FibRecursive(x-1) + FibRecursive(x-2)
}

// Fibonacci calculation for each core, passing the index as the Fibonacci number to calculate
func calculateFibonacci(id int, cpun int, fibNum int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Start Fibonacci calculation (it'll be a high CPU operation)
	fmt.Printf("CPU %d :- Goroutine %d calculating Fibonacci(%d)\n", cpun, id, fibNum)
	result := FibRecursive(fibNum)
	fmt.Printf("CPU %d :- Goroutine %d result: Fibonacci(%d) = %d\n", cpun, id, fibNum, result)
}

func main() {
	// Get the number of CPUs on the machine
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs available: %d\n", numCPU)

	// Set the number of CPUs that Go can use to all available CPUs
	runtime.GOMAXPROCS(numCPU)

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Distribute Fibonacci numbers across CPU cores starting from Fibonacci(12)
	for i := 0; i < numCPU; i++ {
		fibNum := 32 + i                           // Start from Fibonacci(12) and increase with each core
		wg.Add(1)                                  // Increment the WaitGroup counter
		go calculateFibonacci(i+1, i, fibNum, &wg) // Launch a goroutine with a unique Fibonacci number
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Simulate a little pause to allow you to inspect CPU usage
	time.Sleep(5 * time.Second)
	fmt.Println("DONE---------")
}
```

# Closure in golang:
A closure is a special type of anonymous function that can use variables, that declared outside of the function. Closures treat functions as values, allowing us to assign functions to variables, pass functions as arguments, and return functions from other functions.

```go
func main() {
	a := func(i int) int { return i * i }
	v := func(f func(i int) int) func() int {
		fmt.Println("Inside : Address of f", &f)
		c := f(3)
		return func() int {
			fmt.Println("Inside : Address of c", &c, " Value ", c)
			c++
			return c
		}
	}
	f := v(a) // v() returns a function address
	fmt.Println("Inside f is ", f)
	fmt.Println("f Function Called and get", f()) // call the function which v() returns
}
```

```bash
Inside : Address of f 0xc000066040
Inside f is  0x4b1f40
Inside : Address of c 0xc000010120  Value  9
f Function Called and get 10
```
# Panic Defer Recover combo:
panic is use to cause a Runtime Error and Stop the execution. When a function return or panicking then Defer blocks are called according to Last in First out manner, the last defer will execute first. Recover is use to regain the execution from a panicking situation and handle it properly then stop execution. Recover is useful for close any connection like db and websockets etc.

```go
func div(num int) int {
	if num == 0 {
		panic("Not divisible by 0")
	} else {
		return 27 / num
	}
}
func rec() {
	r := recover()
	if r != nil {
		fmt.Println("I am recovering from Panic")
		fmt.Println("I am Fine Now")
	}
}
func main() {
	defer rec()
	fmt.Println(div(0))
	fmt.Println("Main Regained") // Will not executed if divisble by 0
}
```
# Interface 
Interfaces allow us to define contracts, which are abstract methods, have no body/implementations of the methods. A Struct which wants to implements the Interface need to write the body of every abstract methods the interface holds. We can compose interfaces together. An empty interface can hold any type of values. name.(type) give us the Type the interface will hold at runtime. or we can use reflect.TypeOf(name)
```go
func ty(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	default:
		fmt.Println("No idea")
	}
	fmt.Println(reflect.TypeOf(i))
}
func main() {
	ty(67.89)
}
```
```go
func main() {
    var i interface{} = 42 // you can change this to any type

    switch reflect.TypeOf(i) {
    case reflect.TypeOf(0):
        fmt.Println("i is an int")
    case reflect.TypeOf(""):
        fmt.Println("i is a string")
    case reflect.TypeOf(true):
        fmt.Println("i is a bool")
    default:
        fmt.Println("Unknown type:", reflect.TypeOf(i))
    }
}
```
# TypeCasting Vs TypeConversion:

In Typecasting Compiler will automatically convert between types where smaller type value can convert to Larger type value automatically without Code.
Go does not support it.

In TypeConversion in code we can convert between types if those two types are compatible. Go Support this.

# Array and Slices
Array can not grow or shrink at runtime, Slice can. 
Slice are basically reference of an Array and at RUNTIME it will manipulated such a way that it behave like it can grow or shrink.

```go
values:=make([]int, 4)
values=append(values, 2, 56, 12) // OUTPUT: [0 0 0 0 6 9 2 56 12]
```

Empty slice allocate Heap memory, But Nil slice did not allocate Heap memory.

```go
var sl []int // Nil slice

var sl:=[]int{} // Empty slice
```
Deep copy a slice using either the built-in copy or append function to get a duplicated slice. src and dst slices have different backing arrays. You can see the memory address of their backing array are different.

```go
src := []int{1, 2, 3, 4, 5}
dst := append([]int{}, src...)
dst := make([]int, len(src))
```
Shallow copy whose slice header points to the same backing array as the source slice. By assignment, the memory address of the backing array is the same for both of src and dst slices.

```go
src := []int{1, 2, 3, 4, 5}
dst := src
```
# Concurency Primitives:
Concurency Primitives are tools that are provided by any programming languages to handle execution behaviors of Concurent tasks.

In golang we have Mutex, Semaphore, Channels as concurency primitives.

Mutex is used to protect shared resources from being accessed by multiple threads simultaneously.

Semaphore is used to protect shared pool of resources from being accessed by multiple threads simultaneously. Semaphore is a Counter which start from Number of Reosurces. When one thread using the reosurces Semaphore decremented by 1. If semaphore value is 0 then thread will wait untils its value greater than 0. When one thread done with the resources then Semaphore incremented by 1.

Channel is used to communicate via sending and receiving data and provide synchronisation between multiple gorountines. If channel have a value then execution blocked until reader reads from the channel. Channel can be buffered, allowing goroutines to send multiple values without blocking until the buffer is full.

Waitgroup is used when we want the function should wait until goroutines complete its task. Waitgroup has Add() function which increments the wait-counter for each goroutine. Wait() is used for wait until wait-counter became zero. Done() decrement wait-counter and it called when goroutine complete its task.

# Map Synchronisation:

In golang if multiple goroutines try to acess map at same time, then the operations leads to Panic for RACE or DEADLOCK (fatal error: concurrent map read and map write). So we need proper codes for handeling Map. We use MUTEX for LOCK and UNLOCK the Map operations like Read and Write.
```go
func producer(m *map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	vm := *m
	for i := 0; i < 5; i++ {
		mu.Lock()
		vm[i] = fmt.Sprint("$", i)
		mu.Unlock()
	}
	m = &vm
	wg.Done()
}
func consumer(m *map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	vm := *m
	for i := 0; i < 5; i++ {
		mu.RLock()
		fmt.Println(vm[i])
		mu.RUnlock()
	}
	wg.Done()
}
func main() {
	m := make(map[int]string)
	m[0] = "1234"
	m[3] = "2345"
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go producer(&m, &wg, &mu)
		go consumer(&m, &wg, &mu)
	}
	wg.Wait()
}
```
```go
func producer(ch chan<-int){ //ch <- value // Sends value into channel ch
	for i:=0;i<5;i++{
		ch<-i
	}
	close(ch)
}
func consumer(ch <-chan int){ //value := <-ch // Receives from channel ch and assigns it to value
	for n:=range ch{
		fmt.Print("RECV: ",num)
	}
}
for main(){
	ch:=make(chan int)
	go producer(ch)
	consumer(ch)
}
```
```go
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func primeHelper(a []int, ch chan<- map[int]bool, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	defer wg.Done()
	m := make(map[int]bool)
	for i := range a {
		m[a[i]] = isPrime(a[i])
	}
	ch <- m
}
func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	n := 12
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	length := len(arr)
	goroutines := 4
	part := length / goroutines
	ch := make(chan map[int]bool, goroutines)
	ma := make(map[int]bool)
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		s := i * part
		e := s + part
		if e > length {
			e = length
		}
		go primeHelper(arr[s:e], ch, &wg)
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		for k, v := range i {
			ma[k] = v
		}
	}
	fmt.Println(ma)
	fmt.Println("Time Taken: ", time.Since(startTime))
}
```
Two goroutines generate even and odd numbers up to 30 and send them through their respective channels (evenCh and oddCh). The third goroutine prints numbers alternately from these channels, ensuring the sequence. The sync.WaitGroup ensures that the program waits for all goroutines to finish before exiting.
```go
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

	// Goroutine to send even numbers up to 30
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 30; i++ {
			if i%2 == 0 {
				evenCh <- i
			}
		}
		close(evenCh)
	}()

	// Goroutine to send odd numbers up to 30
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 30; i++ {
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
			select {
			case even, ok := <-evenCh:
				if ok {
					fmt.Println("Even:", even)
				}
			case odd, ok := <-oddCh:
				if ok {
					fmt.Println("Odd:", odd)
				}
			}
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
```

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 1: Prints a message and simulates some work
func task1(done chan bool) {
	fmt.Println("Task 1: Starting")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Task 1: Completed")

	// Signal Task 2 that Task 1 is done
	done <- true
}

// Task 2: Prints a message and simulates some work
func task2(done chan bool) {
	// Wait for the signal from Task 1 that it is complete
	<-done

	fmt.Println("Task 2: Starting")
	time.Sleep(3 * time.Second) // Simulate work
	fmt.Println("Task 2: Completed")
}

func main() {
	// Create a channel to signal when Task 1 is done
	done := make(chan bool)

	// Start Task 1 in a goroutine
	go task1(done)

	// Start Task 2 in a goroutine, but it will wait for Task 1 to complete
	go task2(done)

	// Wait for Task 1 and Task 2 to complete before finishing the program
	time.Sleep(6 * time.Second)
	fmt.Println("All tasks completed!")
}
```

# Limited Concurrency
We create a channel sem := make(chan struct{}, concurrencyLimit) with a buffer size of concurrencyLimit (in this case, 5). This means only 5 goroutines can acquire a slot at any given time. Inside each goroutine, sem <- struct{}{} is used to acquire a slot. When a goroutine finishes its task, it releases the slot using <-sem. If the semaphore is full (i.e., 5 goroutines are already running), any additional goroutines will block (wait) until a slot is available.
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when the task is done

	sem <- struct{}{} // Acquire a slot (semaphore)
	fmt.Printf("Goroutine %d is starting\n", id)
	time.Sleep(1 * time.Second) // Simulate work by sleeping for 1 second
	fmt.Printf("Goroutine %d is finished\n", id)
	<-sem // Release the slot (semaphore)
}

func main() {
	var wg sync.WaitGroup
	concurrencyLimit := 5
	sem := make(chan struct{}, concurrencyLimit) // Semaphore with a limit of 5 concurrent goroutines

	// Launch 40 goroutines
	for i := 1; i <= 40; i++ {
		wg.Add(1)            // Add a goroutine to the wait group
		go task(i, sem, &wg) // Start the task as a goroutine
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All goroutines have finished.")
}
```
# Concurent File Read and Write
```go
package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func FileWriter(ch <-chan string, filename string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	defer wg.Done()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for line := range ch {
		mu.Lock()
		_, err = f.WriteString(line + "\n")
		mu.Unlock()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
func FileReader(filename string, interval time.Duration, wg *sync.WaitGroup, mu *sync.RWMutex) {
	defer wg.Done()
	for {
		time.Sleep(interval)
		mu.RLock()
		d, err := os.ReadFile(filename)
		mu.RUnlock()
		if err != nil {
			log.Println("error: ", err)
			continue
		}
		log.Println("Reading: \n", string(d))
	}
}

func main() {
	filename := "log.txt"
	ch := make(chan string, 100)

	var wgR sync.WaitGroup
	var wgW sync.WaitGroup
	var mu sync.RWMutex

	// Start writer goroutine
	wgW.Add(1)
	go FileWriter(ch, filename, &wgW, &mu)

	// Start reader goroutine (reads every seconds)
	wgR.Add(1)
	go FileReader(filename, time.Second, &wgR, &mu)

	// Simulate writing from multiple goroutines
	for i := range 10 {
		msg := fmt.Sprintf("Message from goroutine %d", i)
		ch <- msg
		time.Sleep(time.Second)
	}

	close(ch)
	wgW.Wait()

	// Optional: end reader after some time
	time.Sleep(5 * time.Second)
	fmt.Println("Done.")
}
```
# Listener GoRoutine
```go
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

```
### SOLID Principles:
SOLID priciples are guidelines for designing Code base that are easy to understand maintain and extend over time.

Single Responsibility:- A Struct/Class should have only a single reason to change. Fields of Author shoud not placed inside Book Struct.
```go
type Book struct{
  ISIN string
  Name String
  AuthorID string
}
type Author struct{
  ID string
  Name String
}
```
Assume One Author decided later, he does not want to Disclose its Real Name to Spread. So we can Serve Frontend by Alias instead of Real Name. Without Changing Book Class/Struct, we can add Alias in Author Struct. By that, Existing Authors present in DB will not be affected as Frontend will Change Name only when it Founds that - Alias field is not empty.
```go
type Book struct{
  ISIN string
  Name String
  AuthorID string
}
type Author struct{
  ID string
  Name String
  Alias String
}

```
Open Close:- Struct and Functions should be open for Extension but closed for modifications. New functionality to be added without changing existing Code.
```go
type Shape interface{
	Area() float64
}
type Rectangle struct{
	W float64
	H float64
}
type Circle struct{
	R float64
}
```
Now we want to Calculate Area of Rectangle and Circle, so Rectangle and Circle both can Implements Shape Interface by Write Body of the Area() Function.
```go
func (r Rectangle) Area()float64{
	return r.W * r.H
}
func (c Circle)Area()float64{
	return 3.14 * c.R * c.R
}
```
Now we can create a Function PrintArea() which take Shape as Arguments and Calculate Area of that Shape. So here Shape can be Rectangle, Circle. In Future we can add Triangle Struct which implements Shape interface by writing Body of Area. Now Traingle can be passed to PrintArea() with out modifing the PrintArea() Function.
```go
func PrintArea(shape Shape) {
	fmt.Printf("Area of the shape: %f\n", shape.Area())
}

// In Future
type Triangle struct{
	B float64
	H float54
}
func (t Triangle)Area()float64{
	return 1/2 * t.B * t.H
}

func main(){
	rect:= Rectangle{W:5,H:3}
	cir:=Circle{R:3}
	PrintArea(rect)
	PrintArea(cir)
	// In Future
	tri:=Triangle{B:4,H:8}
	PrintArea(tri)
}
```
Liskov Substitution:- Super class Object can be replaced by Child Class object without affecting the correctness of the program.
```go
type Bird interface{
	Fly() string
}
type Sparrow struct{
	Name string
}
type Penguin struct{
	Name string
}
```
Sparrow and Pengin both are Bird, But Sparrow can Fly, Penguin Not. ShowFly() function take argument of Bird type and call Fly() function. Now as Penguin and Sparrow both are types of Bird, they should be passed as Bird within ShowFly() function.
```go
func (s Sparrow) Fly() string{
	return "Sparrow is Flying"
}
func (p Penguin) Fly() string{
	return "Penguin Can Not Fly"
}

func ShowFly(b Bird){
	fmt.Println(b.Fly())
}
func main() {
	sparrow := Sparrow{Name: "Sparrow"}
	penguin := Penguin{Name: "Penguin"}
  // SuperClass is Bird,  Sparrow, Penguin are the SubClass
	ShowFly(sparrow)
	ShowFly(penguin)
}
```
Interface Segregation:- A class should not be forced to implements interfaces which are not required for the class. Do not couple multiple interfaces together if not necessary then. 
```go
// The Printer interface defines a contract for printers with a Print method.
type Printer interface {
	Print()
}
// The Scanner interface defines a contract for scanners with a Scan method.
type Scanner interface {
	Scan()
}
// The NewTypeOfDevice interface combines Printer and Scanner interfaces for
// New type of devices which can Print and Scan with it new invented Hardware.
type NewTypeOfDevice interface {
	Printer
	Scanner
}
```

Dependecy Inversion:- Class should depends on the Interfaces not the implementations of methods.

```go
// The MessageSender interface defines a contract for 
//sending messages with a SendMessage method.
type MessageSender interface {
	SendMessage(msg string) error
}
// EmailSender and SMSClient structs implement 
//the MessageSender interface with their respective SendMessage methods.
type EmailSender struct{}

func (es EmailSender) SendMessage(msg string) error {
	fmt.Println("Sending email:", msg)
	return nil
}
type SMSClient struct{}

func (sc SMSClient) SendMessage(msg string) error {
	fmt.Println("Sending SMS:", msg)
	return nil
}
type NotificationService struct {
	Sender MessageSender
}
```
The NotificationService struct depends on MessageSender interface, not on concrete implementations (EmailSender or SMSClient). This adheres to Dependency Inversion, because high-level modules (NotificationService) depend on abstractions (MessageSender) rather than details.
```go
func (ns NotificationService) SendNotification(msg string) error {
	return ns.Sender.SendMessage(msg)
}
func main() {
	emailSender := EmailSender{}

	emailNotification := NotificationService{Sender: emailSender}

	emailNotification.SendNotification("Hello, this is an email notification!")
}
```
### Create Post and Get APIs:
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// ---------- Transport -----------

type httpTransport struct {
	Service // Interfaces use for Dependency Inversion
}

func NewHttpTransport(s Service) *httpTransport {
	return &httpTransport{
		Service: s,
	}
}

type AddReq struct {
}
type AddRes struct {
	Success bool `json:"success"`
}
type GetReq struct {
	Id string `json:"id"`
}
type GetUrlReq struct {
	Id string `url:"id"`
}
type GetRes struct {
	Count   int  `json:"Count"`
	Success bool `json:"success"`
}

func (t *httpTransport) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t.AddCounter()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AddRes{Success: true})
}
func (t *httpTransport) Get(w http.ResponseWriter, r *http.Request) {
	var req GetReq
	json.NewDecoder(r.Body).Decode(&req)
	count := t.GetCounter(req.Id)
	json.NewEncoder(w).Encode(GetRes{Count: count, Success: true})
}
func (t *httpTransport) GetUrl(w http.ResponseWriter, r *http.Request) {
	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		json.NewEncoder(w).Encode(GetRes{Count: -1, Success: false})
	}
	id := u["id"][0]
	count := t.GetCounter(id)
	json.NewEncoder(w).Encode(GetRes{Count: count, Success: true})
}

// ---------- Domain -----------
type Counter struct {
	Count int
}

// ---------- Service -----------
type Service interface {
	AddCounter()
	GetCounter(id string) int
}
type service struct {
	Counter Counter
}

func NewService(c Counter) *service {
	return &service{Counter: c}
}
func (s *service) AddCounter() {
	s.Counter.Count++
}
func (s *service) GetCounter(id string) int {
	fmt.Println(id)
	return s.Counter.Count
}

// ---------- Main -----------
func main() {
	// Service
	s := NewService(Counter{})
	// Trnasport
	t := NewHttpTransport(s)
	// Routing
	r := mux.NewRouter()
	r.HandleFunc("/api/add", t.Add).Methods("POST")
	r.HandleFunc("/api/get", t.Get).Methods("GET")
	r.HandleFunc("/api/geturl", t.GetUrl).Methods("GET")
	// Server
	http.ListenAndServe(":3000", r)
}
```

# Pattern Question
```go

/*
1,4,7
2,5,8
3,6,9
*/

package main

import (
	"fmt"
	"sync"
)

func sender1(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i += 3 {
		ch <- i
	}
	close(ch)
}
func sender2(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 20; i += 3 {
		ch <- i
	}
	close(ch)
}
func sender3(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 3; i <= 20; i += 3 {
		ch <- i
	}
	close(ch)
}
func main() {
	wg := sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	wg.Add(3)
	go sender1(ch1, &wg)
	go sender2(ch2, &wg)
	go sender3(ch3, &wg)
	res1 := []int{}
	res2 := []int{}
	res3 := []int{}

	for v := range ch1 {
		res1 = append(res1, v)
	}
	for v := range ch2 {
		res2 = append(res2, v)
	}
	for v := range ch3 {
		res3 = append(res3, v)
	}

	wg.Wait()
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}
```

```go
// O(log n), because with each comparison, the search space is halved, leading to a logarithmic number of comparisons.
// Binary Search function for a sorted slice
func binarySearch(slice []int, target int) int {
    low := 0
    high := len(slice) - 1

    for low <= high {
        mid := (low + high) / 2
        if slice[mid] == target {
            return mid // Target found at index 'mid'
        } else if slice[mid] < target {
            low = mid + 1 // Search in the right half
        } else {
            high = mid - 1 // Search in the left half
        }
    }
    return -1 // Target not found
}
```

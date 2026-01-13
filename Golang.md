
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

In Go, garbage collection (GC) runs automatically to clean up memory your program no longer uses. However, you can adjust how often it runs to improve performance.

GOGC controls how much the heap (memory) can grow before garbage collection starts.

GOGC=200 → GC runs less often (more memory used, but less CPU work).

GOGC=50 → GC runs more often (less memory used, but more CPU work).

GOGC=off → Disables GC (not recommended unless you know what you're doing).

You can set this in the terminal before running your program
GOGC=100 Default
GOGC=200 ./myprogram

If you want to force garbage collection at a specific time, use:
```go
import "runtime"

func main() {
    runtime.GC()
}
```
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
# Go-Routine vs Thread:
Go-Routines are designed for concurrency, meaning multiple tasks can run using context switching. Threads are designed for parallelism, meaning multiple tasks can run simultaneously on multiple CPU cores.

In context switching, imagine multiple tasks need to run at the same time. The Go-Runtime starts Task 1, and after some time, it pauses it and starts Task 2. After some time, it pauses Task 2 and can either start Task 3 or resume Task 1, and so on.

Go-Routines have a dynamic stack size and are managed by the Go-Runtime. Threads have a fixed-size stack and are managed by the OS kernel. This dynamic growth of Stack ensures that memory is used efficiently. As a result, thousands or even millions of Go-Routines can be created and run concurrently with minimal memory overhead. This fixed size stack means that even if the thread doesn’t need all that memory, the space is reserved. As a result, threads consume more memory, which limits how many threads can run concurrently on a system.

The Go scheduler is responsible for deciding when and on which OS-Thread each Go-Routine runs.

When we start a Go-Routine using the `go` keyword, the Go-Runtime adds it to a Run-Queue. This Run-Queue is managed by a special component called a Logical-Processor. The Logical-Processor holds the queue of Go-Routines, manages their state, and helps with scheduling and garbage collection.

A Go-Routine waits in the Run-Queue until an OS-Thread picks it up for execution. Once picked, the Go-Routine starts running.

While running, a Go-Routine can: Be preempted: The runtime can pause it if it runs too long to give others a chance. or Pause itself: If it waits on I/O, channels, mutexes, or syscalls. or Yield: It can voluntarily give up the CPU to let other Go-Routines run.

If a Go-Routine gets blocked (e.g., waiting for data), it enters a waiting state. The Go scheduler then detaches the current OS-Thread so the Logical-Processor can keep working using another thread.
 
A Logical-Processor is tied to an OS-Thread, and a Logical-Processor can only run one goroutine at a time on a particular OS-Thread. So, if the OS-Thread is blocked, we can’t use the same thread to run another goroutine unless we assign a different OS-Thread to the Logical-Processor

Once the blocking condition is resolved (like data is received), the Go-Routine is rescheduled and placed back into a Run-Queue to be picked up again.

When a context switch happens (i.e., the runtime switches from one Go-Routine to another), Go saves the state of the current Go-Routine and loads the state of the next one. The OS-Thread then continues executing from where that new Go-Routine left off.

Because Go-Routines have very small stacks and are scheduled in user space, context switching between them is very fast and much cheaper than switching between OS-Threads.

However, sometimes we might want to force Go to give other Go-Routines a chance to run even if the current Go-Routine is still doing something. This is where runtime.Gosched() comes in. When we call runtime.Gosched(), it tells Go to stop the current Go-Routine for a moment and allow other Go-Routines to run if they need to. It doesn't stop our Go-Routine permanently, but just gives other tasks a chance to run before our Go-Routine continues.

The number of Logical-Processors is controlled by GOMAXPROCS, and it determines how many Go-Routines can run in parallel (at the same time).

runtime.GOMAXPROCS(n): This function sets the maximum number of CPUs that Go will use for running Go-Routines. By default, Go will use all available CPUs. This function allows us to limit it to n CPUs. runtime.NumCPU(): This function returns the number of CPUs available on the current machine.

```go
func task1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Task 1 - Iteration", i)
		if i == 2 {
			// Yield the CPU to other goroutines
			// This tells the Go runtime to yield and let other goroutines run before continuing with task1(). This allows task2 to run when task1 is paused.
			runtime.Gosched()
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func task2() {
	for i := 0; i < 5; i++ {
		fmt.Println("Task 2 - Iteration", i)
		time.Sleep(time.Millisecond * 500)
	}
}
```

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// FibRecursive calculates Fibonacci numbers recursively
func FibRecursive(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// Fibonacci calculation for each core, passing the index as the Fibonacci number to calculate
func calculateFibonacci(id int, cpun int, fibNum int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Start Fibonacci calculation (it'll be a high CPU operation)
	fmt.Printf("CPU %d :- Go-Routine %d calculating Fibonacci(%d)\n", cpun, id, fibNum)
	result := FibRecursive(fibNum)
	fmt.Printf("CPU %d :- Go-Routine %d result: Fibonacci(%d) = %d\n", cpun, id, fibNum, result)
}

func main() {
	// Get the number of CPUs on the machine
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of CPUs available: %d\n", numCPU)

	// Set the number of CPUs that Go can use to all available CPUs
	runtime.GOMAXPROCS(numCPU)

	// Create a WaitGroup to wait for all Go-Routines to finish
	var wg sync.WaitGroup

	// Distribute Fibonacci numbers across CPU cores starting from Fibonacci(32)
	for i := 0; i < numCPU; i++ {
		fibNum := 32 + i                           // Start from Fibonacci(32) and increase with each core
		wg.Add(1)                                  // Increment the WaitGroup counter
		go calculateFibonacci(i+1, i, fibNum, &wg) // Launch a Go-Routine with a unique Fibonacci number
	}

	// Wait for all Go-Routines to finish
	wg.Wait()

	// Simulate a little pause to allow us to inspect CPU usage
	time.Sleep(5 * time.Second)
	fmt.Println("DONE---------")
}
```

# Closure in golang:
A closure is a special type of anonymous function that can access/use variables, that declared outside of the function. Closures treat functions as values, allowing us to assign functions to variables, pass functions as arguments, and return functions from other functions.

```go
func main() {
	a := func(i int) int { return i * i } // A function that takes an int and returns its square.
	v := func(f func(i int) int) func() int {
		fmt.Println("Inside : Address of f", &f) // This is the address of the function parameter `f`, which is a local copy of the function `a` passed into `v`.
		c := f(3)
		return func() int {
			fmt.Println("Inside : Address of c", &c, " Value ", c) // This is the address of the local variable `c` inside `v`. When the closure is returned, it captures `c` and keeps it in its own memory space.
			c++
			return c
		}
	}
	f := v(a)                                     // v() returns a function address
	fmt.Println("Inside f is ", f)                // This is the address of the returned closure function, func() int. It’s a separate function, dynamically created and returned by v.
	fmt.Println("f Function Called and get", f()) // call the function which v() returns // When f() is called, it returns c+1, showing that the closure is preserving the state of c between invocations (i.e., it's a closure over c).
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

```go
func demo() {
    defer func() {
        recover() // recover from outer panic
        panic("panic in defer") // new panic
    }()
    panic("outer panic") 
}
```

` panic("outer panic") ` This initiates stack unwinding. Go starts looking for a defer function to run. The deferred function runs. `recover()` is called. This catches the `"outer panic"` — it stops the unwinding process. At this point, outer panic is suppressed, and the program is safe again. `panic("panic in defer")` is then called. A new panic is initiated during the defer execution. The program crashes again with the new panic "`panic in defer"`. If we run this code in main(), we will see `panic: panic in defer`


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
    var i interface{} = 42 // we can change this to any type

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
# TypeCasting Vs TypeConversion Vs TypeAssertion:

In Typecasting Compiler will automatically convert between types where smaller type value can convert to Larger type value automatically without Code.
Go does not support it.

In TypeConversion in code we can convert between types if those two types are compatible. Go Support this.

		var x int = 10
		var y float64 = float64(x) // Manual conversion

Type assertion is used to extract the concrete value from an interface type. It tells the Go compiler: I know this interface value is actually of this specific type. 

		var i interface{} = "hello"
		s := i.(string)  // i contains a string, so this is valid
		s2, ok := i.(string)

strconv.Atoi() in Go is not a type conversion in the language's strict sense — it's actually a function that parses a string and returns an integer.

# Array and Slices
Array can not grow or shrink at runtime, Slice can. 
Slice are basically reference of an Array and at RUNTIME it will manipulated such a way that it behave like it can grow or shrink.

In Go, the make function is a built-in function used to initialize and allocate memory for slices, maps, and channels

```go
values:=make([]int, 4)
values=append(values, 2, 56, 12) // OUTPUT: [0 0 0 0 6 9 2 56 12]
```

Empty slice allocate Heap memory, But Nil slice did not allocate Heap memory.

```go
var sl []int // Nil slice

var sl:=[]int{} // Empty slice

if sl == nil {
    fmt.Println("The slice is nil")
}
```
Deep copy a slice using either the built-in copy or append function to get a duplicated slice. src and dst slices have different backing arrays. 

```go
src := []int{1, 2, 3, 4, 5}
dst := append([]int{}, src...) // deep copy
dst := make([]int, len(src))
copy(dst, src) // deep copy
```
A slice in Go is a descriptor (header) with three fields: Pointer to the underlying backing array, Length, Capacity
So when we assign a slice to another, we are just copying the header. Both slices point to the same backing array — that's a shallow copy.

By assignment, the memory address of the backing array is the same for both of src and dst slices.

```go
src := []int{1, 2, 3, 4, 5}
dst := src // shallow copy: same backing array

dst[0] = 99
fmt.Println("src:", src) // [99 2 3 4 5]
fmt.Println("dst:", dst) // [99 2 3 4 5]
```

When we append to a slice and it exceeds its current capacity, Go automatically allocates a new underlying array, usually with a larger capacity.  Allocating a new array with a larger capacity (typically doubling the old one, but this growth strategy can vary).

```go
func main() {
    s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    s2 := s1[:5]          // s2 = [0 1 2 3 4], shares backing array with s1
    s2 = append(s2, 100)  // append to s2
    fmt.Println(s2)       // ???
    fmt.Println(s1)       // ???
}
```
s2 shares the same backing array as s1.
s2 has: Length = 5 , Capacity = 10 (same as s1, since slicing up to index 5 doesn't limit capacity).
`s2 = append(s2, 100)` now s1[5] = 100 too! Because it’s the same array.

s2 is [0 1 2 3 4 100].
s1 is [0 1 2 3 4 100 6 7 8 9].

```go
var a int
var p *int
fmt.Println(a, p)

// Output
0 <nil>
Explanation: a defaults to 0, and p is a nil pointer because it’s not initialized.
```
```go
var arr [3]int
fmt.Println(arr)

// OUTPUT
[0 0 0]
Arrays are value types; all elements default to their zero values.
```
```go
m := map[string]int{"a": 1, "b": 2}
delete(m, "c")
fmt.Println(m)

// OUTPUT
map[a:1 b:2]
Deleting a non-existent key is a no-op in Go.
```
```go
var ch chan int
fmt.Println(ch)

// OUTPUT
<nil>
Uninitialized channels are nil.
```
```go
var m map[string]int
fmt.Println(m)

// OUTPUT
map[]
Nil map prints as an empty map. However, writes to it will panic.
```

```go
func cleanup(i int) {
	fmt.Println("cleanup func:", i)
}

func main() {
	i := 10
	defer cleanup(i) // ① defer function call, i evaluated NOW (10)
	i++              // ② i = 11
	defer func() {   // ③ defer anonymous function (closure)
		fmt.Println("anonymous func:", i)
	}()
	i = 100          // ④ i becomes 100
}
/*
The variable i is initialized to 10.

The first defer statement calls cleanup(i).
At this point, the value of i is evaluated immediately, so cleanup(10) is what gets deferred.

Then, i++ increments i from 10 to 11.

The second defer uses an anonymous function (a closure) that references i.
Unlike regular function calls, closures capture variables by reference, not by value.

After that, i = 100 assigns a new value to i.

When main() finishes, Go executes the deferred functions in last-in, first-out (LIFO) order.

The anonymous function is executed first. Since it references i (now 100), it prints:
anonymous func: 100

Then, the cleanup function is executed with the original deferred value 10, so it prints:
cleanup func: 10
*/
```
```go
func test() (result int) {
	defer func() {
		result += 1
	}()
	return 5
}
// OUTPUT
This means the function has a named return variable called result.
This deferred function is registered to run after the return is called, but before the function actually returns to the caller.
return result = 5, then defer adds 1 before returning.
```
```go
func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
// OUTPUT
Answer is non-deterministic. Likely 3 printed 3 times.
Because all closures share the same i, which becomes 3 by the time they run.
```

```go
func main() {
	s := []int{1, 2, 3}
	for i := range s {
		s = append(s, i)
	}
	fmt.Println(s)
}
// Possible infinite loop depending on append growth. Becomes subtle bug.
```

```go
type A struct{}

func (a A) Hello() {
	fmt.Println("Hello from value receiver")
}

func main() {
	var i interface{} = A{}
	i.(interface {
		Hello()
	}).Hello()
}

// OUTPUT
A method Hello() on value receiver A (not pointer).
A{} is stored in the empty interface i, interface{}
Here, we’re asserting that the dynamic value inside i (which is A{}) implements the interface
i.(interface {Hello()}) :- "Hey i, I think we’re holding something that can say Hello(). Please show it to me."
If i really does hold something that has a Hello() method, Go lets we call it.

```
# *interface{} vs interface{}
interface{} is already a reference type. *interface{} don't act like a pointer to the original value inside the interface. 

If we change the value inside *interface{}, it does not change the original variable the interface was holding.

var i interface{} = x   // i now holds a copy of x
var i interface{} = &x // Store a pointer inside the interface, then type-assert and modify

# Escape Analysis
Escape analysis is the compiler's decision process that determines whether a variable should be allocated on the stack or the heap.

If a variable lives beyond the function's scope, it escapes and gets allocated on the heap. If it doesn’t escape, it stays on the stack, which is faster.

Heap allocations are expensive due to garbage collection.

Avoid returning pointers to local variables unless necessary.

```go
func escapeStack() int {
    x := 42
    return x // x stays on stack
}
func escapeHeap() *int {
    x := 42
    return &x // x escapes to heap
}
```
# Concurency Primitives:
Concurency Primitives are tools that are provided by any programming languages to handle execution behaviors of Concurent tasks.

In golang we have Mutex, Semaphore, Channels as concurency primitives.

Mutex is used to protect shared resources from being accessed by multiple threads simultaneously.

Semaphore is used to protect shared pool of resources from being accessed by multiple threads simultaneously. Semaphore is a Counter which start from Number of Reosurces. When one thread using the reosurces Semaphore decremented by 1. If semaphore value is 0 then thread will wait untils its value greater than 0. When one thread done with the resources then Semaphore incremented by 1.

Channel is used to communicate via sending and receiving data and provide synchronisation between multiple gorountines. If channel have a value then execution blocked until reader reads from the channel. Channel can be buffered, allowing Go-Routines to send multiple values without blocking until the buffer is full.

Waitgroup is used when we want the function should wait until Go-Routines complete its task. Waitgroup has Add() function which increments the wait-counter for each Go-Routine. Wait() is used for wait until wait-counter became zero. Done() decrement wait-counter and it called when Go-Routine complete its task.


# Channel
```go
ch := make(chan int)   // Unbuffered
ch := make(chan int, 2) // Buffered
ch <- 42               // Send
val := <-ch            // Receive
```

## Fan-Out / Fan-In
Fan-Out: Distribute work across multiple goroutines.
```go
type Job struct {
    ID  int
    URL string
}

type Result struct {
    JobID   int
    Content string
    Error   error
}
```

```go
func worker(id int, jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        content, err := fetchURL(job.URL)
        results <- Result{JobID: job.ID, Content: content, Error: err}
    }
}
```
```go
numWorkers := 5
for w := 1; w <= numWorkers; w++ {
    go worker(w, jobs, results)
}
```
```go
go func() {
    for i, url := range urls {
        jobs <- Job{ID: i, URL: url}
    }
    close(jobs) // Important to signal workers there's no more work
}()
```
Fan-In: Combine results from multiple channels into one.
```go
var wg sync.WaitGroup
wg.Add(numWorkers)

for i := 0; i < numWorkers; i++ {
    go func() {
        for res := range results {
            process(res)
        }
        wg.Done()
    }()
}
```
## Pipeline Pattern
The Pipeline Pattern in Go is a powerful and idiomatic way to process data in stages, where each stage runs in its own goroutine, and passes results to the next stage via channels.

```go
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}


func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		total := 0
		for n := range in {
			total += n
		}
		out <- total
		close(out)
	}()
	return out
}

func main() {
	in := generator(1, 2, 3, 4, 5)
	sq := square(in)
	result := sum(sq)

	fmt.Println("Result:", <-result) // Output: 55
}
```

## Worker Pool

```go
type Job struct {
	ID      int
	Payload string
}

type Result struct {
	JobID   int
	Outcome string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		outcome := fmt.Sprintf("Processed: %s", job.Payload)
		fmt.Printf("Worker %d handled job %d\n", id, job.ID)
		results <- Result{JobID: job.ID, Outcome: outcome}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 5

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Payload: fmt.Sprintf("task-%d", j)}
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for res := range results {
		fmt.Printf("Result: Job %d -> %s\n", res.JobID, res.Outcome)
	}
}
```
## Merging multiple channels
```go
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) { 
		// This defines an anonymous function called output that takes a single read-only channel of integers as input (named c). It is assigned to the variable output.
        for v := range c {
			// This starts a loop that reads values from channel c. The range statement will keep receiving values until the channel c is closed.
            out <- v
			// Each value received from c is immediately sent to the out channel.
        }
        wg.Done()
    }

    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    go func() {
		// This starts an anonymous goroutine that will run concurrently with the rest of the code. The reason we wrap it in a go statement is to ensure that close(out) happens only after all worker goroutines have completed their tasks.

		// After the wg.Wait() completes, which means all the workers are done, we can safely close the out channel. This is important because: Closing the output channel signals to the receiver that no more data will be sent to it. It prevents a deadlock (if we closed out before all workers finished, the workers might still try to send data to the closed channel, causing a panic).

        wg.Wait()
        close(out)
    }()
    return out
}
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        for i := 0; i < 3; i++ { ch1 <- i }
        close(ch1)
    }()

    go func() {
        for i := 10; i < 13; i++ { ch2 <- i }
        close(ch2)
    }()

    for val := range merge(ch1, ch2) {
        fmt.Println(val)
    }
}
```
# Map Synchronisation:
A race condition occurs when two or more goroutines access shared resources (like variables or memory) concurrently, and at least one of the accesses is a write operation. 

A deadlock occurs when two or more goroutines are blocked forever because they are waiting for each other to release resources or perform some action. In a deadlock situation, no goroutine can proceed, and the program essentially halts. This happens when there is circular waiting for resources.

In golang if multiple Go-Routines try to acess map at same time, then the operations leads to Panic for RACE or DEADLOCK (fatal error: concurrent map read and map write). So we need proper codes for handeling Map. We use MUTEX for LOCK and UNLOCK the Map operations like Read and Write.
```go
func producer(m map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	
	for i := 0; i < 5; i++ {
		mu.Lock()
		m[i] = fmt.Sprint("$", i)
		mu.Unlock()
	}
	wg.Done()
}
func consumer(m map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	for i := 0; i < 5; i++ {
		mu.RLock()
		fmt.Println(m[i])
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
	Go-Routines := 4
	part := length / Go-Routines
	ch := make(chan map[int]bool, Go-Routines)
	ma := make(map[int]bool)
	for i := 0; i < Go-Routines; i++ {
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
Two Go-Routines generate even and odd numbers up to 30 and send them through their respective channels (evenCh and oddCh). The third Go-Routine prints numbers alternately from these channels, ensuring the sequence. The sync.WaitGroup ensures that the program waits for all Go-Routines to finish before exiting.
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

	// Go-Routine to send even numbers up to 30
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

	// Go-Routine to send odd numbers up to 30
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

	// Wait for all Go-Routines to finish
	wg.Wait()
}
```
	Even: 0
	Even: 2
	Even: 4
	Even: 6
	Even: 8
	Even: 10
	Odd: 1
	Odd: 3
	Odd: 5
	Odd: 7
	Odd: 9

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


```
	Odd: 1
	Even: 2
	Odd: 3
	Even: 4
	Odd: 5
	Even: 6
	Odd: 7
	Even: 8
	Odd: 9
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

	// Start Task 1 in a Go-Routine
	go task1(done)

	// Start Task 2 in a Go-Routine, but it will wait for Task 1 to complete
	go task2(done)

	// Wait for Task 1 and Task 2 to complete before finishing the program
	time.Sleep(6 * time.Second)
	fmt.Println("All tasks completed!")
}
```

# Limited Concurrency
We create a channel sem := make(chan struct{}, concurrencyLimit) with a buffer size of concurrencyLimit (in this case, 5). This means only 5 Go-Routines can acquire a slot at any given time. Inside each Go-Routine, sem <- struct{}{} is used to acquire a slot. When a Go-Routine finishes its task, it releases the slot using <-sem. If the semaphore is full (i.e., 5 Go-Routines are already running), any additional Go-Routines will block (wait) until a slot is available.
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
	fmt.Printf("Go-Routine %d is starting\n", id)
	time.Sleep(1 * time.Second) // Simulate work by sleeping for 1 second
	fmt.Printf("Go-Routine %d is finished\n", id)
	<-sem // Release the slot (semaphore)
}

func main() {
	var wg sync.WaitGroup
	concurrencyLimit := 5
	sem := make(chan struct{}, concurrencyLimit) // Semaphore with a limit of 5 concurrent Go-Routines

	// Launch 40 Go-Routines
	for i := 1; i <= 40; i++ {
		wg.Add(1)            // Add a Go-Routine to the wait group
		go task(i, sem, &wg) // Start the task as a Go-Routine
	}

	wg.Wait() // Wait for all Go-Routines to complete
	fmt.Println("All Go-Routines have finished.")
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

	// Start writer Go-Routine
	wgW.Add(1)
	go FileWriter(ch, filename, &wgW, &mu)

	// Start reader Go-Routine (reads every seconds)
	wgR.Add(1)
	go FileReader(filename, time.Second, &wgR, &mu)

	// Simulate writing from multiple Go-Routines
	for i := range 10 {
		msg := fmt.Sprintf("Message from Go-Routine %d", i)
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
# Listener Go-Routine
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
# Select

```go
func main() {
	dataChan := make(chan string)
	errChan := make(chan error)

	// Context with 5-minute timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Simulate a data-producing goroutine
	go func() {
		time.Sleep(2 * time.Second) // simulate work
		dataChan <- "some result"
	}()
	go func() {
		time.Sleep(3 * time.Second) // simulate work
		errChan <- fmt.Errorf("something went wrong")
	}()

	for {
		select {
		case data := <-dataChan:
			fmt.Println("Received data:", data)

		case err := <-errChan:
			fmt.Println("Received error:", err)
			return // or continue depending on need

		case <-ctx.Done():
			fmt.Println("Context expired:", ctx.Err())
			return
		}
	}
}

/*
rupx@dev:~/projects/interview-preparation/cmd$ go run main.go 
Received data: some result
Received error: something went wrong
rupx@dev:~/projects/interview-preparation/cmd$ 
*/
```
In a select statement, a case with a nil channel is disabled. It will never be chosen. Default case will execute.
Accidentally reading from a nil channel without a select causes a deadlock:
```go
var ch chan int // nil channel
select {
case <-ch:
    fmt.Println("Won't happen")
default:
    fmt.Println("Default case")
}
```
```go
<-ch // blocks forever
```
# SOLID Principles:
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

# Design Patterns

- Singleton Pattern:  The Singleton pattern ensures a class has only one instance and provides a global point of access to it. The singleton instance is created only when it is first requested, rather than at the time of class loading. It should be ned to be thread-safe, ensuring that multiple threads can safely access the singleton instance concurrently without creating multiple instances. Singletone n pattern can use at Storing and accessing global settings or configurations throughout the application, database connection pool or a logging service where multiple parts of an application need access to a single resource.

```go
package main

import (
    "fmt"
    "sync"
)

// ConfigManager is a singleton struct for managing configuration settings.
type ConfigManager struct {
    config map[string]string //config is a map storing key-value pairs of configuration settings.
    // other fields as needed
}

var (
    instance *ConfigManager
    once     sync.Once
)

// GetConfigManager returns the singleton instance of ConfigManager.
// This function provides global access to the singleton instance of ConfigManager
func GetConfigManager() *ConfigManager {
    once.Do(func() { // sync.Once ensures that, the initialization code, inside once.Do() is executed exactly once, preventing multiple initializations even with concurrent calls.
        instance = &ConfigManager{
            config: make(map[string]string),
        }
        // Initialize configuration settings here
        instance.initConfig()
    })
    return instance
}

// initConfig simulates loading initial configuration settings.
// This method initializes the initial configuration settings. In a real application, this might involve reading settings from a configuration file, a database, or environment variables.
func (cm *ConfigManager) initConfig() {
    // Load configuration settings from file, database, etc.
    cm.config["server_address"] = "localhost"
    cm.config["port"] = "8080"
    // Add more configuration settings as needed
}

// GetConfig retrieves a specific configuration setting.
func (cm *ConfigManager) GetConfig(key string) string {
    return cm.config[key]
}

func main() {
    // Get the singleton instance of ConfigManager
    configManager := GetConfigManager()

    // Access configuration settings
    fmt.Println("Server Address:", configManager.GetConfig("server_address"))
    fmt.Println("Port:", configManager.GetConfig("port"))
}

```

- Builder Pattern:  The Builder pattern in Go (Golang) is used to construct complex objects step by step. 

```go
package main

import "fmt"

// Product represents the complex object we want to build.
type Product struct {
    Part1 string
    Part2 int
    Part3 bool
}

// Builder interface defines the steps to build the Product.
type Builder interface {
    SetPart1(part1 string)
    SetPart2(part2 int)
    SetPart3(part3 bool)
    Build() Product
}

// ConcreteBuilder implements the Builder interface.
type ConcreteBuilder struct {
    part1 string
    part2 int
    part3 bool
}

func (b *ConcreteBuilder) SetPart1(part1 string) {
    b.part1 = part1
}

func (b *ConcreteBuilder) SetPart2(part2 int) {
    b.part2 = part2
}

func (b *ConcreteBuilder) SetPart3(part3 bool) {
    b.part3 = part3
}

func (b *ConcreteBuilder) Build() Product {
    // Normally, here we could implement additional logic or validation
    // before returning the final product.
    return Product{
        Part1: b.part1,
        Part2: b.part2,
        Part3: b.part3,
    }
}

// Director controls the construction process using a Builder.
type Director struct {
    builder Builder
}

func NewDirector(builder Builder) *Director {
    return &Director{builder: builder}
}

func (d *Director) Construct(part1 string, part2 int, part3 bool) Product {
    d.builder.SetPart1(part1)
    d.builder.SetPart2(part2)
    d.builder.SetPart3(part3)
    return d.builder.Build()
}

func main() {
    // Create a concrete builder instance
    builder := &ConcreteBuilder{}

    // Create a director with the concrete builder
    director := NewDirector(builder)

    // Construct the product
    product := director.Construct("Example", 42, true)

    // Print the constructed product
    fmt.Printf("Constructed Product: %+v\n", product)
}
```


- Factory Pattern: Factory pattern is typically used to encapsulate the instantiation of objects, allowing the client code to create objects without knowing the exact type being created.

Here's an example of implementing the Factory pattern with repository selection based on database type:

```go
package main

import (
	"fmt"
	"errors"
)

// Repository interface defines the methods that all concrete repository implementations must implement.
type Repository interface {
	GetByID(id int) (interface{}, error)
	Save(data interface{}) error
	// Add other methods as needed
}

// MySQLRepository represents a concrete implementation of Repository interface for MySQL.
type MySQLRepository struct {
	// MySQL connection details or any necessary configuration
}

func (r *MySQLRepository) GetByID(id int) (interface{}, error) {
	// Implement MySQL specific logic to fetch data by ID
	return nil, errors.New("not implemented")
}

func (r *MySQLRepository) Save(data interface{}) error {
	// Implement MySQL specific logic to save data
	return errors.New("not implemented")
}

// PostgreSQLRepository represents a concrete implementation of Repository interface for PostgreSQL.
type PostgreSQLRepository struct {
	// PostgreSQL connection details or any necessary configuration
}

func (r *PostgreSQLRepository) GetByID(id int) (interface{}, error) {
	// Implement PostgreSQL specific logic to fetch data by ID
	return nil, errors.New("not implemented")
}

func (r *PostgreSQLRepository) Save(data interface{}) error {
	// Implement PostgreSQL specific logic to save data
	return errors.New("not implemented")
}

// RepositoryFactory is the factory that creates different types of repositories based on the database type.
type RepositoryFactory struct{}

func (f *RepositoryFactory) CreateRepository(databaseType string) (Repository, error) {
	switch databaseType {
	case "mysql":
		return &MySQLRepository{}, nil
	case "postgresql":
		return &PostgreSQLRepository{}, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func main() {
	factory := RepositoryFactory{}

	// Create a MySQL repository
	mysqlRepo, err := factory.CreateRepository("mysql")
	if err != nil {
		fmt.Println("Error creating MySQL repository:", err)
	} else {
		fmt.Println("Created MySQL repository successfully")
		// Use mysqlRepo as needed
	}

	// Create a PostgreSQL repository
	postgresqlRepo, err := factory.CreateRepository("postgresql")
	if err != nil {
		fmt.Println("Error creating PostgreSQL repository:", err)
	} else {
		fmt.Println("Created PostgreSQL repository successfully")
		// Use postgresqlRepo as needed
	}

	// Try to create a repository with an unsupported database type
	_, err = factory.CreateRepository("mongodb")
	if err != nil {
		fmt.Println("Error creating MongoDB repository:", err)
	}
}
```


- Observer Pattern: The Observer pattern defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
```go
package main

import (
	"fmt"
	"time"
)

// Observer defines the interface that all observers (subscribers) must implement.
type Observer interface {
	Update(article string)
}

// Subject defines the interface for the subject (publisher) that observers will observe.
type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	Notify(article string)
}

// NewsService represents a concrete implementation of the Subject interface.
type NewsService struct {
	observers []Observer
}

func (n *NewsService) Register(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NewsService) Deregister(observer Observer) {
	for i, obs := range n.observers {
		if obs == observer {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NewsService) Notify(article string) {
	for _, observer := range n.observers {
		observer.Update(article)
	}
}

// NewsSubscriber represents a concrete implementation of the Observer interface.
type NewsSubscriber struct {
	Name string
}

func (s *NewsSubscriber) Update(article string) {
	fmt.Printf("[%s] Received article update: %s\n", s.Name, article)
}

func main() {
	// Create a news service
	newsService := &NewsService{}

	// Create news subscribers (observers)
	subscriber1 := &NewsSubscriber{Name: "John"}
	subscriber2 := &NewsSubscriber{Name: "Alice"}
	subscriber3 := &NewsSubscriber{Name: "Bob"}

	// Register subscribers with the news service
	newsService.Register(subscriber1)
	newsService.Register(subscriber2)
	newsService.Register(subscriber3)

	// Simulate publishing new articles
	go func() {
		for i := 1; i <= 5; i++ {
			article := fmt.Sprintf("Article %d", i)
			newsService.Notify(article)
			time.Sleep(time.Second)
		}
	}()

	// Let the program run for a moment to see the output
	time.Sleep(6 * time.Second)

	// Deregister subscriber2
	newsService.Deregister(subscriber2)

	// Simulate publishing another article after deregistration
	article := "Final Article"
	newsService.Notify(article)

	// Let the program run for a moment to see the final output
	time.Sleep(time.Second)
}

```
Advantages of the Observer Pattern:

    Loose coupling: The subject  and observers they don't need to know each other's details beyond the Observer interface.

    Supports multiple observers: The pattern allows for any number of observers to be registered with a subject, and they can all react independently to changes.

    Ease of extension: You can easily add new observers without modifying the subject or other observers.

When to Use the Observer Pattern:

    Event-driven systems: When changes in one object require updates in other related objects (e.g., UI components reflecting changes in underlying data).

    Decoupling behavior: When you want to decouple an abstraction from its implementation so that the two can vary independently.


- Decorator Pattern: The Decorator pattern allows behavior to be added to individual objects, dynamically, without affecting the behavior of other objects from the same class. 

In this example, the Decorator pattern allows us to dynamically add behaviors (enhancements to attack or defense) to an object (the player) without affecting its core functionality. This pattern is useful when you need to extend an object's functionality in a flexible and modular way, which is especially relevant in game development and other software scenarios where objects can have varied and dynamic behaviors.

```go

package main

import "fmt"

// Player interface represents the basic functionalities a player must have.
type Player interface {
	Attack() int
	Defense() int
}

// BasePlayer represents the basic attributes of a player.
type BasePlayer struct {
	attack  int
	defense int
}

func (p *BasePlayer) Attack() int {
	return p.attack
}

func (p *BasePlayer) Defense() int {
	return p.defense
}

// ArmorDecorator enhances a player's defense.
// ArmorDecorator enhances a player's defense by adding an additional armor value.
type ArmorDecorator struct {
	player Player
	armor  int
}

func (a *ArmorDecorator) Attack() int {
	return a.player.Attack()
}

func (a *ArmorDecorator) Defense() int {
	return a.player.Defense() + a.armor
}

// WeaponDecorator enhances a player's attack.
// WeaponDecorator enhances a player's attack by adding an additional attack value.
type WeaponDecorator struct {
	player Player
	attack int
}

func (w *WeaponDecorator) Attack() int {
	return w.player.Attack() + w.attack
}

func (w *WeaponDecorator) Defense() int {
	return w.player.Defense()
}

func main() {
	// Create a base player
	player := &BasePlayer{
		attack:  10,
		defense: 5,
	}

	fmt.Println("Base Player:")
	fmt.Printf("Attack: %d, Defense: %d\n", player.Attack(), player.Defense())

	// Add armor to the player
	playerWithArmor := &ArmorDecorator{
		player: player,
		armor:  5,
	}

	fmt.Println("Player with Armor:")
	fmt.Printf("Attack: %d, Defense: %d\n", playerWithArmor.Attack(), playerWithArmor.Defense())

	// Add a weapon to the player with armor
	playerWithArmorAndWeapon := &WeaponDecorator{
		player: playerWithArmor,
		attack: 8,
	}

	fmt.Println("Player with Armor and Weapon:")
	fmt.Printf("Attack: %d, Defense: %d\n", playerWithArmorAndWeapon.Attack(), playerWithArmorAndWeapon.Defense())
}
```
```bash
# OUTPUT
Base Player:
Attack: 10, Defense: 5
Player with Armor:
Attack: 10, Defense: 10
Player with Armor and Weapon:
Attack: 18, Defense: 10
```
- Strategy Pattern: The Strategy pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable. It lets the algorithm vary independently from clients that use it. The strategy pattern is a behavioral design pattern that enables selecting an algorithm at runtime from a family of algorithms. It allows a client to choose from a variety of algorithms or strategies without altering their structure. This pattern is useful when you have multiple algorithms that can be interchangeable depending on the context.

```go
// Defines a contract for sorting algorithms. Any concrete sorting algorithm (BubbleSort, QuickSort) must implement this interface.
type SortStrategy interface {
    Sort([]int) []int
}
```
```go
ype BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}
```

```go
type QuickSort struct{}

func (qs *QuickSort) Sort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    pivot := arr[0]
    var less, greater []int
    for _, v := range arr[1:] {
        if v <= pivot {
            less = append(less, v)
        } else {
            greater = append(greater, v)
        }
    }
    sorted := append(append(qs.Sort(less), pivot), qs.Sort(greater)...)
    return sorted
}

```
```go
// Holds a reference to a strategy object (SortStrategy) and provides methods to set and use different sorting algorithms. It delegates the sorting task to the current strategy.
type SortContext struct {
    strategy SortStrategy
}

func NewSortContext(strategy SortStrategy) *SortContext {
    return &SortContext{strategy: strategy}
}

func (sc *SortContext) SetStrategy(strategy SortStrategy) {
    sc.strategy = strategy
}

func (sc *SortContext) SortArray(arr []int) []int {
    return sc.strategy.Sort(arr)
}

```
```go
func main() {
    arr := []int{64, 25, 12, 22, 11}

    // Use Bubble Sort
    bubbleSort := &BubbleSort{}
    context := NewSortContext(bubbleSort)
    sorted := context.SortArray(arr)
    fmt.Println("Sorted using Bubble Sort:", sorted)

    // Use Quick Sort
    quickSort := &QuickSort{}
    context.SetStrategy(quickSort)
    sorted = context.SortArray(arr)
    fmt.Println("Sorted using Quick Sort:", sorted)
}
```
The strategy pattern allows the client (main function) to select different algorithms dynamically at runtime.
Algorithms are encapsulated in their own classes (BubbleSort, QuickSort), adhering to the Single Responsibility Principle.
Adding new sorting algorithms (MergeSort, InsertionSort, etc.) involves creating new classes that implement SortStrategy.

# File Operations
```go
package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Writer goroutine: sends log messages into channel
func writer(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Writer %d: log entry %d\n", id, i)
		ch <- msg
		time.Sleep(100 * time.Millisecond) // simulate delay
	}
}

// File writer goroutine: consumes channel and writes to file
func fileWriter(file *os.File, ch <-chan string, done chan<- struct{}) {
	for msg := range ch {
		_, err := file.WriteString(msg)
		if err != nil {
			fmt.Println("Write error:", err)
		}
	}
	// signal completion
	done <- struct{}{}
}

func main() {
	// Open file for append
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Channels
	logCh := make(chan string, 100) // buffered channel
	done := make(chan struct{})

	// Start file writer goroutine
	go fileWriter(file, logCh, done)

	var wg sync.WaitGroup

	// Start multiple writers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go writer(i, logCh, &wg)
	}

	// Wait for writers to finish
	wg.Wait()
	close(logCh) // close channel so fileWriter exits

	// Wait for fileWriter to finish
	<-done

	fmt.Println("All logs written to file.")
}
```
```go
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

```
## Concurent File Write and Read
```go
package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Writer: Appends logs until its work is done
func tailWriter(file *os.File, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 200; i++ {
		_, err := file.WriteString(fmt.Sprintf("Writer %d: log entry %d\n", id, i))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
		time.Sleep(50 * time.Millisecond)
		// After each write, it sleeps for 50 milliseconds to mimic real-world logging where entries don’t come instantly.
	}
}

// Tail Reader: Monitors new entries until the context is cancelled . This function continuously monitors the file for new log entries until it is told to stop.
func tailReader(ctx context.Context, file *os.File, name string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// It uses a buffered reader (bufio.NewReader) and maintains an offset to track how far it has read in the file.
	// Use a persistent reader and offset tracking
	reader := bufio.NewReader(file)
	offset := int64(0)

	for { // Inside an infinite loop, it checks if the context is cancelled. If so, it prints a shutdown message and exits. Otherwise, it seeks to the current offset in the file and tries to read a line.
		select {
		case <-ctx.Done():
			fmt.Printf("Reader %s (%d) received shutdown signal.\n", name, id)
			return
		default:
			// Ensure we are at the right position
			_, err := file.Seek(offset, io.SeekStart)
			if err != nil {
				return
			}
			// If a line is read, it prints it with the reader’s name and ID, then updates the offset.
			line, err := reader.ReadString('\n')
			if len(line) > 0 {
				fmt.Printf("[%s Reader %d] %s", name, id, line)
				offset += int64(len(line))
				// Reset reader state after seek to ensure next read is fresh
				reader.Reset(file) // After seeking, the reader resets its buffer to ensure fresh reads.
			}

			if err == io.EOF { // If it hits io.EOF, meaning no new data is available, it sleeps briefly and retries.
				// Wait for more data or shutdown signal
				time.Sleep(100 * time.Millisecond)
				continue
			}
			if err != nil {
				fmt.Printf("Reader %s error: %v\n", name, err)
				return
			}
		}
	}
}

func main() {
	logFile := "server.log"
	// Opens the file for writing with flags: append, write-only, create if not exists, and truncate (start fresh).
	writeFile, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer writeFile.Close()

	// Separate handles for readers to maintain independent offsets. Creates two separate file handles for reading (readFile1 and readFile2). This is important because each reader maintains its own offset independently.
	readFile1, _ := os.Open(logFile)
	defer readFile1.Close()
	readFile2, _ := os.Open(logFile)
	defer readFile2.Close()

	// Context for managing Reader lifecycles. Sets up a context with cancellation (ctx, cancel := context.WithCancel(...)) to control when readers should stop.
	ctx, cancel := context.WithCancel(context.Background())
	// Defines two WaitGroups: one for writers and one for readers.
	var writerWg sync.WaitGroup
	var readerWg sync.WaitGroup
	/*
		The program launches two writer goroutines.

		    Each writer is added to the writerWg.

		    They run concurrently, writing logs into the same file.

		    This simulates multiple processes writing logs simultaneously.
	*/
	// 1. Start Writers
	for i := 1; i <= 2; i++ {
		writerWg.Add(1)
		go tailWriter(writeFile, i, &writerWg)
	}
	/*
		Two reader goroutines are launched.

		    Each reader is added to the readerWg.

		    They monitor the file independently, printing out new log entries as they appear.

		    The names "Alpha" and "Beta" are used to distinguish them.

		This simulates multiple monitoring processes tailing the same log file.

	*/
	// 2. Start Tail Readers
	readerWg.Add(1)
	go tailReader(ctx, readFile1, "Alpha", 1, &readerWg)
	readerWg.Add(1)
	go tailReader(ctx, readFile2, "Beta", 2, &readerWg)

	// 3. Wait for writers to finish their work
	writerWg.Wait()
	fmt.Println(">>> Writers finished. Waiting 1s for readers to catch up...")
	time.Sleep(1 * time.Second)
	/*
		    The program waits for all writers to finish (writerWg.Wait()).

		    Once writers are done, it prints a message and sleeps for 1 second to give readers time to catch up.

		    Then it calls cancel() to signal the readers to stop.

		    Finally, it waits for all readers to finish (readerWg.Wait()).

		    Prints a clean exit message.

		This ensures that the program exits gracefully, with no dangling goroutines.
	*/
	// 4. Signal readers to stop and wait for them
	cancel()
	readerWg.Wait()
	fmt.Println("Main: System exited cleanly.")
}

```
```go
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

```

# Status Code

			200 OK: The request was successful, and the server returned the requested data.

			201 Created: The request was successful, and as a result, a new resource was created (commonly used with POST).

			202 Accepted: The request has been accepted for processing, but the processing is not complete.

			400 Bad Request: The server could not understand the request due to invalid syntax.

			401 Unauthorized: The client must authenticate itself to get the requested response (missing or invalid authentication credentials).

			403 Forbidden: The server understood the request but refuses to authorize it (the client does not have permission).

			404 Not Found: The server could not find the requested resource.

			405 Method Not Allowed: The method specified in the request is not allowed for the resource.

			406 Not Acceptable: The server cannot generate a response that is acceptable according to the Accept headers.

			408 Request Timeout: The client did not send a request in time.

			500 Internal Server Error: The server encountered an unexpected condition that prevented it from fulfilling the request.

			501 Not Implemented: The server does not support the functionality required to fulfill the request.

			502 Bad Gateway: The server, while acting as a gateway or proxy, received an invalid response from the upstream server.

			503 Service Unavailable: The server is currently unable to handle the request (e.g., due to overload or maintenance).

			504 Gateway Timeout: The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server.

			505 HTTP Version Not Supported: The server does not support the HTTP protocol version that was used in the request.

			511 Network Authentication Required: The client needs to authenticate to gain network access.

In HTTP, PUT replaces the entire resource with the new data, while PATCH applies partial updates to only the specified fields. PUT is always idempotent (repeating the same request has the same effect), whereas PATCH may not be idempotent depending on implementation
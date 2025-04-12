
/*
1,4,7
2,5,8
3,6,9
*/

```go
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


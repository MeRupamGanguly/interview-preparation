package main

import "fmt"

func cleanup(i int) {
	fmt.Println("cleanup func:", i)
}

func main() {
	i := 10
	defer cleanup(i)

	i++

	defer func() {
		fmt.Println("anonymous func:", i)
	}()

	i = 100
}

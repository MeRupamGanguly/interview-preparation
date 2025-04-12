package main

import "fmt"

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

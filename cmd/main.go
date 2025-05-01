package main

import (
	"fmt"
)

func main() {
	a := func(i int) int { return i * i } // A function a that takes an int and returns its square.
	v := func(f func(i int) int) func() int {
		fmt.Println("Inside : Address of f", &f) // This is the address of the function parameter f, which is a local copy of the function a passed into v.
		c := f(3)
		return func() int {
			fmt.Println("Inside : Address of c", &c, " Value ", c) // This is the address of the local variable c inside v. When the closure is returned, it captures c and keeps it in its own memory space.
			c++
			return c
		}
	}
	f := v(a)                                     // v() returns a function address
	fmt.Println("Inside f is ", f)                // This is the address of the returned closure function, func() int. It’s a separate function, dynamically created and returned by v.
	fmt.Println("f Function Called and get", f()) // call the function which v() returns // When f() is called, it returns c+1, showing that the closure is preserving the state of c between invocations (i.e., it's a closure over c).
}

/*
Inside : Address of f 0xc000066040
Inside f is  0x4b1f40
Inside : Address of c 0xc000010120  Value  9
f Function Called and get 10
*/

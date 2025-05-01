package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{2, 0, 6, 0, 9, 12, 0, 46}
	x := Zero2Front(arr)
	arr = []int{2, 0, 6, 0, 9, 12, 0, 46}
	y := Zero2FrontOrdered(arr)
	fmt.Println(x, y)
}
func Zero2FrontOrdered(arr []int) []int {
	l := 0
	r := 0
	for range arr {
		if arr[r] == 0 {

			if r-l == 1 {
				arr[r] = arr[l]
				arr[l] = 0
			} else {
				rr := r
				for range arr[l : rr+1] {
					arr[rr] = arr[rr-1]
					rr--
				}
				arr[l] = 0
			}
			l++
		}
		r++
	}
	return arr
}
func Zero2Front(arr []int) []int {
	l := 0
	r := 0
	loopCount := 0
	for range arr {
		loopCount++
		if arr[r] == 0 {
			arr[r] = arr[l]
			arr[l] = 0
			l++
		}
		r++
	}
	log.Println("Iteration Number", loopCount)
	return arr
}

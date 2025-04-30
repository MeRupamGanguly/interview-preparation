package main

import "log"

func main() {
	arr := []int{12, 14, 15, 1, 16, 5, 13, 18}
	log.Println("Array is ", arr)
	w := 4
	l := 0
	r := w - 1
	maxSum := 0
	sum := func(a []int) int {
		log.Println("Sum of ", a)
		sum := 0
		for i := range a {
			sum = sum + a[i]
		}
		return sum
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for r <= len(arr) {
		s := sum(arr[l:r])
		l++
		r++
		maxSum = max(maxSum, s)
		log.Println("Sum is ", s)
	}
	log.Println("MAX SUM IS ", maxSum)
}

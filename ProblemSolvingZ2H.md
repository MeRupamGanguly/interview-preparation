# Problem Solving Zero to Hero

## Pattern 1 : Two Pointer OR Window Pattern
Two Pointer is a technique where you use two indices (pointers) to traverse a data structure, usually arrays or strings, to perform optimized searches or comparisons. We use when Finding pairs/triplets in sorted arrays, Reversing sequences, Merging sorted arrays. 

### Given a sorted array, find two numbers that add up to a given target.

Use two pointers: left at the start and right at the end.

```go
func TwoSum(arr []int, target int) (int, int) {
	l := 0
	r := len(arr) - 1
	loopCount := 0
	for l < r {
		loopCount++
		sum := arr[l] + arr[r]
		if target == sum {
			log.Println("Iteration Number", loopCount)
			return l, r
		} else {
			if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return -1, -1
}
```
### Reverse a Slice.

```go
func Reverse(arr []int) []int {
	l := 0
	r := len(arr) - 1
	loopCount := 0
	for l < r {
		loopCount++
		t := arr[l]
		arr[l] = arr[r]
		arr[r] = t
		l++
		r--
	}
	log.Println("Iteration Number", loopCount)
	return arr
}
```

### Move 0s to front
```go
func main() {
	arr := []int{2, 0, 6, 0, 9, 12, 0, 46}
	x := Zero2Front(arr)
	arr = []int{2, 0, 6, 0, 9, 12, 0, 46}
	y := Zero2FrontOrdered(arr)
	fmt.Println(x, y)
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
```
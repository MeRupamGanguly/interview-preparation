
# Problem-Solving Patterns

| Pattern Name               | What It’s For                           | Example Problems                            |
|---------------------------|-----------------------------------------|---------------------------------------------|
| 1. **Two Pointers**       | Scan from two ends / two speeds         | Palindrome, Pair Sum, Remove Duplicates     |
| 2. **Sliding Window**     | Work with subarrays/substrings          | Longest Substring, Max Sum in Range         |
| 3. **DFS**                 | Go deep into trees/graphs               | Maze Solver, Binary Tree Paths              |
| 4. **BFS**                 | Explore level by level                  | Shortest Path, Social Network               |
| 5. **Backtracking**        | Try all options (like guessing)         | Sudoku, N-Queens, Word Search               |
| 6. **Binary Search**       | Search sorted data fast                 | Number Finder, First/Last Position          |
| 7. **Heap / Priority Queue** | Get top/bottom items fast             | K Largest Numbers, Task Scheduling          |
| 8. **Dynamic Programming** | Break big problems into small parts     | Fibonacci, Knapsack, Coin Change            |

---
# Two Pointers — "Scaning with Two Fingers":

Imagine we have an Sorted Array, need to check if there exist any pair of numbers such that adding those number give a target k.

```go
func TwoSum(arr []int, target int) bool {
    for i := 0; i < len(arr); i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i]+arr[j] == target {
                return true
            }
        }
    }
    return false
}
```
This is O(n²) Solution as we use 2 nested Loops. The outer loop picks one element `a[i]` at a time from the array. For each of these elements `a[i]`, the inner loop goes through the remaining elements `a[j]` ahead of it to form every possible pair `(a[i],a[j])` and check their sum `arr[i]+arr[j] == target` .

Now to optimize it:
Start looking with two pointers: one at the beginning, and one at the end. Based on the sum of their values, move the pointers inward to get closer to the target without checking every possible pair. This is O(n) Solution. This approach only works if the array is sorted.

```go
sort.Ints(arr) // Sort the array if Sorted array not given.
```
```go
func hasPairWithSum(arr []int, target int) bool {
    left := 0
    right := len(arr) - 1
    for left < right {
        sum := arr[left] + arr[right]
        if sum == target {
            return true
        } else if sum < target {
            left++ // Increase sum by moving left pointer right
        } else {
            right-- // Decrease sum by moving right pointer left
        }
    }
    return false
}
```

So in Two Pointers Pattern we have to Figure out 2 Satge:
1. Pointers Initiallization. 
2. Pointers Moving Conditions.


Use **two fingers (pointers)** to walk through an array or list by these possible scenerio. 
- Start from **both ends** You put one pointer at the beginning and the other at the end of the Array. You move them toward each other for checking any specific condtion matched or not. Also you need a Condition after that the Checking is stoped(like after all elements checked of that array).

Like To check if a string is a palindrome, compare characters from the beginning and the end, moving towards the center. If any pair of characters doesn't match, it's not a palindrome; otherwise, it is.

```go
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1 // Pointers Initiallization

	for left < right { // specific condition after that the Checking is stoped
		if s[left] != s[right] { // specific condtion matched or not
			return false
		}
		left++
		right--
	}
	return true
}
```

- Move both **left to right** Both pointers start at the beginning, and move forward through the list. One goes faster (called fast), the other goes slower or only moves when (specific condition satisfied)needed (called slow). Like scanning a paper: one eye reads every word (fast), the other keeps track of important stuff (slow). 

```go
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
```
- Move one **fast**, one **slow** One pointer (fast) moves twice as fast as the other (slow). This is like a race between a rabbit and a turtle. If there's a cycle (loop), they will eventually meet. If there's no loop, the fast one will reach the end first.
```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
```

| Type                       | How It Works                       | Use It When...                         | Example                            |
|----------------------------|------------------------------------|----------------------------------------|------------------------------------|
| **Opposite Direction**   | One from start, one from end       | Sorted data or symmetry                | Pair Sum, Palindrome               |
| **Same Direction**      | Both move left to right            | Filter/compress/change in-place        | Remove Duplicates, Merge Arrays    |
| **Fast & Slow**         | One moves faster                   | Find cycles or middle                  | Linked List Cycle, Middle Node     |


1. Two Sum II - Input Array Is Sorted

Problem: Given a sorted array of integers and a target value, return the 1-based indices of the two numbers such that they add up to the target.

Logic:
Use two pointers (left and right) starting from opposite ends of the sorted array.
If their sum equals the target, return indices. If the sum is less, move left forward. If more, move right backward.

```go
func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left + 1, right + 1}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    return nil // Guaranteed to find a solution, so this line won't be reached
}

```
2. Valid Palindrome

Problem: Check if a given string is a palindrome, considering only alphanumeric characters and ignoring case.

Logic:
Use two pointers (left and right) to skip non-alphanumeric characters and compare lowercase versions.
Move inward; if characters don't match, it's not a palindrome.

```go
func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
            left++
        } else if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
            right--
        } else if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
            return false
        } else {
            left++
            right--
        }
    }
    return true
}

```

3. Remove Duplicates from Sorted Array

Problem: Remove duplicates in-place from a sorted array and return the new length.

Logic:
Use a slow pointer to mark the position of the next unique number.
Move the fast pointer to scan the array and copy over new values when a change is detected.

```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    slow := 1
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[fast-1] {
            nums[slow] = nums[fast]
            slow++
        }
    }
    return slow
}
```

4. Move Zeroes

Problem: Move all zero elements to the end of the array while maintaining the relative order of non-zero elements.

Logic:
Use a pointer lastNonZero to track where the next non-zero element should go.
Swap non-zero elements into place and let zeroes naturally move to the end.

```go
func moveZeroes(nums []int) {
    lastNonZero := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[lastNonZero], nums[i] = nums[i], nums[lastNonZero]
            lastNonZero++
        }
    }
}
```

5. Longest Substring Without Repeating Characters

Problem: Find the length of the longest substring without repeating characters.

Logic:
Use a sliding window and a map to track character positions.
If a repeat is found, shift the window’s left edge and update the max length.

```go
func lengthOfLongestSubstring(s string) int {
    charIndexMap := make(map[byte]int)
    left, maxLength := 0, 0
    for right := 0; right < len(s); right++ {
        if index, found := charIndexMap[s[right]]; found && index >= left {
            left = index + 1
        }
        charIndexMap[s[right]] = right
        if right-left+1 > maxLength {
            maxLength = right - left + 1
        }
    }
    return maxLength
}
```

6. Linked List Cycle

Problem: Detect if a linked list has a cycle.

Logic:
Use slow and fast pointers; the fast pointer moves twice as fast as the slow.
If there's a cycle, they will eventually meet inside the loop.

```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

7. Middle of the Linked List

Problem: Return the middle node of a singly linked list.

Logic:
Again use slow and fast pointers. Fast moves 2 steps, slow moves 1 step.
When fast reaches the end, slow is at the middle.

```go
func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    while fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
```

8. Sort Colors

Problem: Given an array with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent.

Logic:
This is the Dutch National Flag problem. Use three pointers: left for 0s, right for 2s, and current to traverse.
Swap values to their respective regions in-place.

```go
func sortColors(nums []int) {
	left, right, current := 0, len(nums)-1, 0
	for current <= right {
		switch nums[current] {
		case 0:
			nums[left], nums[current] = nums[current], nums[left]
			left++
			current++
		case 1:
			current++
		case 2:
			nums[right], nums[current] = nums[current], nums[right]
			right--
		}
	}
}
```

9. Container With Most Water

Problem: Given n non-negative integers representing the height of walls, find two walls that together with the x-axis forms a container that holds the most water.

Logic:
Use two pointers from both ends. Calculate area = width × min(height).
Move the shorter wall inward to potentially find a larger area.

```go
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		width := right - left
		h := min(height[left], height[right])
		area := width * h
		maxArea = max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

# Sliding Window Pattern

The Sliding Window pattern involves using two pointers to represent a "window" over a subset of elements in a data structure (like an array or string). By adjusting these pointers, you can efficiently explore all possible contiguous subarrays or substrings that satisfy certain conditions.

## 🔍 Types of Sliding Window

| Type                  | Description                                                                 | Use Cases                                           | Example Problem                                      |
|-----------------------|------------------------------------------------------------------------------|-----------------------------------------------------|------------------------------------------------------|
| Fixed Size Window     | Window size is predetermined (e.g., `k` elements).                          | Maximum/Minimum sum, average, etc.                 | Maximum sum of subarray of size `k`                  |
| Variable Size Window  | Window size adjusts dynamically based on conditions (e.g., distinct items). | Longest/shortest subarray based on certain criteria | Longest substring with at most `k` distinct characters |

When to Use Sliding Window

- Contiguous Subarrays/Substrings: Problems requiring analysis of consecutive elements.

- Optimization: When a brute-force solution involves nested loops, sliding window can reduce time complexity to O(n).

- Dynamic Conditions: When the subset size or content needs to adjust based on certain criteria.

Common Problem Patterns

- Fixed Size Window: Calculate sum, average, or other metrics over a subarray of fixed length.

- Variable Size Window: Find the longest/shortest subarray that meets a condition (e.g., sum ≤ k, at most k distinct characters).

- Two Pointer Technique: Often used in conjunction with sliding window to handle problems like palindrome checking or cycle detection.


Maximum Sum Subarray of Size k
Given an array of integers and a number k, find the maximum sum of any contiguous subarray of size k.


```go
func maxSumSubarray(arr []int, k int) int {
    if len(arr) < k {
        return -1 // Invalid input
    }
    var maxSum, windowSum int
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }
    maxSum = windowSum
    for i := k; i < len(arr); i++ {
        windowSum += arr[i] - arr[i-k]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }
    return maxSum
}
```

Longest Substring with At Most k Distinct Characters
Given a string and an integer k, find the length of the longest substring that contains at most k distinct characters.

```go
func longestSubstringKDistinct(s string, k int) int {
    charCount := make(map[byte]int)
    left, maxLength := 0, 0
    for right := 0; right < len(s); right++ {
        charCount[s[right]]++
        for len(charCount) > k {
            charCount[s[left]]--
            if charCount[s[left]] == 0 {
                delete(charCount, s[left])
            }
            left++
        }
        if right-left+1 > maxLength {
            maxLength = right - left + 1
        }
    }
    return maxLength
}
```

Subarray with Sum Less Than or Equal to k
Given an array of positive integers and a number k, find the length of the smallest subarray whose sum is greater than or equal to k.

```go
func minSubArrayLen(k int, nums []int) int {
    left, sum, minLength := 0, 0, len(nums)+1
    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= k {
            if right-left+1 < minLength {
                minLength = right - left + 1
            }
            sum -= nums[left]
            left++
        }
    }
    if minLength == len(nums)+1 {
        return 0
    }
    return minLength
}
```
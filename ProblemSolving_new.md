```go
func reverseString(s string) string {
    runes := []rune(s)
    left, right := 0, len(runes)-1

    for left < right {
        runes[left], runes[right] = runes[right], runes[left]
        left++
        right--
    }

    return string(runes)
}
```
```go
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
```
```go
func findMissingNumber(nums []int) int {
    n := len(nums) + 1
    expectedSum := n * (n + 1) / 2
    actualSum := 0

    for _, num := range nums {
        actualSum += num
    }

    return expectedSum - actualSum
}
```
```go
func findDuplicate(nums []int) int {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] {
            return num
        }
        seen[num] = true
    }
    return -1 // return -1 if no duplicate exists
}
```
```go
func isPalindrome(s string) bool {
    runes := []rune(s)
    left, right := 0, len(runes)-1

    for left < right {
        if runes[left] != runes[right] {
            return false
        }
        left++
        right--
    }
    return true
}
```
```go
func findKthLargest(nums []int, k int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))
    return nums[k-1]
}
```
```go
func areAnagrams(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    count := make(map[rune]int)
    for _, ch := range s1 {
        count[ch]++
    }
    for _, ch := range s2 {
        count[ch]--
        if count[ch] < 0 {
            return false
        }
    }
    return true
}
```
```go
// Implement a Stack Using Queues

type MyStack struct {
    queue1 []int
    queue2 []int
}

func Constructor() MyStack {
    return MyStack{queue1: []int{}, queue2: []int{}}
}

func (this *MyStack) Push(x int) {
    this.queue1 = append(this.queue1, x)
}

func (this *MyStack) Pop() int {
    if len(this.queue1) == 0 {
        return -1
    }
    for len(this.queue1) > 1 {
        this.queue2 = append(this.queue2, this.queue1[0])
        this.queue1 = this.queue1[1:]
    }
    popVal := this.queue1[0]
    this.queue1 = this.queue2
    this.queue2 = []int{}
    return popVal
}

func (this *MyStack) Top() int {
    if len(this.queue1) == 0 {
        return -1
    }
    return this.queue1[len(this.queue1)-1]
}

func (this *MyStack) Empty() bool {
    return len(this.queue1) == 0
}
```
```go
func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    
    prev := countAndSay(n - 1)
    result := ""
    count := 1

    for i := 1; i < len(prev); i++ {
        if prev[i] == prev[i-1] {
            count++
        } else {
            result += fmt.Sprintf("%d%c", count, prev[i-1])
            count = 1
        }
    }

    result += fmt.Sprintf("%d%c", count, prev[len(prev)-1])
    return result
}
```
```go
func permute(s string) []string {
    var res []string
    var backtrack func(s []rune, start int)

    backtrack = func(s []rune, start int) {
        if start == len(s) {
            res = append(res, string(s))
            return
        }

        for i := start; i < len(s); i++ {
            s[start], s[i] = s[i], s[start]
            backtrack(s, start+1)
            s[start], s[i] = s[i], s[start] // backtrack
        }
    }

    backtrack([]rune(s), 0)
    return res
}
```
```go
func intersection(nums1 []int, nums2 []int) []int {
    result := []int{}
    map1 := make(map[int]bool)

    for _, num := range nums1 {
        map1[num] = true
    }

    for _, num := range nums2 {
        if map1[num] {
            result = append(result, num)
            map1[num] = false // To avoid duplicates in the result
        }
    }

    return result
}
```


```go
func isValid(s string) bool {
    stack := []rune{}

    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 {
                return false
            }
            top := stack[len(stack)-1]
            if (char == ')' && top == '(') || (char == '}' && top == '{') || (char == ']' && top == '[') {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}
```
```go
func reverse(x int) int {
    result := 0
    for x != 0 {
        digit := x % 10
        result = result*10 + digit
        x /= 10
    }

    if result > (1<<31)-1 || result < -(1<<31) {
        return 0
    }
    return result
}
```

```go
func isPowerOfThree(n int) bool {
    if n <= 0 {
        return false
    }

    for n%3 == 0 {
        n /= 3
    }
    return n == 1
}
```
```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    index := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}
```
```go
func firstUniqChar(s string) int {
    charCount := make(map[rune]int)

    for _, char := range s {
        charCount[char]++
    }

    for i, char := range s {
        if charCount[char] == 1 {
            return i
        }
    }
    return -1
}
```
```go
// Maximum Subarray (Kadane's Algorithm)

// Problem: Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

func maxSubArray(nums []int) int {
    maxSoFar := nums[0]
    maxEndingHere := nums[0]

    for i := 1; i < len(nums); i++ {
        maxEndingHere = max(nums[i], maxEndingHere+nums[i])
        maxSoFar = max(maxSoFar, maxEndingHere)
    }

    return maxSoFar
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    fmt.Println(maxSubArray(nums)) // Output: 6
}
```



## 1. Two Pointers: The Buddy System 👫
A technique where two pointers traverse a data structure (usually an array or string) in a coordinated way to solve problems efficiently.

Imagine you and your friend are searching for treasure in a sorted list. 
You start at the left, your friend start at the right and walk toward each other, checking items as you go. 


Two pointers start at different positions.
Or They can be start from same Position then they move differently according specific conditions.
Or They can start at opposite ends and meet in the middle. 

Efficiency: O(n) time, O(1) space.

1. Two Sum (Sorted Array) - Asked by Google (50+ times), Amazon (30+ times)

Problem: Given a sorted array of integers and a target, find two numbers that add up to the target.

Explanation: We use two pointers starting from both ends. If the sum is too small, we move the left pointer right. If too large, we move the right pointer left.

```go
func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1
    
    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            return []int{left+1, right+1} // 1-based index
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    return []int{-1, -1} // not found
}
```

```go
func twoSum(nums []int, target int) []int {
    mapNums := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if index, found := mapNums[complement]; found {
            return []int{index, i}
        }
        mapNums[num] = i
    }
    return nil
}
```

2. Remove Duplicates from Sorted Array - Asked by Microsoft (40+ times), Facebook (30+ times)

Problem: Given a sorted array, remove duplicates in-place and return the new length.

Explanation: The slow pointer tracks the position of the last unique element, while fast explores the array. When we find a new unique element, we move it next to slow.

```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    slow := 0
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    return slow + 1
}
```

3. 3Sum - Asked by Google (45+ times), Facebook (35+ times)

Problem: Find all unique triplets in the array that sum to zero.

Explanation: We fix one number (i) and then use two pointers (left and right) to find pairs that sum to -nums[i]. Sorting helps us avoid duplicates efficiently.

```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue // skip duplicates
        }
        
        left, right := i+1, len(nums)-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                // Skip duplicates
                for left < right && nums[left] == nums[left-1] {
                    left++
                }
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }
    return result
}
```

4. Valid Palindrome - Asked by Apple (30+ times), LinkedIn (25+ times)

Problem: Given a string, determine if it's a palindrome after converting to lowercase and removing non-alphanumeric characters.

Explanation: We use two pointers starting from both ends, skipping non-alphanumeric characters, and comparing characters case-insensitively until they meet in the middle.

```go
func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    
    for left < right {
        // Skip non-alphanumeric from left
        for left < right && !isAlphanumeric(s[left]) {
            left++
        }
        // Skip non-alphanumeric from right
        for left < right && !isAlphanumeric(s[right]) {
            right--
        }
        
        if toLower(s[left]) != toLower(s[right]) {
            return false
        }
        left++
        right--
    }
    return true
}

func isAlphanumeric(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return c + 32
    }
    return c
}
```
5. Trapping Rain Water - Asked by Facebook (40+ times), Google (30+ times)

Problem: Given n non-negative integers representing an elevation map, compute how much water it can trap after raining.

Explanation: We track the maximum heights from both ends. The water trapped at any point depends on the minimum of the two maximums.


```go
func trap(height []int) int {
    if len(height) < 3 {
        return 0
    }
    
    left, right := 0, len(height)-1
    leftMax, rightMax := height[left], height[right]
    water := 0
    
    for left < right {
        if leftMax < rightMax {
            left++
            leftMax = max(leftMax, height[left])
            water += leftMax - height[left]
        } else {
            right--
            rightMax = max(rightMax, height[right])
            water += rightMax - height[right]
        }
    }
    return water
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

6. Move Zeroes - Asked by Amazon (35+ times), Microsoft (25+ times)

Problem: Move all 0's to the end while maintaining the relative order of non-zero elements.

Explanation: The slow pointer maintains the position for the next non-zero element, while fast finds non-zero elements to swap.

```go
func moveZeroes(nums []int) {
    slow := 0
    for fast := 0; fast < len(nums); fast++ {
        if nums[fast] != 0 {
            nums[slow], nums[fast] = nums[fast], nums[slow]
            slow++
        }
    }
}
```

7. Squares of a Sorted Array - Asked by Apple (30+ times), Uber (20+ times)

Problem: Given a sorted array in non-decreasing order, return squares of each number also in sorted order.

Explanation: We fill the result array from the end, comparing absolute values from both ends to handle negative numbers properly.

```go
func sortedSquares(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    left, right := 0, n-1
    
    for i := n-1; i >= 0; i-- {
        if abs(nums[left]) > abs(nums[right]) {
            result[i] = nums[left] * nums[left]
            left++
        } else {
            result[i] = nums[right] * nums[right]
            right--
        }
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

8. Longest Substring Without Repeating Characters - Asked by Google (50+ times), Amazon (40+ times)

Problem: Find the length of the longest substring without repeating characters.

Explanation: We maintain a sliding window using two pointers and a set to track characters in the current window.

```go
func lengthOfLongestSubstring(s string) int {
    charSet := make(map[byte]bool)
    left := 0
    maxLength := 0
    
    for right := 0; right < len(s); right++ {
        for charSet[s[right]] {
            delete(charSet, s[left])
            left++
        }
        charSet[s[right]] = true
        if right-left+1 > maxLength {
            maxLength = right - left + 1
        }
    }
    return maxLength
}
```

9. Merge Sorted Arrays - Asked by Microsoft (35+ times), Facebook (30+ times)

Problem: Merge two sorted arrays into one sorted array in O(1) space (nums1 has enough space).

Explanation: We start from the end of both arrays and merge backwards to avoid overwriting elements in nums1.

```go
func merge(nums1 []int, m int, nums2 []int, n int) {
    p1, p2 := m-1, n-1
    p := m + n - 1
    
    for p1 >= 0 && p2 >= 0 {
        if nums1[p1] > nums2[p2] {
            nums1[p] = nums1[p1]
            p1--
        } else {
            nums1[p] = nums2[p2]
            p2--
        }
        p--
    }
    
    // Copy remaining elements from nums2
    for p2 >= 0 {
        nums1[p] = nums2[p2]
        p2--
        p--
    }
}
```

## 2. Sliding Window Pattern
Imagine you're holding a pair of binoculars and scanning a long landscape (array) in front of you. The binoculars let you focus on a fixed section of the view — just a few objects at a time — and you can slide your view across the scene.

Imagine We have a binoculars/window (a subarray) that moves through the array.
We use two pointers, usually called start and end, to mark the boundaries of the window.
The window "slides" by moving the start or end pointer as needed/Satify specific conditions.

start and end both begin at the start of the array. 
Expand the window by moving/increment end → to include more elements. 
Shrink the window by moving/increment start → to exclude elements from the front. 
You adjust the window based on a condition. 
Track the best solution during the sliding process.

Efficiency: O(n) time, O(min(m, n)) space (where m is character set size).


1. Maximum Sum Subarray of Size K (Fixed Window) Asked by: Google (50+), Amazon (40+), Microsoft (30+)
Problem: Given an array of integers and a number k, find the maximum sum of any contiguous subarray of size k.

Explanation:
Fixed-size window of length k.
Slide the window by subtracting the leftmost element (nums[i-k]) and adding the new right element (nums[i]).
Track the maximum sum as the window slides.

```go
func maxSumSubarray(nums []int, k int) int {
    if len(nums) < k {
        return 0
    }

    maxSum, windowSum := 0, 0
    for i := 0; i < k; i++ {
        windowSum += nums[i]
    }
    maxSum = windowSum

    for i := k; i < len(nums); i++ {
        windowSum += nums[i] - nums[i-k] // Slide the window
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }
    return maxSum
}
```

2. Minimum Window Substring (Variable Window) Asked by: Google (50+), Amazon (40+), Microsoft (30+)

Problem: Given strings s and t, find the smallest substring in s that contains all characters of t.


Explanation:
Sliding window expands until all t characters are included (count == len(t)).
Shrinks from the left to find the smallest valid window.
Tracks character frequencies (window and target maps).

```go
func minWindow(s string, t string) string {
    target := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        target[t[i]]++
    }

    left, count, minLen := 0, 0, math.MaxInt32
    result := ""
    window := make(map[byte]int)

    for right := 0; right < len(s); right++ {
        char := s[right]
        if _, exists := target[char]; exists {
            window[char]++
            if window[char] <= target[char] {
                count++
            }
        }

        for count == len(t) { // Valid window found
            if right-left+1 < minLen {
                minLen = right - left + 1
                result = s[left : right+1]
            }
            leftChar := s[left]
            if _, exists := target[leftChar]; exists {
                window[leftChar]--
                if window[leftChar] < target[leftChar] {
                    count--
                }
            }
            left++
        }
    }
    return result
}
```

3. Longest Substring with At Most K Distinct Characters Asked by: Facebook (40+), Uber (30+), Amazon (25+)

Problem: Find the longest substring with at most k distinct characters.

Explanation:
Expands window until >k distinct characters.
Shrinks from left to maintain ≤k distinct characters.
Updates maxLen when a longer valid window is found.


```go
func lengthOfLongestSubstringKDistinct(s string, k int) int {
    freq := make(map[byte]int)
    left, maxLen := 0, 0

    for right := 0; right < len(s); right++ {
        freq[s[right]]++
        for len(freq) > k { // Shrink if more than k distinct chars
            freq[s[left]]--
            if freq[s[left]] == 0 {
                delete(freq, s[left])
            }
            left++
        }
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}
```
4. Maximum Consecutive Ones III (Flip at most K zeros) Asked by: Google (35+), Microsoft (25+), Facebook (20+)

Problem: Given a binary array and k, find the maximum number of consecutive 1s if you can flip at most k 0s.

Explanation:
Tracks zeroCount (number of 0s in the window).
Shrinks window if zeroCount > k.
Updates maxLen for the longest valid window.

```go
func longestOnes(nums []int, k int) int {
    left, maxLen, zeroCount := 0, 0, 0

    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            zeroCount++
        }
        for zeroCount > k { // Shrink if more than k zeros flipped
            if nums[left] == 0 {
                zeroCount--
            }
            left++
        }
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}
```

5. Permutation in String (Anagram Check) Asked by: Amazon (30+), Google (25+), Microsoft (20+)

Problem: Given two strings s1 and s2, return true if s2 contains a permutation of s1.

Explanation:
Fixed-size window of len(s1).
Tracks character counts (s1Count vs s2Count).
Slides window while checking if counts match (matches == 26).

```go
func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    s1Count := make([]int, 26)
    s2Count := make([]int, 26)

    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]-'a']++
        s2Count[s2[i]-'a']++
    }

    matches := 0
    for i := 0; i < 26; i++ {
        if s1Count[i] == s2Count[i] {
            matches++
        }
    }

    left := 0
    for right := len(s1); right < len(s2); right++ {
        if matches == 26 {
            return true
        }

        // Add new right character
        char := s2[right] - 'a'
        s2Count[char]++
        if s2Count[char] == s1Count[char] {
            matches++
        } else if s2Count[char] == s1Count[char]+1 {
            matches--
        }

        // Remove left character
        char = s2[left] - 'a'
        s2Count[char]--
        if s2Count[char] == s1Count[char] {
            matches++
        } else if s2Count[char] == s1Count[char]-1 {
            matches--
        }
        left++
    }
    return matches == 26
}
```

6. Fruit Into Baskets (At Most 2 Types) Asked by: Google (25+), Amazon (20+), Uber (15+)

Problem: Find the longest subarray with at most 2 distinct values (like collecting fruits into 2 baskets).

Explanation:
Tracks fruit types in basket (hashmap).
Shrinks if >2 types.
Updates maxLen for the longest valid window.

```go
func totalFruit(fruits []int) int {
    basket := make(map[int]int)
    left, maxLen := 0, 0

    for right := 0; right < len(fruits); right++ {
        basket[fruits[right]]++
        for len(basket) > 2 { // Shrink if more than 2 types
            basket[fruits[left]]--
            if basket[fruits[left]] == 0 {
                delete(basket, fruits[left])
            }
            left++
        }
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}
```
7. Minimum Size Subarray Sum (Sum ≥ Target) Asked by: Google (40+), Microsoft (30+), Facebook (25+)

Problem: Given an array of positive integers and target, find the smallest subarray with sum ≥ target.

Explanation:
Expands window while sum < target.
Shrinks to find the smallest valid window (sum ≥ target).
Tracks minLen of valid subarrays.
```go
func minSubArrayLen(target int, nums []int) int {
    left, sum, minLen := 0, 0, math.MaxInt32

    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target { // Shrink while sum ≥ target
            if right-left+1 < minLen {
                minLen = right - left + 1
            }
            sum -= nums[left]
            left++
        }
    }
    if minLen == math.MaxInt32 {
        return 0
    }
    return minLen
}
```

## 3. DFS (Depth-First Search)
DFS is a technique used to traverse or search through graphs and trees.
DFS means going deep into one path until you can’t go any further, then going back (backtracking) and trying another path.

You pick a path and keep walking forward.
If you hit a dead end, you turn around.
Then you try a different path you didn’t take before.
You keep doing this until everything is explored.

Start from a node (root or any starting point).
Go deep into one neighbor, then its neighbor, and so on...
When you can’t go deeper, backtrack and try another branch.

Imagine you're going down a path, and you have to keep track of where you've been so you can return to previous places (backtracking).
When you reach a new node, you push it onto the stack (you take a mental note of the place).
If you hit a dead-end or can’t go further, you pop the stack (backtrack) and return to the most recent place you visited, where you can try the next possible path.

Efficiency: O(n) time (visits every node once).


1. Problem: Number of Islands (LeetCode 200)

Problem Statement:
Given a 2D grid of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.

Explanation:
This solution iterates through each cell in the grid. When a '1' (land) is found, it triggers a DFS to mark all connected '1's as visited by changing them to '0's. Each DFS call corresponds to discovering a new island.

```go
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0' // Mark as visited
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}
```

2. Problem: Binary Tree Paths (LeetCode 257)

Problem Statement:
Given the root of a binary tree, return all root-to-leaf paths in any order.

Explanation:
This solution uses a recursive DFS approach to traverse the binary tree. The constructPaths function appends the current node's value to the path string and recurses down to the left and right children. When a leaf node is reached, the current path is added to the result list.


```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}
	constructPaths(root, "", &paths)
	return paths
}

func constructPaths(node *TreeNode, path string, paths *[]string) {
	if node == nil {
		return
	}
	if path != "" {
		path += "->"
	}
	path += strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, path)
	} else {
		constructPaths(node.Left, path, paths)
		constructPaths(node.Right, path, paths)
	}
}
```

3. Problem: Course Schedule (LeetCode 207)

Problem Statement:
There are numCourses courses you have to take, labeled from 0 to numCourses - 1. Some courses have prerequisites. For example, to take course 0 you have to first take course 1, which is expressed as a pair: 0,1 . Given the total number of courses and a list of prerequisite pairs, determine if you can finish all courses.

Explanation:
This solution uses DFS to detect cycles in the course prerequisite graph. Each course is marked as unvisited, visiting, or visited. A cycle is detected if a course is revisited while still being in the visiting state.

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	visited := make([]int, numCourses)

	// Build the graph
	for _, prereq := range prerequisites {
		graph[prereq[0]] = append(graph[prereq[0]], prereq[1])
	}

	var dfs func(int) bool
	dfs = func(course int) bool {
		if visited[course] == 1 {
			return false // Cycle detected
		}
		if visited[course] == 2 {
			return true // Already processed
		}
		visited[course] = 1 // Mark as visiting
		for _, prereq := range graph[course] {
			if !dfs(prereq) {
				return false
			}
		}
		visited[course] = 2 // Mark as visited
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}
```

## 4. BFS (Breadth-First Search)
BFS is a technique used to traverse or search through graphs and trees.
It visits everything close to the starting point first, then goes outwards like a wave. It checks everything closer to the start before going deeper.

It means exploring all the nearby nodes first, before moving on to nodes that are further away.


A queue is like a line at a ticket counter. The first person who gets in line is the first one to be served.

Imagine we're walking down a street full of houses. We start at House A. We want to visit all the neighboring houses around House A first (like House B, House C, and House D).

To keep track of which houses we need to visit, we add them to a queue, like writing them down in the order we plan to visit them.

We enqueue (add) each neighbor of House A (B, C, and D) into the queue, putting them in line to visit next.

Once we finish visiting House A, we dequeue (remove) the first house in line (let’s say House B) and go to visit that one.

After visiting House B, we check if it has any neighbors (let's say House E, F, and G). We enqueue these new neighbors (E, F, G) to the back of the queue.

Now the queue looks like this: [C, D, E, F, G]. This means we will visit House C, then House D, then House E, and so on.

We continue the process: dequeue the next house (House C), visit it, and add any of its neighbors to the queue. Repeat until all the houses are visited.

This method ensures we explore the street level by level. First, we visit all of House A’s neighbors (B, C, D), then we move on to the next set of neighbors (E, F, G), and so on.


Efficiency: O(n) time and space (queue stores all nodes).


1. Number of Islands

Problem:
Given an m x n 2D binary grid grid representing a map of '1's (land) and '0's (water), return the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.

Explanation:
This solution uses BFS to traverse each unvisited land cell ('1'). Starting from each unvisited land cell, BFS explores all connected land cells, marking them as visited by changing them to '0'. Each BFS call corresponds to discovering a new island.


```go
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func bfs(grid [][]byte, i, j int) {
	queue := []struct{ x, y int }{{i, j}}
	grid[i][j] = '0' // Mark as visited
	directions := []struct{ x, y int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			x, y := cell.x+dir.x, cell.y+dir.y
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '1' {
				grid[x][y] = '0' // Mark as visited
				queue = append(queue, struct{ x, y int }{x, y})
			}
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid)) // Output: 1
}
```

2. Rotting Oranges

Problem:
You are given a grid containing oranges where each cell can have one of the following three values:

    0 - representing an empty cell

    1 - representing a fresh orange

    2 - representing a rotten orange

Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten. Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

Explanation:
This solution uses BFS to simulate the rotting process. It starts by enqueuing all initially rotten oranges and then spreads the rot to adjacent fresh oranges in all four directions. The process continues until no fresh oranges remain or it's determined that some fresh oranges can't be reached.

```go

func orangesRotting(grid [][]int) int {
	queue := []struct{ x, y int }{}
	freshCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, struct{ x, y int }{i, j})
			}
		}
	}
	if freshCount == 0 {
		return 0
	}
	if len(queue) == 0 {
		return -1
	}
	directions := []struct{ x, y int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	minutes := -1
	for len(queue) > 0 {
		for i := len(queue); i > 0; i-- {
			cell := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				x, y := cell.x+dir.x, cell.y+dir.y
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					grid[x][y] = 2
					freshCount--
					if freshCount == 0 {
						return minutes + 1
					}
					queue = append(queue, struct{ x, y int }{x, y})
				}
			}
		}
		minutes++
	}
	return -1
}

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid)) // Output: 4
}
```

## 5. Binary Search

Binary Search is an efficient algorithm for finding an item from a sorted list or array. Instead of checking each element one by one, it repeatedly divides the search range in half, narrowing down the possible locations of the target element until it's found (or determined to be absent).

The array must be sorted for binary search to work correctly.

Start with the whole array `left = 0, right = length - 1`: Find the middle of the array `middle = (left + right) / 2`.

Compare the middle element with the target value:

If it matches, you're done — you’ve found your target.

If the target is smaller, narrow the search to the left half of the array `right to middle - 1`.

If the target is larger, narrow the search to the right half of the array `left to middle + 1`.

Then Again  Find the middle of the right half or left half of and Repeat the process, each time halving the search range, until you either find the target or the search range is empty.

Explanation:
This function performs a standard binary search. It initializes two pointers, left and right, to the start and end of the array, respectively. In each iteration, it calculates the middle index mid and compares the middle element with the target. Depending on the comparison, it adjusts the search range by updating either left or right. The loop continues until the target is found or the search range is exhausted.

```go
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 5)) // Output: 4
	fmt.Println(binarySearch(arr, 11)) // Output: -1
}
```
Efficiency: O(log n) time.



Find First and Last Occurrence

Problem:
Given a sorted array of integers and a target value, find the first and last occurrence of the target value.

Explanation:
To find the first occurrence, the algorithm searches for the target and continues searching in the left half even after finding it. Similarly, to find the last occurrence, it searches for the target and continues searching in the right half after finding it. This ensures that the first and last occurrences are correctly identified.

```go
func findFirstAndLast(arr []int, target int) (int, int) {
	first, last := -1, -1

	// Find first occurrence
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			first = mid
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Find last occurrence
	left, right = 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			last = mid
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return first, last
}

func main() {
	arr := []int{1, 2, 2, 2, 3, 4, 5}
	fmt.Println(findFirstAndLast(arr, 2)) // Output: (1, 3)
}
```

4. Search in Rotated Sorted Array

Problem:
Given a rotated sorted array and a target value, search for the target in the array.

A Rotated Sorted Array is an array that was originally sorted in ascending order but then "rotated" (shifted) at some pivot point.

`[1, 2, 3, 4, 5, 6, 7]` if you "rotate" it at index 3, you move the first part to the end:-> `[4, 5, 6, 7, 1, 2, 3]`

Explanation:
The algorithm first determines which half of the array is sorted. If the left half is sorted, it checks if the target lies within the range of the left half. If so, it narrows the search to the left half; otherwise, it searches the right half. The same logic applies if the right half is sorted. This ensures that the target is found efficiently.

```go
func search(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[left] <= arr[mid] {
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	arr := []int{6, 7, 9, 15, 19, 2, 3}
	fmt.Println(search(arr, 15)) // Output: 3
	fmt.Println(search(arr, 5))  // Output: -1
}
```

Find Minimum in Rotated Sorted Array

Problem:
Given a rotated sorted array, find the minimum element.

Explanation:
The algorithm uses binary search to find the pivot point where the rotation occurs. If the middle element is greater than the rightmost element, the minimum must be in the right half; otherwise, it's in the left half. The loop continues until left equals right, at which point arr[left] is the minimum element.

```go
func findMin(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left]
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(arr)) // Output: 0
}
```

## Dynamic Programming

Memoization is a way of saving the result of a function once you've calculated it — so you don't have to recalculate it again later.

Memoization prevents this by storing the result of each subproblem in a map, array, or matrix, so it’s only computed once.

We want to calculate the n-th Fibonacci number. 0, 1, 1, 2, 3, 5, 8, 13, 21, ...
```go
func fib(n int) int {
	if n <= 1 {
		return n
	}
    // Each number is the sum of the previous two:
	return fib(n-1) + fib(n-2)
}
```
Instead of recalculating each Fibonacci number over and over again, we calculate it once and store it.

```go

func fib(n int, memo map[int]int) int {
    // Step 1: If result is already in memo, return it
	if val, found := memo[n]; found {
		return val
	}
    // Step 2: Base case
	if n <= 1 {
		return n
	}
    // Step 3: Recursive case, but save the result in memo
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
    // Step 4: Return the result
	return memo[n]
}

func main() {
	memo := make(map[int]int)
	fmt.Println(fib(50, memo)) // Output: 12586269025
}
```
You are climbing a staircase with n steps.
Each time you can either climb 1 or 2 steps.
Your task is to compute: how many distinct ways can you climb to the top?

If you are on step n, the last move you made must have been:
- From step n-1 (a 1-step jump)
- From step n-2 (a 2-step jump)

`ways(n) = ways(n-1) + ways(n-2)`

```go
func climbStairs(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = climbStairs(n-1, memo) + climbStairs(n-2, memo)
	return memo[n]
}
```

    ways(5) = ways(4) + ways(3)
        = (ways(3) + ways(2)) + (ways(2) + ways(1))
        = ((ways(2) + ways(1)) + 2) + (2 + 1)
        = (2 + 1 + 2) + (2 + 1)
        = 5 + 3 = 8

1. 0/1 Knapsack Problem

Problem:
Given a set of items, each with a weight and a value, determine the maximum value that can be obtained by selecting a subset of items such that their total weight does not exceed a given capacity.

For each item, you have two options:

    Include it: add its value to the total and subtract its weight from the remaining capacity.

    Exclude it: don’t include the item, just move to the next item.

`knapsack(i, W) = max(knapsack(i-1, W), value[i] + knapsack(i-1, W-weight[i]))`

```go
func knapsackRecursive(wt, val []int, W int, n int) int {
    if n == 0 || W == 0 {
        return 0 // Base case: no items or no capacity left
    }
    if wt[n-1] > W {
        return knapsackRecursive(wt, val, W, n-1) // Can't include item
    }
    include := val[n-1] + knapsackRecursive(wt, val, W-wt[n-1], n-1) // Include item
    exclude := knapsackRecursive(wt, val, W, n-1) // Exclude item
    return max(include, exclude)
}
```
Explanation:
This solution uses a 1D DP array to store the maximum value achievable for each weight capacity. By iterating through each item and updating the DP array in reverse order, we ensure that each item is only considered once.

```go
func knapsack(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([]int, capacity+1)

	for i := 0; i < n; i++ {
		for w := capacity; w >= weights[i]; w-- {
			dp[w] = max(dp[w], dp[w-weights[i]]+values[i])
		}
	}
	return dp[capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weights := []int{1, 2, 3}
	values := []int{10, 20, 30}
	capacity := 4
	fmt.Println(knapsack(weights, values, capacity)) // Output: 50
}
```

2. Longest Common Subsequence (LCS)

Problem:
Given two strings, find the length of their longest common subsequence.

If the last characters of both strings match:

    The LCS of the two strings is 1 + LCS of the rest of both strings.

If they don’t match:

    We have two possibilities: either exclude the last character from the first string or from the second string, and take the maximum length.

```go
func lcsRecursive(s1, s2 string, i, j int) int {
    if i == 0 || j == 0 {
        return 0 // Base case: no common subsequence
    }
    if s1[i-1] == s2[j-1] {
        return 1 + lcsRecursive(s1, s2, i-1, j-1) // Match: include character
    }
    return max(lcsRecursive(s1, s2, i-1, j), lcsRecursive(s1, s2, i, j-1)) // No match: take max of excluding one string
}
```
    For s1 = "abcde", s2 = "ace", the LCS is "ace", and the length is 3.

    Compare characters:

        a from s1 matches a from s2 → 1 match, continue.

        b doesn’t match c, so consider two branches (exclude b or c).

        Continue until you find the longest subsequence.

Explanation:
The LCS problem is solved using a 2D DP table where dp[i][j] represents the length of the LCS of substrings X[0..i-1] and Y[0..j-1]. The table is filled based on character matches and previously computed values.

```go
func lcs(X, Y string) int {
	m, n := len(X), len(Y)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	X := "AGGTAB"
	Y := "GXTXAYB"
	fmt.Println(lcs(X, Y)) // Output: 4
}

```

3. Coin Change Problem

Problem:
Given a set of coin denominations and a total amount, determine the minimum number of coins needed to make that amount.

Recursive Approach:

    Key Insight:
    For each coin, we can either:

        Use it (subtract coin value from the amount and recurse).

        Skip it and check with the remaining coins.

```go
func coinChangeRecursive(coins []int, amount int) int {
    if amount == 0 {
        return 0 // No coins needed for 0 amount
    }
    if amount < 0 {
        return -1 // Invalid case
    }
    min := int(^uint(0) >> 1) // Large value for comparison
    for _, coin := range coins {
        res := coinChangeRecursive(coins, amount-coin)
        if res >= 0 && res < min {
            min = res + 1
        }
    }
    if min == int(^uint(0)>>1) {
        return -1
    }
    return min
}
```
    For coins [1, 2, 5] and target 11:

    Use 5 first: coinChangeRecursive([1, 2, 5], 11-5) = coinChangeRecursive([1, 2, 5], 6)

    Continue until the target is reduced to 0.


Explanation:
This solution uses a 1D DP array where dp[i] represents the minimum number of coins needed to make amount i. The array is updated by considering each coin denomination.

```go
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount)) // Output: 3
}
```

4. Longest Increasing Subsequence (LIS)

Problem:
Given an array of integers, find the length of the longest increasing subsequence.

Explanation:
The LIS problem is solved using a 1D DP array where dp[i] represents the length of the longest increasing subsequence that ends with nums[i]. The array is updated by comparing each element with previous elements.

```go
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLen := 0
	for _, length := range dp {
		maxLen = max(maxLen, length)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums)) // Output: 4
}
```
    Explanation of Code:

    Initialization:

        We initialize a dp array where dp[i] represents the length of the longest increasing subsequence that ends at index i. Initially, we set each element of dp to 1, because every element can at least form a subsequence of length 1.

    Filling the dp array:

        For each element nums[i] (starting from the second element), we check all previous elements nums[j] (where j < i).

        If nums[i] > nums[j], then nums[i] can extend the subsequence ending at nums[j]. So, we update dp[i] as max(dp[i], dp[j] + 1).

    Result:

        The answer is the maximum value in the dp array, which gives us the length of the longest increasing subsequence.

Optimized Approach (Using Binary Search):

To improve the time complexity, we can use a more efficient approach that leverages Binary Search. The idea is to maintain a list that helps us find the correct position for each element in the increasing subsequence.
Key Insight:

    We maintain an array tails[], where tails[i] represents the smallest possible tail value of any increasing subsequence of length i+1.

    For each number in nums, we use binary search to find its position in the tails array. This allows us to efficiently update the subsequence length.

```go
package main

import "fmt"
import "sort"

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    tails := []int{}
    
    for _, num := range nums {
        // Find the position of 'num' in tails using binary search
        pos := binarySearch(tails, num)
        if pos == len(tails) {
            tails = append(tails, num) // num is larger than any element in tails
        } else {
            tails[pos] = num // Replace existing element with num
        }
    }
    
    return len(tails) // The length of tails array is the length of the LIS
}

func binarySearch(tails []int, target int) int {
    left, right := 0, len(tails)-1
    for left <= right {
        mid := left + (right-left)/2
        if tails[mid] == target {
            return mid
        } else if tails[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}

func main() {
    nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
    result := lengthOfLIS(nums)
    fmt.Println("Length of LIS:", result)  // Output: 4
}
```

## Linked List

A Linked List is a linear data structure where elements (called nodes) are stored in a sequence, and each node contains two parts:
- Data: The value of the node.
- Next: A reference (or pointer) to the next node in the list.

InsertAtNthPosition Method:

    The method allows you to insert a new node at the n-th position in the list.

    Position 0: If n is 0, the new node is inserted at the head (beginning of the list).

    Traversing the List: If n > 0, the method traverses the list until it reaches the (n-1)-th position and inserts the new node at the n-th position.

    Out of Bounds: If n exceeds the list’s length, an error is returned.

    Insertion Logic: The new node's next is set to point to the old n-th node, and the (n-1)-th node’s next is updated to point to the new node.


```go
package main

import "fmt"

// Define a Node structure
type Node struct {
    data int
    next *Node
}

// Define the LinkedList structure
type LinkedList struct {
    head *Node
}

// Method to insert a node at the N-th position
func (list *LinkedList) InsertAtNthPosition(n int, data int) error {
    newNode := &Node{data: data}

    // Case 1: Insert at the beginning (n = 0)
    if n == 0 {
        newNode.next = list.head
        list.head = newNode
        return nil
    }

    // Traverse to the (n-1)-th position
    current := list.head
    position := 0
    for current != nil && position < n-1 {
        current = current.next
        position++
    }

    // Case 2: If the position is greater than the length of the list
    if current == nil {
        return fmt.Errorf("position %d is out of bounds", n)
    }

    // Case 3: Insert the new node at the n-th position
    newNode.next = current.next
    current.next = newNode

    return nil
}

// Method to insert a node at the beginning
func (list *LinkedList) InsertAtBeginning(data int) {
    newNode := &Node{data: data}
    newNode.next = list.head
    list.head = newNode
}

// Method to insert a node at the end
func (list *LinkedList) InsertAtEnd(data int) {
    newNode := &Node{data: data}
    if list.head == nil {
        list.head = newNode
        return
    }
    lastNode := list.head
    for lastNode.next != nil {
        lastNode = lastNode.next
    }
    lastNode.next = newNode
}

// Method to delete a node with a specific value
func (list *LinkedList) DeleteNode(data int) {
    if list.head == nil {
        return
    }

    if list.head.data == data {
        list.head = list.head.next
        return
    }

    current := list.head
    for current.next != nil && current.next.data != data {
        current = current.next
    }

    if current.next == nil {
        fmt.Println("Node with value", data, "not found.")
        return
    }

    current.next = current.next.next
}

// Method to search for a node with a specific value
func (list *LinkedList) Search(data int) bool {
    current := list.head
    for current != nil {
        if current.data == data {
            return true
        }
        current = current.next
    }
    return false
}

// Method to display the entire linked list
func (list *LinkedList) Display() {
    if list.head == nil {
        fmt.Println("List is empty.")
        return
    }
    current := list.head
    for current != nil {
        fmt.Print(current.data, " ")
        current = current.next
    }
    fmt.Println()
}
```

## Binary Tree
A binary tree is a type of data structure in which each node has at most two children, referred to as the left child and the right child. This makes the tree structure easy to navigate and manage.

Root: The topmost node of the tree. It’s the starting point for any traversal.

Node: An individual element in the tree that holds data.

Left Child: The child node that appears on the left side of the parent node.

Right Child: The child node that appears on the right side of the parent node.

Leaf: A node that has no children (both left and right children are empty).

Height: The length of the longest path from the node to a leaf node.

Depth: The length of the path from the root to a particular node.


Understanding the Traversal Types:

In-order traversal: Visit the left subtree, the root node, and then the right subtree. In-order traversal is commonly used in binary search trees (BSTs) because it visits nodes in sorted order.

        Example Order: Left → Root → Right

Pre-order traversal: Visit the root node first, then the left subtree, followed by the right subtree.

        Example Order: Root → Left → Right

Post-order traversal: Visit the left subtree first, then the right subtree, and finally the root node.

        Example Order: Left → Right → Root

```go
// TreeNode represents a node in the binary tree.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

```go
// In-order traversal (Left → Root → Right)
func inOrderTraversal(root *TreeNode) {
    if root == nil {
        return
    }
    // Traverse left subtree
    inOrderTraversal(root.Left)
    // Visit root
    fmt.Print(root.Val, " ")
    // Traverse right subtree
    inOrderTraversal(root.Right)
}
```
```go
// Pre-order traversal (Root → Left → Right)
func preOrderTraversal(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Print(root.Val, " ")
    preOrderTraversal(root.Left)
    preOrderTraversal(root.Right)
}
```

```go
// Post-order traversal (Left → Right → Root)
func postOrderTraversal(root *TreeNode) {
    if root == nil {
        return
    }
    postOrderTraversal(root.Left)
    postOrderTraversal(root.Right)
    fmt.Print(root.Val, " ")
}
```

```go
func main() {
    // Construct a sample binary tree:
    //        1
    //       / \
    //      2   3
    //     / \ / \
    //    4  5 6  7

    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 4}
    root.Left.Right = &TreeNode{Val: 5}
    root.Right.Left = &TreeNode{Val: 6}
    root.Right.Right = &TreeNode{Val: 7}

    // In-order traversal: 4 2 5 1 6 3 7
    fmt.Println("In-order Traversal:")
    inOrderTraversal(root)
    fmt.Println()

    // Pre-order traversal: 1 2 4 5 3 6 7
    fmt.Println("Pre-order Traversal:")
    preOrderTraversal(root)
    fmt.Println()

    // Post-order traversal: 4 5 2 6 7 3 1
    fmt.Println("Post-order Traversal:")
    postOrderTraversal(root)
    fmt.Println()
}
```
```bash
In-order Traversal:
4 2 5 1 6 3 7 
Pre-order Traversal:
1 2 4 5 3 6 7 
Post-order Traversal:
4 5 2 6 7 3 1
```

## Binary Search Tree (BST):
        10
       /  \
      5   20
     / \   / \
    2   7 15  25

A binary search tree is a special type of binary tree in which the nodes follow these properties:

- The left subtree of a node contains only nodes with values less than the node’s value.

- The right subtree of a node contains only nodes with values greater than the node’s value.

This makes searching for a value in a BST very efficient, as you can ignore half of the tree at each step, resulting in a time complexity of O(log n) for search operations in a balanced BST.

Binary Search Trees (BSTs) are used in databases and file systems for quick lookups, insertions, and deletions.

Priority Queues can be implemented using heaps, a type of binary tree.



## Heap / Priority Queue
A heap is a special type of binary tree used to efficiently manage and retrieve the maximum or minimum value in a collection. 

- Max Heap: In a max heap, the parent node is always greater than or equal to its child nodes. The largest element is always at the root.
- Min Heap: In a min heap, the parent node is always smaller than or equal to its child nodes. The smallest element is always at the root.

A heap is often represented as an array, where for any element at index i:

- The left child is at index 2*i + 1
- The right child is at index 2*i + 2
- The parent is at index (i-1) // 2

This structure ensures that the heap can be efficiently implemented and that operations like insertion and removal can be done in O(log n) time.

Heap Sort: A sorting algorithm that uses the heap data structure.

Heap Operations:

- Insertion:

When you insert an element into a heap, you add it at the end of the tree (or the last index in the array) and heapify it (adjust the heap to maintain the heap property).

- Deletion:

The element with the highest priority (root of the heap) is removed. After removal, the heap needs to be restructured to maintain the heap property.

- Heapify:

This is the process of adjusting the heap to restore the heap property. This can happen after insertion or deletion, and it involves comparing the parent with its children and swapping them if necessary.

A priority queue is a data structure that allows elements to be inserted and removed based on their priority rather than the order they were added. It is commonly implemented using a heap.

- In a max-priority queue, the element with the highest priority is always removed first.

- In a min-priority queue, the element with the lowest priority is removed first.

In both cases, the heap ensures that you can access the highest or lowest priority element efficiently.


```go
package main

import (
	"fmt"
)
type Task struct {
	name     string
	priority int
}

type PriorityQueue struct {
	data []Task
}

type MinHeap struct {
	data []int
}

// To insert a task into the priority queue
func (pq *PriorityQueue) Enqueue(task Task) {
	pq.data = append(pq.data, task)
	pq.upHeap(len(pq.data) - 1) // "Heapify up" to restore max-heap property
}

// Insert a new element into the heap
func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.upHeap(len(h.data) - 1)
}

// To remove the task with highest priority
func (pq *PriorityQueue) Dequeue() Task {
	if len(pq.data) == 0 {
		panic("priority queue is empty")
	}
	top := pq.data[0]
	last := pq.data[len(pq.data)-1]
	pq.data = pq.data[:len(pq.data)-1]
	if len(pq.data) > 0 {
		pq.data[0] = last
		pq.downHeap(0) // "Heapify down" to fix the order
	}
	return top
}

// Remove and return the smallest element
func (h *MinHeap) Pop() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}
	root := h.data[0]
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 0 {
		h.data[0] = last
		h.downHeap(0)
	}
	return root
}

// Move the element at index up to restore heap property
func (h *MinHeap) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[i] >= h.data[parent] {
			break
		}
		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}
func (pq *PriorityQueue) upHeap(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.data[i].priority <= pq.data[parent].priority {
			break
		}
		pq.data[i], pq.data[parent] = pq.data[parent], pq.data[i]
		i = parent
	}
}

// Move the element at index down to restore heap property
func (h *MinHeap) downHeap(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func (pq *PriorityQueue) downHeap(i int) {
	n := len(pq.data)
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && pq.data[left].priority > pq.data[largest].priority {
			largest = left
		}
		if right < n && pq.data[right].priority > pq.data[largest].priority {
			largest = right
		}
		if largest == i {
			break
		}
		pq.data[i], pq.data[largest] = pq.data[largest], pq.data[i]
		i = largest
	}
}

func main() {
	h := &MinHeap{}

	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)
	h.Insert(6)

	fmt.Println("Min:", h.Pop()) // Should print 1
	fmt.Println("Min:", h.Pop()) // Should print 3
	fmt.Println("Min:", h.Pop()) // Should print 5
}
```

```go
func main() {
	pq := &PriorityQueue{}

	pq.Enqueue(Task{name: "Task A", priority: 2})
	pq.Enqueue(Task{name: "Task B", priority: 5})
	pq.Enqueue(Task{name: "Task C", priority: 1})
	pq.Enqueue(Task{name: "Task D", priority: 4})

	for len(pq.data) > 0 {
		task := pq.Dequeue()
		fmt.Printf("Processing: %s (priority %d)\n", task.name, task.priority)
	}
}
```
To make a Max-Heap, just change the comparison signs:

`if h.data[i] <= h.data[parent] // for upHeap`

AND

`if h.data[left] > h.data[largest] // for downHeap`

The logic is otherwise the same — just reversed comparisons.


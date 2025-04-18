
# A : TWO POINTERS

Imagine We have a row of boxes (an array), and we want to look at two boxes at a time. Instead of using one finger to check  use two fingers (pointers) to move across the boxes from different directions. And yoweu move them closer or far to each other based on some rule. Reduces nested loops to single traversal.

Ask ourself:

- ❓ `Is the data sorted?` If the data is sorted, you can make smart decisions about which direction to move the pointers (left or right). 

- ❓ `Am I comparing elements?` If you're comparing elements (like looking for a pair of numbers, or a matching pattern), the Two Pointer approach helps avoid checking each element against every other element, making your solution faster.

- ❓ `Do I need to optimize time over checking all pairs?` Instead of checking every single pair, you’re reducing the number of checks by moving the pointers intelligently, like narrowing down possibilities.

- ❓ `Is this about ranges, windows, or reversals?` "Ranges" or "windows" refer to parts of the array or list that you're focusing on. Then Two Pointers are perfect because they let you expand or shrink the window efficiently.

If yes, 💥 boom — you're probably looking at a Two Pointer problem.

## 1: Reverse a slice using two pointers

- We start with two pointers: left at the beginning and right at the end of the array.
- The loop continues as long as left is less than right.
- Inside the loop, we swap the values at arr[left] and arr[right].
- Then we move left forward and right backward to keep swapping inward.
- When they meet (or cross), the array is completely reversed in place.

```go
left, right := 0, len(arr)-1
for left < right {
	arr[left], arr[right] = arr[right], arr[left]
	left++
	right--
}
```
## 2: Move 0s to front

- It uses two pointers: left marks the position for placing non-zero elements, and right iterates through the array.
- If nums[right] is 0, it swaps nums[left] and nums[right] and then increments left. This ensures that every non-zero element gets placed at the front and zeros are shifted to the end.
- The loop runs in one pass, making the solution efficient with O(n) time complexity and O(1) space complexity.

```go
left := 0
for right := 0; right < len(nums); right++ {
    if nums[right] == 0 {
        nums[left], nums[right] = nums[right], nums[left]
        left++
    }
}
```
## 3: Find if a pair with a given sum exists in a sorted array

- The array arr is assumed to be sorted in increasing order.
- We start with two pointers: left at the start, and right at the end of the array.
- We check the sum of the two numbers: arr[left] + arr[right].
- If the sum is less than the target, we move left forward to get a bigger number; if it's more, we move right back to get a smaller number.
- If the sum matches the target, we return true — we found the pair!


```go
left, right := 0, len(arr)-1
for left < right {
	sum := arr[left] + arr[right]
	if sum == target {
		return true
	} else if sum < target {
		left++
	} else {
		right--
	}
}
```
## 4:  Remove duplicates from sorted array

- The array nums is sorted, so all duplicates are next to each other.
- We use two pointers: fast goes through the array, and slow keeps track of the last unique element found.
- Whenever nums[fast] != nums[slow], it means we found a new unique number.
- We move slow forward and copy that unique number to the slow position.
- At the end, nums[:slow+1] contains all the unique elements in order, and the rest can be ignored.

```go
slow := 0
for fast := 1; fast < len(nums); fast++ {
	if nums[fast] != nums[slow] {
		slow++
		nums[slow] = nums[fast]
	}
}
return nums[:slow+1]
```

## 5: Longest Substring Without Repeating Characters

- `seen := make(map[byte]int)` — A map (seen) keeps track of each character's position in the string s. The key is the character, and the value is its most recent index.
- `left, maxLen := 0, 0` — Two variables are initialized: left (which marks the start of the window) and maxLen (which will store the maximum length of substrings without duplicates).
- We loop through the string with right as the pointer: `for right := 0; right < len(s); right++` — As we go through each character, we check if it's already in the seen map and whether its stored position is within the current window (left to right).
- If the character is found inside the current window, we move left to the right of the previous occurrence to ensure no duplicates exist in the window. `left = seen[ch] + 1`.
- For the character ch, update its position in the map to the current index right.
- Every time we do, we compare it to our current longest (maxLen), and if the new one is longer, we update it.
- We update the maxLen by comparing the current window's length (right - left + 1) with the longest substring found so far. We add +1 because both left and right are inclusive, so the window length must count both ends.

```go
seen := make(map[byte]int)
left, maxLen := 0, 0
for right := 0; right < len(s); right++ {
    ch := s[right]
    if _, ok := seen[ch]; ok && seen[ch] >= left {
        left = seen[ch] + 1 // shrink from left
    }
    seen[ch] = right
    maxLen = max(maxLen, right-left+1)
}
```







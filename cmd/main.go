package main

import "fmt"

// Function to find the maximum of two values
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function to find the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int) // Map to store the last index of each character
	left, maxLen := 0, 0       // Initialize left pointer and maxLen to 0

	// Iterate through the string with the right pointer
	for right := range len(s) {
		ch := s[right] // Current character at the right pointer

		// If the character is already seen and its index is within the current window
		if _, ok := seen[ch]; ok && seen[ch] >= left {
			left = seen[ch] + 1 // Move the left pointer right past the duplicate
		}

		// Update the last index of the character
		seen[ch] = right

		// Calculate the length of the current valid substring
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen // Return the length of the longest substring
}

func main() {
	s := "abcabcbb"                       // Example string
	result := lengthOfLongestSubstring(s) // Call the function
	fmt.Println("Length of longest substring without repeating characters:", result)
}

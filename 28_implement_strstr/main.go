package main

import "fmt"

/*
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().



Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Example 3:

Input: haystack = "", needle = ""
Output: 0
*/

func strStr(haystack string, needle string) int {
	// Best Case:  θ(4) Cost: 1 + 1 + 1 + 1 = 4
	// Worst Case: θ(n * m) Cost: 2n + 1 + 1 + 1 =  2n + 3

	n := len(haystack) // ------> 1
	m := len(needle)   // ------> 1

	if m == 0 { // ------> 1
		return 0 // ------> 1
	}

	var j int = 0 // ------> 1

	for i := 0; i < n; i++ { // ------> n, m

		if haystack[i] == needle[j] { // ------> 1
			j++ // ------> 1
		} else {
			i = i - j // ------> 1
			j = 0     // ------> 1
		}

		if j == m { // ------> 1
			return i - j + 1 // ------> 1
		}
	}

	return -1 // ------> 1s
}

func main() {
	fmt.Println(strStr("hello", "llr"))
	//fmt.Println(strStr("mississippi", "issipi"))
}

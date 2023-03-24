package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

/*
Given a string s consisting of some words separated by some number of spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.


Constraints:

1 <= s.length <= 104
s consists of only English letters and spaces ' '.
There will be at least one word in s.
*/

func lengthOfLastWord(s string) int {
	// Best Case:  θ(n) Cost: 1 + n + 1 + 1 + 1 + 1 + 1 = n + 6
	// Worst Case: θ(n) Cost: 1 + 2n + 2 + 1 =  2n + 4

	var count int = 0 // ----> 1

	for i := len(s) - 1; i >= 0; i-- { // ----> n
		if s[i] != ' ' { // ----> 1
			count++ // ----> 1
		} else {
			if count > 0 { // ----> 1
				break // ----> 1
			}
		}
	}

	return count // ----> 1
}

func main() {
	inputs := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
	}

	for _, v := range inputs {
		fmt.Println(lengthOfLastWord(v))
	}

	fmt.Println(uuid.NewV4().String())

}

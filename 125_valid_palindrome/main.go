package main

import (
	"fmt"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/

func isPalindrome(s string) bool {

	var ok bool
	var comp rune
	endPointer := len(s) - 1

	for k := 0; k <= len(s); {

		v := rune(s[k])

		if endPointer <= k {
			return true
		}

		if v, ok = convert(v); !ok {
			k++
			continue
		}

		if comp, ok = convert(rune(s[endPointer])); !ok {
			endPointer--
			continue
		}

		if v != comp {
			return false
		}

		endPointer--
		k++
	}

	return true
}

func convert(v rune) (rune, bool) {

	if 65 <= v && v <= 90 {
		return v + 32, true
	}

	if (v < 97 || 122 < v) && !(48 <= v && v <= 57) {
		return v, false
	}

	return v, true

}

func main() {
	inputs := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
	}

	for _, v := range inputs {
		fmt.Println(isPalindrome(v))
	}
}

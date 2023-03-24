package main

import "fmt"

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	m := make(map[string]string)

	m["("] = ")"
	m["["] = "]"
	m["{"] = "}"

	var stack []string

	for _, c := range s {
		if value, found := m[string(c)]; found {
			// push closing char
			stack = append(stack, value)
		} else {
			n := len(stack)

			// top element = stack[n-1]

			if n > 0 && stack[n-1] == string(c) {
				// pop element
				stack = stack[:n-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	inputs := []string{
		"()",
		"()[]{}",
		"(]",
	}

	for _, v := range inputs {
		fmt.Println(isValid(v))
	}

}

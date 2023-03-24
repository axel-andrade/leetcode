package main

import (
	"fmt"
	"sort"
)

/* Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

func containsDuplicateWithMap(nums []int) bool {
	m := make(map[int]int)

	for i, num := range nums {
		if _, found := m[num]; found {
			fmt.Println("encontrou")
			return true
		}

		m[num] = i
		fmt.Println(m)
	}

	return false
}

func containsDuplicateWithSort(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func main() {
	inputs := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}

	for _, v := range inputs {
		fmt.Println(containsDuplicateWithSort(v))
	}

}

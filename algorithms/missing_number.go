/**
 * Missing Number
 *
 * Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
 * find the one that is missing from the array.
 *
 * For example,
 * Given nums = [0, 1, 3] return 2.
 * Note:
 * Your algorithm should run in linear runtime complexity.
 * Could you implement it using only constant extra space complexity?
 *
 * Credits:
 * Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
 */

package main

import "fmt"

func main() {
	nums := []int{0, 4, 1, 3}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	ret := 0
	for i := 0; i < len; i++ {
		ret ^= i + 1
		ret ^= nums[i]
	}
	return ret
}

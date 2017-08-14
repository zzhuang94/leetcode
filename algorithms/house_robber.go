/**
 * House Robber
 *
 * You are a professional robber planning to rob houses along a street.
 * Each house has a certain amount of money stashed, the only constraint
 * stopping you from robbing each of them is that adjacent houses have
 * security system connected and it will automatically contact the police
 * if two adjacent houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of each house,
 * determine the maximum amount of money you can rob tonight without alerting the police.
 *
 * Credits:
 * Special thanks to @ifanchu for adding this problem and creating all test cases.
 * Also thanks to @ts for adding additional test cases.
 */

package main

import "fmt"

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	if len == 1 {
		return nums[0]
	}
	a, b := nums[0], nums[0]
	if nums[1] > nums[0] {
		b = nums[1]
	}
	for i := 2; i < len; i++ {
		r1 := a + nums[i]
		r2 := b
		a = b
		if r1 >= r2 {
			b = r1
		} else {
			b = r2
		}
	}
	return b
}

/**
 * Given an array of integers, return indices of the two numbers
 * such that they add up to a specific target.
 * You may assume that each input would have exactly one solution,
 * and you may not use the same element twice.
 *
 * Example:
 * Given nums = [2, 7, 11, 15], target = 9,
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 */

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15, 3}
	target := 14
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	len := len(nums)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

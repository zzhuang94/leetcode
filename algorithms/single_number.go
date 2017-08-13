/**
 * Single Number
 *
 * Given an array of integers, every element appears twice except for one. Find that single one.
 *
 * Note:
 * Your algorithm should have a linear runtime complexity.
 * Could you implement it without using extra memory?
 */

package main

import "fmt"

func main() {
	nums := []int{600, 2, 3, 4, 5, 4, 3, 2, 5, 600}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	for i := 0; i < len-1; i++ {
		nums[i+1] ^= nums[i]
	}
	return nums[len-1]
}

/**
 * Maximum Subarray
 *
 * Find the contiguous subarray within an array
 * (containing at least one number) which has the largest sum.
 *
 * For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
 * the contiguous subarray [4,-1,2,1] has the largest sum = 6.
 */

package main

import "fmt"

func main() {
	nums := []int{-100, -20, -30, 10, -5}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	t := 0
	max := nums[0]
	for i := 0; i < len; i++ {
		t += nums[i]
		if t > max {
			max = t
		}
		if t < 0 {
			t = 0
		}
	}
	return max
}

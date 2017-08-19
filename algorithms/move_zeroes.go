/**
 * Move Zeroes
 *
 * Given an array nums, write a function to move all 0's to the end of it
 * while maintaining the relative order of the non-zero elements.
 *
 * For example, given nums = [0, 1, 0, 3, 12], after calling your function,
 * nums should be [1, 3, 12, 0, 0].
 *
 * Note:
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 * Credits:
 * Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
 */

package main

func main() {
	nums := []int{}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	l := len(nums)
	step := 0
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			step++
		} else if step > 0 {
			nums[i-step] = nums[i]
		}
	}
	for i := 0; i < step; i++ {
		nums[l-1-i] = 0
	}
}

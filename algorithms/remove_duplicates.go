/**
 * Remove Duplicates from Sorted Array
 *
 * Given a sorted array, remove the duplicates in place
 * such that each element appear only once and return the new length.
 *
 * Do not allocate extra space for another array,
 * you must do this in place with constant memory.
 *
 * For example,
 * Given input array nums = [1,1,2],
 *
 * Your function should return length = 2, with the first two elements
 * of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.
 */

package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	len := len(nums)
	n_l := len
	for i := 0; i < len-1; i++ {
		if nums[i] == nums[i+1] {
			n_l--
		}
	}
	return n_l
}

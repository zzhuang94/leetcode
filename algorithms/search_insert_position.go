/**
 * Search Insert Position
 *
 * Given a sorted array and a target value, return the index if the target is found.
 * If not, return the index where it would be if it were inserted in order.
 *
 * You may assume no duplicates in the array.
 *
 * Here are few examples.
 * [1,3,5,6], 5 → 2
 * [1,3,5,6], 2 → 1
 * [1,3,5,6], 7 → 4
 * [1,3,5,6], 0 → 0
 */

package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 2))
}

func searchInsert(nums []int, target int) int {
	len := len(nums)
	for i := 0; i < len; i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len
}

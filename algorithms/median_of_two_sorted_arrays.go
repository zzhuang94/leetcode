/**
 * 4. Median of Two Sorted Arrays
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
 * You may assume nums1 and nums2 cannot be both empty.
 * Example 1:
 * nums1 = [1, 3]
 * nums2 = [2]
 * The median is 2.0
 * Example 2:
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * The median is (2 + 3)/2 = 2.5package main
 */

package main

import "fmt"

func main() {
	a := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(a)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sortList := []int{}
	for i, j := 0, 0; ; {
		if i >= len(nums1) && j >= len(nums2) {
			break
		} else if i >= len(nums1) || j < len(nums2) && nums1[i] >= nums2[j] {
			sortList = append(sortList, nums2[j])
			j++
		} else {
			sortList = append(sortList, nums1[i])
			i++
		}
	}
	l := len(sortList)
	return float64(sortList[(l-1)/2]+sortList[l/2]) / 2.0
}

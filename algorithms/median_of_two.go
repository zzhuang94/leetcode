/**
 * Median of Two Sorted Arrays
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
 *
 * Example 1:
 * nums1 = [1, 3]
 * nums2 = [2]
 * The median is 2.0
 *
 * Example 2:
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * The median is (2 + 3)/2 = 2.5
 */
package main

import "fmt"

func main() {
	nums1 := []int{1, 4}
	nums2 := []int{}
	fmt.Println(float64(3 / 2))
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fmt.Println("###", nums1, nums2)
	l1 := len(nums1)
	l2 := len(nums2)

	if l1 == 0 {
		if l2 == 0 {
			return 0
		} else {
			t := nums2[l2/2] + nums2[(l2-1)/2]
			return float64(t / 2.0)
		}
	} else {
		if l2 == 0 {
			t := nums1[l1/2] + nums1[(l1-1)/2]
			return float64(t / 2.0)
		}
	}
	if l1 == 1 && l2 == 1 {
		fmt.Println(nums1[0], nums2[0])
		t := (nums1[0] + nums2[0])
		fmt.Println("==", t)
		return float64(t) / 2.0
	}

	if nums1[l1-1] > nums2[l2-1] {
		return findMedianSortedArrays(nums1[0:(l1-1)/2+1], nums2[(l2-1)/2+1:l2])
	} else {
		return findMedianSortedArrays(nums2[0:(l2-1)/2+1], nums1[(l1-1)/2+1:l1])
	}
	return 0
}

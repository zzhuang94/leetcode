/**
 * Intersection of Two Arrays
 *
 * Given two arrays, write a function to compute their intersection.
 *
 * Example:
 * Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].
 *
 * Note:
 * Each element in the result must be unique.
 * The result can be in any order.
 */

package main

import "fmt"

func main() {
	nums1 := []int{}
	nums2 := []int{}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	s := []int{}
	for _, n := range nums1 {
		m[n] = 1
	}
	for _, n := range nums2 {
		_, ok := m[n]
		if ok {
			m[n] = 2
		}
	}
	for k, v := range m {
		if v == 2 {
			s = append(s, k)
		}
	}
	return s
}

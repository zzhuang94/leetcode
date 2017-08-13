/**
 * Pascal's Triangle II
 *
 * Given an index k, return the kth row of the Pascal's triangle.
 *
 * For example, given k = 3,
 * Return [1,3,3,1].
 *
 * Note:
 * Could you optimize your algorithm to use only O(k) extra space?
 */

package main

import "fmt"

func main() {
	fmt.Println(getRow(1))
}

func getRow(rowIndex int) []int {
	t := []int{1}
	if rowIndex <= 1 {
		return t
	}
	for n := 2; n <= rowIndex; n++ {
		for i := n - 2; i > 0; i-- {
			t[i] = t[i] + t[i-1]
		}
		t = append(t, 1)
	}
	return t
}

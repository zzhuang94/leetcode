/**
 * Pascal's Triangle
 *
 * Given numRows, generate the first numRows of Pascal's triangle.
 *
 * For example, given numRows = 5,
 * Return
 *
 * [
 *      [1],
 *     [1,1],
 *    [1,2,1],
 *   [1,3,3,1],
 *  [1,4,6,4,1]
 * ]
 */

package main

import "fmt"

func main() {
	ret := generate(10)
	for _, r := range ret {
		fmt.Println(r)
	}
}

func generate(numRows int) [][]int {
	ret := [][]int{}
	if numRows <= 0 {
		return ret
	}
	ret = append(ret, []int{1})
	for n := 1; n < numRows; n++ {
		t := []int{1}
		r := ret[n-1]
		for i := 0; i < n-1; i++ {
			t = append(t, r[i]+r[i+1])
		}
		t = append(t, 1)
		ret = append(ret, t)
	}
	return ret
}

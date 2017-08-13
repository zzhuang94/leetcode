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
	t := []int{1}
	ret = append(ret, t)
	if numRows <= 1 {
		return ret
	}
	for n := 2; n <= numRows; n++ {
		for i := n - 2; i > 0; i-- {
			t[i] = t[i] + t[i-1]
		}
		t = append(t, 1)
		ret = append(ret, t)
	}
	return ret
}

/**
 * Excel Sheet Column Number
 *
 * Related to question Excel Sheet Column Title
 * Given a column title as appear in an Excel sheet, return its corresponding column number.
 *
 * For example:
 *
 * A -> 1
 * B -> 2
 * C -> 3
 * ...
 * Z -> 26
 * AA -> 27
 * AB -> 28
 * Credits:
 * Special thanks to @ts for adding this problem and creating all test cases.
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 1000; i++ {
		s := convertToTitle(i)
		ii := titleToNumber(s)
		if i != ii {
			fmt.Println("-----", i, s, ii)
		}
		fmt.Println(i, s, ii)
	}
}

func titleToNumber(s string) int {
	len := len(s)
	if len == 0 {
		return 0
	}
	sum := 0
	for i, c := range s {
		t := float64(int(c)-64) * math.Pow(26, float64(len-1-i))
		sum += int(t)
	}
	return sum
}

func convertToTitle(n int) string {
	if n < 1 {
		return ""
	}
	if n < 27 {
		return string(64 + n)
	}
	if n%26 == 0 {
		return convertToTitle(n/26-1) + "Z"
	}
	return convertToTitle(n/26) + convertToTitle(n%26)
}

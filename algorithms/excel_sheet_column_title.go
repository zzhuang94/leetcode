/**
 * Excel Sheet Column Title
 *
 * Given a positive integer, return its corresponding column title as appear in an Excel sheet.
 *
 * For example:
 *
 * 1 -> A
 * 2 -> B
 * 3 -> C
 * ...
 * 26 -> Z
 * 27 -> AA
 * 28 -> AB
 * Credits:
 * Special thanks to @ifanchu for adding this problem and creating all test cases.
 */

package main

import "fmt"

func main() {
	for i := 600; i <= 800; i++ {
		fmt.Println(i, convertToTitle(i), " ")
	}
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

/**
 * Power of Two
 *
 * Given an integer, write a function to determine if it is a power of two.
 *
 * Credits:
 * Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
 */

package main

import "fmt"

func main() {
	for i := -20; i < 20; i++ {
		fmt.Println(i, isPowerOfTwo(i))
	}
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

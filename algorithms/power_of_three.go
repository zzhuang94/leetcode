/**
 * Power of Three
 *
 * Given an integer, write a function to determine if it is a power of three.
 *
 * Follow up:
 * Could you do it without using any loop / recursion?
 *
 * Credits:
 * Special thanks to @dietpepsi for adding this problem and creating all test cases.
 *
 * // 1162261467 is 3^19,  3^20 is bigger than int
 * return ( n>0 &&  1162261467%n==0 );
 */

package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, isPowerOfThree(i))
	}
}

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for {
		if n == 1 {
			return true
		} else if n%3 == 0 {
			n = n / 3
		} else {
			return false
		}
	}
}

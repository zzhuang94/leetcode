/**
 * Power of Four
 *
 * Given an integer (signed 32 bits), write a function to check whether it is a power of 4.
 *
 * Example:
 * Given num = 16, return true. Given num = 5, return false.
 *
 * Follow up: Could you solve it without loops/recursion?
 *
 * Credits:
 * Special thanks to @yukuairoy for adding this problem and creating all tests
 */

package main

import "fmt"

func main() {
	for i := 0; i < 1000000; i++ {
		if isPowerOfFour(i) {
			fmt.Println(i)
		}
	}
}

func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && (num-1)%3 == 0
}

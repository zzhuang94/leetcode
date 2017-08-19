/*
 * Add Digits
 *
 * Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
 *
 * For example:
 * Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.
 * Follow up:
 * Could you do it without any loop/recursion in O(1) runtime?
 * Credits:
 * Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
 */

package main

import "fmt"

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(i, addDigits(i))
	}
}

func addDigits(num int) int {
	// return 1 + (num-1)%9
	for {
		if num < 10 {
			return num
		}
		num = num/10 + num%10
	}
}

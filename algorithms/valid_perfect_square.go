/**
 * Valid Perfect Square
 *
 * Given a positive integer num, write a function which returns True if num is a perfect square else False.
 * Note: Do not use any built-in library function such as sqrt.
 *
 * Example 1:
 * Input: 16
 * Returns: True
 * Example 2:
 * Input: 14
 * Returns: False
 * Credits:
 * Special thanks to @elmirap for adding this problem and creating all test cases.
 */

package main

import "fmt"

func main() {
	for i := 1; i < 200; i++ {
		if isPerfectSquare(i) {
			fmt.Println(i)
		}
	}
}

func isPerfectSquare(num int) bool {
	l, h := 1, num
	for {
		m := (l + h) / 2
		t := m * m
		if t == num {
			return true
		} else if t > num {
			h = m
		} else {
			l = m
		}
		if h-l <= 1 {
			return false
		}
	}
}

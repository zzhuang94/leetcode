/**
 * Happy Number
 *
 * Write an algorithm to determine if a number is "happy".
 *
 * A happy number is a number defined by the following process:
 * Starting with any positive integer, replace the number by the sum of the
 * squares of its digits, and repeat the process until the number equals
 * 1 (where it will stay), or it loops endlessly in a cycle which does not
 * include 1. Those numbers for which this process ends in 1 are happy numbers.
 *
 * Example: 19 is a happy number
 * 12 + 92 = 82
 * 82 + 22 = 68
 * 62 + 82 = 100
 * 12 + 02 + 02 = 1
 *
 * Credits:
 * Special thanks to @mithmatt and @ts for adding this problem and creating all test cases.
 */

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, isHappy(i))
	}
}

func isHappy(n int) bool {
	m := map[int]int{n: n}
	for {
		len := len(strconv.Itoa(n))
		t_sum := 0
		for i := len - 1; i >= 0; i-- {
			r := int(math.Pow(10, float64(i)))
			t := n / r
			t_sum += t * t
			n = n % r
		}
		n = t_sum
		if n == 1 {
			return true
		}
		_, ok := m[n]
		if ok {
			return false
		}
		m[n] = n
	}
}

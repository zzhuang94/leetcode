/**
 * Sum of Two Integers
 *
 * Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.
 *
 * Example:
 * Given a = 1 and b = 2, return 3.
 *
 * Credits:
 * Special thanks to @fujiaozhu for adding this problem and creating all test cases.
 */

package main

import "fmt"

func main() {
	for i := -10; i < 10; i++ {
		fmt.Println(i, getSum(i, i+1))
	}
}

func getSum(a int, b int) int {
	sum := a
	for {
		if b == 0 {
			return sum
		}
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
}

/**
 * Climbing Stairs
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 * Each time you can either climb 1 or 2 steps.
 * In how many distinct ways can you climb to the top?
 *
 * Note: Given n will be a positive integer.
 */

package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		fmt.Println(i, climbStairs(i))
	}
}

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	a, b := 1, 1
	for i := 1; i <= n; i++ {
		a, b = b, a+b
	}
	return a
}

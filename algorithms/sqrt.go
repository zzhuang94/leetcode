/**
 * Sqrt(x)
 *
 * Implement int sqrt(int x).
 * Compute and return the square root of x.
 */

package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	g, l := x, 1
	for {
		fmt.Println(l, g)
		if g-l <= 1 {
			return l
		}
		mid := (g + l) >> 1
		if mid*mid < x {
			l = mid
		} else if mid*mid == x {
			return mid
		} else {
			g = mid
		}
	}
}

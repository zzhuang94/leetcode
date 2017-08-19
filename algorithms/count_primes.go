/*
 * Count Primes
 *
 * Description:
 * Count the number of prime numbers less than a non-negative number, n.
 *
 * Credits:
 * Special thanks to @mithmatt for adding this problem and creating all test cases.
 */

package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, countPrimes(i))
	}
	fmt.Println(499979, countPrimes(499979))
}

func countPrimes(n int) int {
	a := make([]int, n)
	for i := 2; i*i < n; i++ {
		if a[i] == 0 {
			for j := i; i*j < n; j++ {
				a[i*j] = 1
			}
		}
	}
	t := 0
	for i := 2; i < n; i++ {
		if a[i] == 0 {
			t++
		}
	}
	return t
}

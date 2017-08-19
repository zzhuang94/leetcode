/**
 * Ugly Number

 * Write a program to check whether a given number is an ugly number.
 *
 * Ugly numbers are positive numbers whose prime factors only include 2, 3, 5
 * For example, 6, 8 are ugly while 14 is not ugly since it includes another prime factor 7.
 *
 * Note that 1 is typically treated as an ugly number.
 * Credits:
 * Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
 */

package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, isUgly(i))
	}
	i := 10031234234234254
	fmt.Println(i, isUgly(i))
}

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for {
		if num == 1 {
			return true
		} else if num%2 == 0 {
			num /= 2
		} else if num%3 == 0 {
			num /= 3
		} else if num%5 == 0 {
			num /= 5
		} else {
			return false
		}
	}
}

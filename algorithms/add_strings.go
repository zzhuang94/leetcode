/**
 * Add Strings
 *
 * Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.
 *
 * Note:
 * The length of both num1 and num2 is < 5100.
 * Both num1 and num2 contains only digits 0-9.
 * Both num1 and num2 does not contain any leading zero.
 * You must not use any built-in BigInteger library or convert the inputs to integer directly.
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addStrings("999", "1"))
}

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	ret := ""
	bit := 0
	for i := 1; ; i++ {
		b1, b2 := 0, 0
		if l1-i >= 0 {
			b1 = int(num1[l1-i] - 48)
		}
		if l2-i >= 0 {
			b2 = int(num2[l2-i] - 48)
		}
		t := bit + b1 + b2
		if l1-i <= 0 && l2-i <= 0 {
			return strconv.Itoa(t) + ret
		}
		bit = t / 10
		ret = strconv.Itoa(t%10) + ret
	}
	return ret
}

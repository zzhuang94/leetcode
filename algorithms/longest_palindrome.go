/**
 * Longest Palindrome
 *
 * Given a string which consists of lowercase or uppercase letters,
 * find the length of the longest palindromes that can be built with those letters.
 * This is case sensitive, for example "Aa" is not considered a palindrome here.
 * Note:
 * Assume the length of given string will not exceed 1,010.
 *
 * Example:
 * Input:
 * "abccccdd"
 * Output:
 * 7
 * Explanation:
 * One longest palindrome that can be built is "dccaccd", whose length is 7.
 */

package main

import "fmt"

func main() {
	fmt.Println(byte('A'))
	fmt.Println(byte('Z'))
	fmt.Println(byte('a'))
	fmt.Println(byte('z'))
	fmt.Println(longestPalindrome("zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"))
}

func longestPalindrome(s string) int {
	l := make([]int, 68)
	for _, c := range s {
		l[c-65]++
	}
	sum := 0
	f := true
	for _, v := range l {
		if f && v%2 == 1 {
			sum++
			f = false
		}
		sum += v / 2 * 2
	}
	return sum
}

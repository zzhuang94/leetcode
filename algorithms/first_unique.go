/*
 * First Unique Character in a String
 *
 * Given a string, find the first non-repeating character in it and return it's index.
 * If it doesn't exist, return -1.
 *
 * Examples:
 *
 * s = "leetcode"
 * return 0.
 *
 * s = "loveleetcode",
 * return 2.
 * Note: You may assume the string contain only lowercase letters.
 */

package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("aaaa"))
	fmt.Println(firstUniqChar("abcd"))
	fmt.Println(firstUniqChar("leeztcodel"))
}

func firstUniqChar(s string) int {
	t := make([]int, 26)
	for _, b := range s {
		t[b-97]++
	}
	for i, b := range s {
		if t[b-97] == 1 {
			return i
		}
	}
	return -1
}

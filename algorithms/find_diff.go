/**
 * Find the Difference
 *
 * Given two strings s and t which consist of only lowercase letters.
 * String t is generated by random shuffling string s and then add one more letter at a random position.
 * Find the letter that was added in t.
 *
 * Example:
 * Input:
 * s = "abcd"
 * t = "abcde"
 * Output:
 * e
 *
 * Explanation:
 * 'e' is the letter that was added.
 */

package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("cabcd", "cbcbda"))
}

func findTheDifference(s string, t string) byte {
	m := make(map[rune]int)
	for _, c := range s {
		_, ok := m[c]
		if ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	for _, c := range t {
		v, ok := m[c]
		if !ok || v == 0 {
			return byte(c)
		} else {
			m[c]--
		}
	}
	return 0
}
/*
 * Isomorphic Strings
 *
 * Given two strings s and t, determine if they are isomorphic.
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 * All occurrences of a character must be replaced with another character while preserving the
 * order of characters. No two characters may map to the same character but a character may map to itself.
 *
 * For example,
 * Given "egg", "add", return true.
 * Given "foo", "bar", return false.
 * Given "paper", "title", return true.
 *
 * Note:
 * You may assume both s and t have the same length.
 */

package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("ccec", "aaba"))
	fmt.Println(isIsomorphic("acb", "cba"))
}

func isIsomorphic(s string, t string) bool {
	l := len(s)
	if l != len(t) {
		return false
	}
	ms := make(map[byte]byte)
	mt := make(map[byte]byte)
	for i := 0; i < l; i++ {
		ss := s[i]
		tt := t[i]
		cs, ok := ms[ss]
		if !ok {
			ms[ss] = tt
		} else if cs != tt {
			return false
		}
		ct, ok := mt[tt]
		if !ok {
			mt[tt] = ss
		} else if ct != ss {
			return false
		}
	}
	return true
}

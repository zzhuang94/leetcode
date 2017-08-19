/**
 * Word Pattern
 *
 * Given a pattern and a string str, find if str follows the same pattern.
 * Here follow means a full match, such that there is a bijection
 * between a letter in pattern and a non-empty word in str.
 *
 * Examples:
 * pattern = "abba", str = "dog cat cat dog" should return true.
 * pattern = "abba", str = "dog cat cat fish" should return false.
 * pattern = "aaaa", str = "dog cat cat dog" should return false.
 * pattern = "abba", str = "dog dog dog dog" should return false.
 * Notes:
 * You may assume pattern contains only lowercase letters, and str
 * contains lowercase letters separated by a single space.
 *
 * Credits:
 * Special thanks to @minglotus6 for adding this problem and creating all test cases.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abc", "fds fsdf sdf"))
}

func wordPattern(pattern string, str string) bool {
	if pattern == "" && str == "" {
		return true
	}
	l := len(pattern)
	arr := strings.Split(str, " ")
	if l != len(arr) {
		return false
	}
	mp := make(map[byte]string)
	ms := make(map[string]byte)
	for i := 0; i < l; i++ {
		p := pattern[i]
		s := arr[i]
		vp, ok := mp[p]
		if !ok {
			mp[p] = s
		} else if vp != s {
			return false
		}
		vs, ok := ms[s]
		if !ok {
			ms[s] = p
		} else if vs != p {
			return false
		}
	}
	return true
}

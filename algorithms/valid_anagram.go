/*
 * Valid Anagram
 *
 * Given two strings s and t, write a function to determine if t is an anagram of s.
 *
 * For example,
 * s = "anagram", t = "nagaram", return true.
 * s = "rat", t = "car", return false.
 *
 * Note:
 * You may assume the string contains only lowercase alphabets.
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your solution to such case?
 */

package main

import "fmt"

func main() {
	fmt.Println(isAnagram("kobe", "okbe"))
	fmt.Println(isAnagram("zhuangzhuohuang", "huangzhuozhuang"))
	fmt.Println(isAnagram("hahah", "hhhaa"))
}

func isAnagram(s string, t string) bool {
	l := len(s)
	if l != len(t) {
		return false
	}
	ms := make(map[byte]int)
	mt := make(map[byte]int)
	for i := 0; i < l; i++ {
		si := s[i]
		ti := t[i]
		_, ok := ms[si]
		if ok {
			ms[si]++
		} else {
			ms[si] = 1
		}
		_, ok = mt[ti]
		if ok {
			mt[ti]++
		} else {
			mt[ti] = 1
		}
	}
	for k, v := range ms {
		cc, ok := mt[k]
		if !ok || cc != v {
			return false
		}
	}
	return true
}

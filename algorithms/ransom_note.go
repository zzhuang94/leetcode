/**
 * Ransom Note
 *
 * Given an arbitrary ransom note string and another string containing letters
 * from all the magazines, write a function that will return true if the ransom
 * note can be constructed from the magazines ; otherwise, it will return false.
 * Each letter in the magazine string can only be used once in your ransom note.
 *
 * Note:
 * You may assume that both strings contain only lowercase letters.
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 */

package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("", ""))
}

func canConstruct(ransomNote string, magazine string) bool {
	mr := make(map[rune]int)
	for _, b := range ransomNote {
		_, ok := mr[b]
		if ok {
			mr[b]++
		} else {
			mr[b] = 1
		}
	}
	mm := make(map[rune]int)
	for _, b := range magazine {
		_, ok := mm[b]
		if ok {
			mm[b]++
		} else {
			mm[b] = 1
		}
	}
	for k, v := range mr {
		c, ok := mm[k]
		if !ok || v > c {
			return false
		}
	}
	return true
}

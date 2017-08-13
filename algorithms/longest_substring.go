/**
 * Given a string, find the length of the longest substring without repeating characters.
 *
 * Examples:
 * Given "abcabcbb", the answer is "abc", which the length is 3.
 * Given "bbbbb", the answer is "b", with the length of 1.
 * Given "pwwkew", the answer is "wke", with the length of 3.
 * Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

package main

import "fmt"

func main() {
	s := "kobebryantzzhuang94"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	len := len(s)
	max := 0
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if s[i] == s[j] {
				if max < j {
					t_m := lengthOfLongestSubstring(s[0:j])
					if max < t_m {
						max = t_m
					}
				}
				if max < len-i-1 {
					t_m := lengthOfLongestSubstring(s[i+1 : len])
					if max < t_m {
						max = t_m
					}
				}
				return max
			}
		}
	}
	return len
}

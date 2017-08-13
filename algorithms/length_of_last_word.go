/**
 * Length of Last Word
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 * Note: A word is defined as a character sequence consists of non-space characters only.
 *
 * For example,
 * Given s = "Hello World",
 * return 5.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "kobe  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	s_l := len(s)
	if s_l == 0 {
		return 0
	}
	i := strings.LastIndex(s, " ")
	return s_l - i - 1
}

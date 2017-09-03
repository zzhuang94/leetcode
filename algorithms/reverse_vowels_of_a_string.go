/**
 * Reverse Vowels of a String
 *
 * Write a function that takes a string as input and reverse only the vowels of a string.
 *
 * Example 1:
 * Given s = "hello", return "holle".
 *
 * Example 2:
 * Given s = "leetcode", return "leotcede".
 *
 * Note:
 * The vowels does not include the letter "y".
 */

package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("Aa"))
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	m := map[rune]byte{'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	x := []rune(s)
	for a, b := 0, len(x)-1; a < b; {
		_, oka := m[x[a]]
		if !oka {
			a++
		}
		_, okb := m[x[b]]
		if !okb {
			b--
		}
		if oka && okb {
			x[a], x[b] = x[b], x[a]
			a++
			b--
		}
	}
	return string(x)
}

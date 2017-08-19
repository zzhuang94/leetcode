/**
 * Reverse String
 *
 * Write a function that takes a string as input and returns the string reversed.
 *
 * Example:
 * Given s = "hello", return "olleh".
 */

package main

import "fmt"

func main() {
	fmt.Println(reverseString("kobe"))
	fmt.Println(reverseString("abc"))
	fmt.Println(reverseString(""))
}

func reverseString(s string) string {
	a := []rune(s)
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-1-i] = a[l-1-i], a[i]
	}
	return string(a)
}

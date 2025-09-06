package main

import (
	"fmt"
	"unicode"
)

// 0 ms
func isPalindrome(s string) bool {
	str := []rune{}
	for _, l := range []rune(s) {
		if unicode.IsDigit(l) || unicode.IsLetter(l) {
			str = append(str, unicode.ToLower(l))
		}
	}

	swapped := []rune{}
	i := len(str) - 1
	for i >= 0 {
		swapped = append(swapped, str[i])
		i--
	}
	return string(swapped) == string(str)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}

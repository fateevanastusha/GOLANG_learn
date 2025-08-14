package main

import (
	"fmt"
)

func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res int
	n := len(s)

	for i := 0; i < n; i++ {
		v := symbols[s[i]]
		if (i+1 < n) && (symbols[s[i+1]] > v) {
			res += symbols[s[i+1]] - v
			i++
		} else {
			res += v
		}
	}
	return res
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))

}

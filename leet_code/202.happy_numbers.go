package main

import (
	"fmt"
)

// 0 ms
func isHappy(n int) bool {
	for n != 1 && n != 2 && n != 3 && n != 4 {
		res := 0
		for n > 0 {
			res += (n % 10) * (n % 10)
			n = n / 10
		}
		n = res
	}
	return n == 1
}

func main() {
	fmt.Println(isHappy(49))
}

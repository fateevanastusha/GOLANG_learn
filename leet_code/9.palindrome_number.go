package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverseNum := func(n int) int {
		res := 0
		for n > 0 {
			res = (res * 10) + (n % 10)
			n = n / 10
		}
		return res
	}

	return reverseNum(x) == x
}

func main() {
	fmt.Println(isPalindrome(4321234))

}

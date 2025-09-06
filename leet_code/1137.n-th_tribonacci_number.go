package main

import "fmt"

// 0ms
func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a, b, c := 0, 1, 1
	for i := 1; i < n-1; i++ {
		a, b, c = b, c, a+b+c
	}

	return c
}

func main() {
	n := 25
	fmt.Println(tribonacci(n))
}

package main

import "fmt"

// 0ms
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return b

}

func main() {
	n := 3
	fmt.Println(fib(n))
}

package main

import "fmt"

// func climbStairs(n int) int {
// 	var c func(start int) int
// 	c = func(start int) int {
// 		if start > n {
// 			return 0
// 		}
// 		if start == n {
// 			return 1
// 		}
// 		return c(start+1) + c(start+2)
// 	}
// 	return c(0)
// }

// 0 ms
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2

	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(climbStairs(5))
}

//start = 0
//end = n

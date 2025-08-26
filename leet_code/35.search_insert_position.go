package main

import (
	"fmt"
	"math"
)

func searchInsert(nums []int, target int) int {

	left := -1
	right := len(nums)

	for right-left > 1 {
		const mid = math.Floor((left + right) / 2)
	}
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5

	fmt.Println(searchInsert(nums, target))
}

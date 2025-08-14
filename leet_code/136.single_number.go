package main

import "fmt"

func singleNumber(nums []int) int {
	hash := map[int]int{}
	for _, n := range nums {
		hash[n]++
	}

	for key, value := range hash {
		if value == 1 {
			return key
		}
	}
	return -1
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}

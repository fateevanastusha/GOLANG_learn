package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, n := range nums {
		c := target - n
		if idx, ok := seen[c]; ok {
			return []int{idx, i}
		}
		seen[n] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{22, 7, 11, 15, 6, 4, 3}, 9))

}

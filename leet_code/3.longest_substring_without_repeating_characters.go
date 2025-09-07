package main

import "fmt"

// 1 ms (best res)
func lengthOfLongestSubstring(s string) int {
	max := 0
	left := 0
	sr := []rune(s)
	seen := make(map[rune]int)
	for right := range sr {
		rightS := sr[right]
		for seen[rightS] != 0 {
			delete(seen, sr[left])
			left++
		}
		seen[rightS] = 1
		res := right - left + 1
		if max < res {
			max = res
		}
	}
	return max
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

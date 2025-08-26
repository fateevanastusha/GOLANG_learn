package main

import (
	"fmt"
	"strings"
)

// 0 ms
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	h := "sadbutsad"
	n := "sad"

	fmt.Println(strStr(h, n))
}

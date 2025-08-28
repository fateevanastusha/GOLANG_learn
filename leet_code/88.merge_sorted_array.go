package main

import (
	"fmt"
	"math"
)

func quickSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	mid := math.Ceil(float64(length)/2) - 1

	var less, equal, greater []int
	pivot := nums[int(mid)]
	for _, el := range nums {
		if el > pivot {
			greater = append(greater, el)
		}
		if el < pivot {
			less = append(less, el)
		}
		if el == pivot {
			equal = append(equal, el)
		}
	}

	result := append(quickSort(less), equal...)
	result = append(result, quickSort(greater)...)

	return result
}

// 1 ms
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2[:n])
	newArr := quickSort(nums1)
	for i := range nums1 {
		nums1[i] = newArr[i]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

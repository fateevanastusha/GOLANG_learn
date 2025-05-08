package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(v uint) uint {
		var res string

		for v > 0 {
			num := v % 10
			if num != 0 && num%2 == 0 {
				res = strconv.Itoa(int(num)) + res
			}
			v = v / 10
		}
		if res == "" {
			return 100
		}
		returnValue, _ := strconv.Atoi(res)
		return uint(returnValue)
	}
	fmt.Println(fn(0))
	fmt.Println(fn(727178))
	fmt.Println(fn(55555))
}

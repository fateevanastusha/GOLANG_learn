package main

import (
	"fmt"
	"strings"
)

// 4ms
// func longestCommonPrefix(strs []string) string {
// 	r := []rune(strs[0])
// 	for _, s := range strs[1:] {
// 		s := []rune(s)
// 		//сначала обрезаем лишние символы, если r > s, то обрезаем r до длины s, и наоборот
// 		sl := len(s)
// 		rl := len(r)
// 		if sl > rl {
// 			s = s[:rl]
// 		}
// 		if rl > sl {
// 			r = r[:sl]
// 		}
// 		//сравниваем текущую строку с r с конца. если символы не равны - удаляем символ
// 		nSlice := len(s)
// 		i := len(s)
// 		for i >= 0 {
// 			if string(s[0:i]) != string(r[0:i]) {
// 				nSlice--
// 			}
// 			i--
// 		}
// 		r = r[:nSlice]
// 	}
// 	return string(r)
// }

// 0ms
func longestCommonPrefix(strs []string) string {
	r := []rune(strs[0])
	strs = strs[1:]
	n := len(strs)
	str := " " + strings.Join(strs, " ")
	i := len(r) - 1
	//смотрим сколкьо раз встречается " flower", если меньше < n, то обрезаем на один символ и снова проверяем
	for i >= 0 {
		if strings.Count(str, " "+string(r)) != n {
			r = r[0:i]
		} else {
			break
		}
		i--
	}

	return string(r)
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

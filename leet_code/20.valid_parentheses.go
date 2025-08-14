package main

import (
	"fmt"
	"strings"
)

// 4ms
// func isValid(str string) bool {
// 	for strings.Contains(str, "{}") || strings.Contains(str, "[]") || strings.Contains(str, "()") {
// 		str = strings.ReplaceAll(str, "{}", "")
// 		str = strings.ReplaceAll(str, "()", "")
// 		str = strings.ReplaceAll(str, "[]", "")
// 	}
// 	return len(str) == 0
// }

// 0ms
func isValid(str string) bool {
	stack := []string{}
	m := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	//следующий элемент - либо закрывающая для текущего, либо открывающая. если первый случай - делаем pop (снимаем стэк)
	for _, char := range []rune(str) {
		s := string(char)
		//если открывающая или первая - просто пушим
		if len(stack) == 0 || strings.Contains("{([", s) {
			stack = append(stack, s)
			continue
		}

		//если закрывающая, чекаем, является ли предыдущая открывающей для нее
		lastValue := stack[len(stack)-1]
		if m[lastValue] != s {
			//если нет - ошибка, выходим
			return false
		}
		//если да - просто снимаем их со стека
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0

}

func main() {
	s := "{{{{}"
	fmt.Println(isValid(s))
}

package main

import "fmt"

func main() {
	/*
		IF
	*/
	if true {
		fmt.Println(1)
	}

	//можно объявить переменную прямо в блоке if
	if v := 23 + 6; v > 25 {
		fmt.Println(2)
	}

	//можно несколько условий

	var a, b = 78, 78

	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}

	/*
		SWITCH
		- в switch можно использовать любой тип данных
	*/

	var i = 2
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

	//в go нет слова break, зато есть обратное ему - fallthrough, если его использовать, то он повалится дальше вниз
	v := 42
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}
	// Вывод:
	// 42
	// 1
	// default

	//в go можно использовать произвольные условия в каждом блоке case

	var c uint32
	fmt.Scan(&c)
	switch {
	case 1 <= c && c <= 9:
		fmt.Println("от 1 до 9")
	case 100 <= c && c <= 250:
		fmt.Println("от 100 до 250")
	case 1000 <= c && c <= 6000:
		fmt.Println("от 1000 до 6000")
	}

}

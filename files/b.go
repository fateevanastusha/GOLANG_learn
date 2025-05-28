package main

import (
	"fmt"
	"log"
	"os"
)

var filename = "new.txt"

func main() {
	/*
		вывод stdout, stderr -
		это файлы в операционной системе, в которые попадает все, что пишется в консоль.
		под капотом -
		var (
			Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
			Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
			Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
		)
		Stdin
		Stdout - поток, выводящий в терминал
		Stderr - то же самое, что поток вывода, только с ошибками

		получается, когда мы печатаем что-то через fmt, то мы записываем эти данные в файл stdout, к которому
		привязана консоль - поэтому мы видим их в консоли, либо когда делаем scan - то читаем данные из файла
		stdin.
	*/

	fmt.Println("-------------------------------------------------------------------------------------------------------------")
	output()
	fmt.Println("-------------------------------------------------------------------------------------------------------------")
	input()
	fmt.Println("-------------------------------------------------------------------------------------------------------------")
	args()
	fmt.Println("-------------------------------------------------------------------------------------------------------------")

}

func output() {
	fmt.Println("OUTPUT FUNCTION\n")

	count, err := fmt.Println("some text")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)

	/*
		если мы хотим записать данные в stdout, то нужно использовать как io.Writer - os.Stdout (c stderr аналогично).
		получается, что os.Stdout имплементирует интерфейс io.Writer.
		Таким образом можно будет перенаправлять логи в другие места, из этих файлов - stderr, stdout.
	*/

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//fmt.Println - всего лишь обертка над Fprintln, которая передает ему первым аргументом os.Stdout. Мы можем использовать
	//его в своих целях и передать что-то другое имплементирующее интерфейс io.Writer.
	count, err = fmt.Fprintln(file, "hello\nworld!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func input() {

	fmt.Println("INPUT FUNCTION\n")

	/*
		fmt.Scan получает данные из файла stdin, все что мы туда вводим - в итоге оказывается здесь. fmt.Scan - это обертка
		над Fscan, в который он первым аргументом передает os.Stdin (который имплементирует интерфейс io.Reader). Затем он
		читает файл stdin (или другой, который ему передать).
	*/

	var text1, text2 string

	fmt.Println("please enter TWO values:")
	count, err := fmt.Scan(&text1, &text2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TEXTS = ", text1, text2, " COUNT = ", count) // "" "" [кол-во символов]

	fmt.Println("please enter TWO values:")
	count, err = fmt.Fscan(os.Stdin, &text1, &text2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TEXTS =", text1, text2, "COUNT =", count) // "" "" [кол-во символов]

	//попробуем передать другой файл, кроме os.Stdin
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	//считывает по одной строке, закончит работу как только найдет \n
	count, err = fmt.Fscanln(file, &text1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TEXTS =", text1, "COUNT =", count) //TEXTS = hello COUNT = 1

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

}

func args() {

	fmt.Println("AGRS FUNCTION\n")

	/*
		os.Args содержит слайс строк, аргументов запуска приложения. Все, что пишется при запуске после названия
		файла - и есть аргументы (пример - go run main.go test)
	*/

	for _, arg := range os.Args {
		fmt.Println("this is arg =", arg)
	}
}

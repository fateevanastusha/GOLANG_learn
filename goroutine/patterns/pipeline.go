package main

import (
	"fmt"
	"time"
)

/*
	pipeline
	переливаем один канал в другой
*/

func main() {
	fmt.Println("PIPELINE\n")

	reader(doubler(writer()))
}

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

func doubler(readCh <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for v := range readCh {
			time.Sleep(500 * time.Millisecond)
			ch <- v * 2
		}

		close(ch)
	}()

	return ch
}

func reader(ch <-chan int) {

	for range 10 {
		fmt.Println(<-ch)
	}

}

/*
	Напишите функции merge и fillChan.

	Функция fillChan:
		•	на вход получает целое число n;
		•	возвращает канал;
		•	пишет в этот канал n чисел от 0 до n-1.

	Функция merge:
		•	получает на вход массив каналов cs;
		•	возвращает канал;
		•	параллельно читает из каждого канала из cs и пишет полученное значение в возвращаемый канал.

	// merge - соединяет каналы в один
	func merge(cs ...<-chan int) <-chan int {
		// напишите ваш код здесь
	}

	// fillChan - заполняет канал числами от 0 до n-1
	func fillChan(n int) <-chan int {
		// напишите ваш код здесь
	}

	func main() {
		a := fillChan(2)
		b := fillChan(3)
		c := fillChan(4)
		d := merge(a, b, c)
		for v := range d {
			fmt.Println(v)
		}
	}
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)
	d := merge(a, b, c)
	for v := range d {
		fmt.Println(v)
	}
	fmt.Println(time.Since(start).Round(time.Microsecond))
}

func fillChan(n int) <-chan int {
	channel := make(chan int, n)
	go func() {
		defer close(channel)
		for i := 0; i < n; i++ {
			channel <- i
		}
	}()

	return channel
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, channel := range cs {
		go func(channel <-chan int) {
			defer wg.Done()
			for v := range channel {
				out <- v
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

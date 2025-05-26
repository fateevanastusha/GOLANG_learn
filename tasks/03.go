/*

	Напишите функции generator и squarer.

	Функция generator принимает на вход контекст и слайс целых чисел, элементы которого
	последовательно записываются в возвращаемый канал.
	Функция square принимает на вход контекст и канал целых чисел. Функция последовательно
	читает из канала числа, возводит их в квадрат и пишет в возвращаемый канал.

	Обе функции должны уметь завершаться по отмене контекста.

	func main() {
		ctx := context.Background()
		pipeline := squarer(ctx, generator(ctx, 1, 2, 3))
		for x := range pipeline {
			fmt.Println(x)
		}
	}

	func generator(ctx context.Context, in ...int) <-chan int {
		// напишите ваш код здесь
	}

	func squarer(ctx context.Context, in <-chan int) <-chan int {
		// напишите ваш код здесь
	}

*/

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))
	for x := range pipeline {
		fmt.Println(x)
	}
}

func generator(ctx context.Context, in ...int) <-chan int {
	// напишите ваш код здесь
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	// напишите ваш код здесь
}

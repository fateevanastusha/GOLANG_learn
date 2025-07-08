package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//SELECT CASE

	/*
		- выбирает любой из готовых case, где операция с каналом может быть выполнена немедленно — будь то чтение или запись.
		1.	Готовые операции — select проверяет, какие case могут быть выполнены прямо сейчас:
			•	Если можно прочитать из канала (в нём есть данные) → этот case готов.
			•	Если можно записать в канал (в нём кто-то ждет или он не заблокирован) → этот case готов.
		2.	Случайный выбор — если несколько case готовы одновременно, Go выбирает один случайным образом.
		3.	Блокировка — если ни один case не готов, select блокируется и ждет, пока хотя бы один case не станет готовым.
		4.	default (если есть) — выполняется немедленно, если никакой другой case не готов.

		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
			}
		}

		пример использования:
	*/

	fibonacci := func(c, quit chan int) {
		x, y := 0, 1
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
			}
		}
	}

	chanf := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-chanf)
		}
		quit <- 0
	}()
	fibonacci(chanf, quit)

	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// quit

	//-------------------------------------------------------------------------------------------------------------------

	/*
		select - блокирующий оператор.
		1) (v := <-ch) - ждет пока кто-то напишет в какой-то канал, если напишет одновременно - то сработает рандомный
		канал. Если в каналы никто никогда не напишет - дедлок. Отработает один раз и выйдет.
		2) (<-time.After(1 * time.Second)) - выйдет через отрезок времени, если другие кейсы до этого момента не сработают.
		3) (<-timer.C) - сработает когда таймер закончится.
		4) (<-ctx.Done()) - сработает когда контекст будет Done.
		5) (<-ch3) - выйдет когда канал закроется.
		6) default - сработает если все остальные операции заблокированы, он не будет дожидаться (даже если где-то пытаемся
		дописать значения выше).

		- по итогу из всех сработает тот, что завершится первый. Все пункты работают по принципу закрытия канала.
	*/

	selectF := func() {
		ch1 := make(chan int)
		ch2 := make(chan int)
		timer := time.NewTimer(1 * time.Millisecond)
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond)
		defer cancel()
		ch3 := make(chan int)
		close(ch3)

		select {
		case v := <-ch1: //1
			fmt.Println("v =", v, "from ch1")
		case v := <-ch2:
			fmt.Println("v =", v, "from ch2")
		case <-time.After(1 * time.Second): //2
			fmt.Println("exited by after")
		case <-timer.C: //3
			fmt.Println("exited by timer")
		case <-ctx.Done(): //4
			fmt.Println("exited by context")
		case <-ch3: //5
			fmt.Println("exited by channel")
		default: //6
		}
	}

	selectF()
}

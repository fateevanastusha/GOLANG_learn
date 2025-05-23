package main

import (
	"fmt"
	"time"
)

func main() {

	//1) Sleep

	// func Sleep(d Duration)
	// программа засыпает на заданное время
	time.Sleep(time.Second * 2) // спим ровно 2 секунды

	fmt.Println("after 2 seconds") //увидим через 2 секунды

	fmt.Println("----------------------------------------------------------------------------------------------")

	//2) After

	// func After(d Duration) <-chan Time - возвращает readonly канал
	/*
		создает канал, который через заданное время вернет значение. Ожидание происходит асинхронно, поэтому
		пока может выполняться другой код.
	*/
	timer := time.After(time.Second * 2)
	go func() {
		fmt.Println(<-timer) // значение будет получено из канала ровно через 2 секунды (тип значения time.Time)
	}()

	/*
		Этой фукнцией очень полезно пользоваться для таймаута. Например, мы ждем ответа от сети максимум 5 секунд -
		если за 5 секунд ответа нет, то выкидываем ошибку. Сделать это можно так:

		timeout := time.After(2 * time.Second)

		for i := 0; i < 5; i++ {

			select { // Оператор select

			case gopherID := <-c: // Ждет, когда проснется гофер

				fmt.Println("gopher ", gopherID, " has finished sleeping")

			case <-timeout: // Ждет окончания времени

				fmt.Println("my patience ran out")

				return // Сдается и возвращается

			}

		}

		- если гофер не успеет проснуться за 2 секунды, то терпение заканчиватется. Если успеет - то пишет, что
		гофер наконец проснулся.
		Данный паттерн полезен, когда вам нужно ограничить время на выполнения определенной операции. Поместив действие
		внутрь горутины и отправив его каналу, когда тот завершен, можно добиться фиксированного времени практически для
		всего в Go.
	*/

	time.Sleep(time.Second * 3)

	fmt.Println("----------------------------------------------------------------------------------------------")

	//3) Tick
	// func Tick(d Duration) <-chan Time - возвращает readonly канал
	/*
		создает канал, который будет посылать сигналы постоянно через заданный промежуток времени. Канал
		будет небуферизированный, поэтому пока не возьмешь значение - следующее он не отправит, поэтому
		значения накапливаться не могут.
	*/
	ticker := time.Tick(time.Second * 2)
	//прослойка, чтобы следить сколько там значений
	countTicker := make(chan time.Time, 10)
	count := 0

	for {
		countTicker <- <-ticker
		fmt.Println("очередной тик")
		count++
		if count == 3 {
			break
		}
	}

	fmt.Println("after tick")

	// очередной тик
	// очередной тик
	// очередной тик

	fmt.Println("----------------------------------------------------------------------------------------------")

	//4) Timer
	/*
		- похож на After, но можно вторнуться в его работу. Во время выполнения его можно
		остановить или изменить время его выполнения.
	*/

	t := time.NewTimer(time.Second) // создаем новый таймер, который сработает через 1 секунду
	go func() {
		<-t.C // C - канал, который должен вернуть значение через заданное время
	}()
	t.Stop() // но мы можем остановить таймер и раньше установленного времени

	t.Reset(time.Second * 2) // пока таймер не сработал, мы можем сбросить его, установив новый срок выполнения
	<-t.C

	fmt.Println("----------------------------------------------------------------------------------------------")

	//5) Ticker
	/*
		- работает так же как и Tick, но может быть остановлен в процессе работы.
		func NewTicker(d Duration) *Ticker // создаем новый Ticker
		func (t *Ticker) Stop() // останавливаем Ticker
	*/

	//пример
	work := func() <-chan struct{} { //возвращает readonly канал, с any типом значений в нем
		done := make(chan struct{}) // канал для синхронизации горутин

		go func() {
			defer close(done) // синхронизирующий канал будет закрыт, когда функция завершит свою работу

			stop := time.NewTimer(time.Second * 2) //через сколько завершить работу

			tick := time.NewTicker(time.Millisecond * 200)
			defer tick.Stop() // освободим ресурсы, при завершении работы функции

			for {
				select {
				case <-stop.C: // C - канал, который должен вернуть значение через заданное время
					// stop - Timer, который через 1 секунду даст сигнал завершить работу
					return
				case <-tick.C: // C - канал, который должен вернуть значение через заданное время
					// tick - Ticker, посылающий сигнал выполнить работу каждый 200 миллисекунд
					fmt.Println("тик-так")
				}
			}
		}()

		return done
	}

	<-work()
	/*
	 * тик-так
	 * тик-так
	 * тик-так
	 * тик-так
	 */

	fmt.Println("----------------------------------------------------------------------------------------------")

}

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Модуль time стандартной библиотеки предоставляет инструменты для работы с датой и временем,
		включая создание объектов времени, их форматирование и преобразование.
	*/

	//1) type Time
	/*
		Структура Time в Go используется для представления конкретных даты и времени. Мы можем создавать
		экземпляры этой структуры с помощью различных функций из пакета time.
	*/

	// Получаем текущее время
	currentTime := time.Now()

	// Создаем время с помощью конкретных значений
	customTime := time.Date(
		2020,     // год
		time.May, // месяц
		15,       // день
		10,       // часы
		13,       // минуты
		12,       // секунды
		45,       // наносекунды
		time.UTC, // временная зона
	)

	// Создаем время, используя секунды и наносекунды, прошедшие с начала эпохи Unix
	unixTime := time.Unix(
		150000, // секунды
		1,      // наносекунды
	)

	// Форматируем и выводим время в строковом виде
	// Format принимает строку-шаблон, преобразует структуру Time в строку по этому шаблону
	fmt.Println(currentTime.Format("02-01-2006 15:04:05")) // текущее время в указанномы dd-mm-yyyy hh:mm:ss
	fmt.Println(customTime.Format("02-01-2006 15:04:05"))  // 15-05-2020 10:13:12
	fmt.Println(unixTime.Format("02-01-2006 15:04:05"))    // 02-01-1970 22:40:00

	/*
		Шаблон всегда основан на дате "Mon Jan 2 15:04:05 MST 2006" (понедельник, 2 января 2006 года, 15:04:05,
		время в часовом поясе MST). Логика этих значений в следующем:
		Месяц, первый месяц - Январь = 01
		День, второй день месяца = 02
		Часы, 15:00 (для 24 часа) или 03:00 (для 12) = 15
		Минуты, 4-ая минута = 04
		Секунды, 5-я секунды = 05
		Год, 6-й год (2006 или 06) = 2006
		Часовой пояс

		Также в time есть константы-шаблоны:

		Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
		ANSIC       = "Mon Jan _2 15:04:05 2006"
		UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
		RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
		RFC822      = "02 Jan 06 15:04 MST"
		RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
		RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
		RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
		RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
		RFC3339     = "2006-01-02T15:04:05Z07:00"
		RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		Kitchen     = "3:04PM"

		// Handy time stamps.
		Stamp      = "Jan _2 15:04:05"
		StampMilli = "Jan _2 15:04:05.000"
		StampMicro = "Jan _2 15:04:05.000000"
		StampNano  = "Jan _2 15:04:05.000000000"
		DateTime   = "2006-01-02 15:04:05"
		DateOnly   = "2006-01-02"
		TimeOnly   = "15:04:05"

		вызывать их так: time.[template_name]
	*/

	fmt.Println("----------------------------------------------------------------------------------------------")

	//2) конвертация строк в Time

	// Функция Parse парсит строку в соответствии с заданным шаблоном
	firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
	if err != nil {
		panic(err)
	}

	// LoadLocation находит временную зону в справочнике IANA
	// https://www.iana.org/time-zones
	loc, err := time.LoadLocation("Asia/Yekaterinburg")
	if err != nil {
		panic(err)
	}

	// Функция ParseInLocation парсит строку в указанной временной зоне
	secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
	if err != nil {
		panic(err)
	}

	fmt.Println(firstTime.Format(time.RFC3339))  // 2020-05-15T17:45:00Z
	fmt.Println(secondTime.Format(time.RFC3339)) // 2020-05-15T17:45:10+05:00

	fmt.Println("----------------------------------------------------------------------------------------------")

	//3) конвертация Time в строку

	// func (t Time) Format(layout string) string
	c := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	fmt.Println(c.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:12

	fmt.Println("----------------------------------------------------------------------------------------------")

	//4) методы Time

	current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// func (t Time) Date() (year int, month Month, day int)
	fmt.Println(current.Date()) // 2020 May 15

	// func (t Time) Year() int
	fmt.Println(current.Year()) // 2020

	// func (t Time) Month() Month
	fmt.Println(current.Month()) // May

	// func (t Time) Day() int
	fmt.Println(current.Day()) // 15

	// func (t Time) Clock() (hour, min, sec int)
	fmt.Println(current.Clock()) // 17 45 12

	// func (t Time) Hour() int
	fmt.Println(current.Hour()) //17

	// func (t Time) Minute() int
	fmt.Println(current.Minute()) // 45

	// func (t Time) Second() int
	fmt.Println(current.Second()) // 12

	// func (t Time) Unix() int64
	fmt.Println(current.Unix()) // 1589546712

	// func (t Time) Weekday() Weekday
	fmt.Println(current.Weekday()) // Friday

	// func (t Time) YearDay() int
	fmt.Println(current.YearDay()) // 136

	fmt.Println("----------------------------------------------------------------------------------------------")

	//4) сравнение структур Time

	firstTimeRandom := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
	secondTimeRandom := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

	// func (t Time) After(u Time) bool
	// true если позже
	fmt.Println(firstTimeRandom.After(secondTimeRandom)) // true

	// func (t Time) Before(u Time) bool
	// true если раньше
	fmt.Println(firstTimeRandom.Before(secondTimeRandom)) // false

	// func (t Time) Equal(u Time) bool
	// true если равны
	fmt.Println(firstTimeRandom.Equal(secondTimeRandom)) // false

	fmt.Println("----------------------------------------------------------------------------------------------")

	//5) изменение Time

	now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// func (t Time) Add(d Duration) Time
	// изменяет дату в соответствии с параметром - "продолжительностью"
	future := now.Add(time.Hour * 12) // перемещаемся на 12 часов вперед

	// func (t Time) AddDate(years int, months int, days int) Time
	// изменяет дату в соответствии с параметрами - количеством лет, месяцев и дней
	past := now.AddDate(-1, -2, -3) // перемещаемся на 1 год, два месяца и 3 дня назад

	// func (t Time) Sub(u Time) Duration
	// вычисляет время, прошедшее между двумя датами
	fmt.Println(future.Sub(past)) // 10332h0m0s

	/*
		Add и AddDate могут использоваться и отрицательные значения, это позволяет не только «добавлять» время (что видно
		из названий методов), но и «отнимать» его.
	*/

	fmt.Println("----------------------------------------------------------------------------------------------")

	//6) time.Month

	/*
		Это всего лишь объявленные на уровне модуля time константы, которые выглядят следующим образом:

		type Month int

		const (
			January Month = 1 + iota
			February
			March
			April
			May
			June
			July
			August
			September
			October
			November
			December
		)
	*/

	// fmt.Println("----------------------------------------------------------------------------------------------")

	//7) time.Duration

	/*
		представляет из себя int64, определяющий количество наносекунд, прошедших между двумя моментами времени.
	*/

	nowW := time.Now()
	pastW := nowW.AddDate(0, 0, -1)
	futureW := nowW.AddDate(0, 0, 1)

	// func Since(t Time) Duration
	// вычисляет период между текущим моментом и заданным временем в прошлом
	fmt.Println(time.Since(pastW).Round(time.Second)) // 24h0m0s

	// func Until(t Time) Duration
	// вычисляет период между текущим моментом и заданным временем в будущем
	fmt.Println(time.Until(futureW).Round(time.Second)) // 24h0m0s

	// func ParseDuration(s string) (Duration, error)
	// преобразует строку в Duration с использованием аннотаций:
	// "ns" - наносекунды,
	// "us" - микросекунды,
	// "ms" - миллисекунды,
	// "s" - секунды,
	// "m" - минуты,
	// "h" - часы.
	dur, err := time.ParseDuration("1h12m3s")
	if err != nil {
		panic(err)
	}
	fmt.Println(dur.Round(time.Hour).Hours())     // 1
	fmt.Println(dur.Round(time.Minute).Minutes()) // 72
	fmt.Println(dur.Round(time.Second).Seconds()) // 4323

	/*
		так же у типа Duration есть еще другие методы, позволяющие получить часть целых значений
		(часов, минут, секунд и тды)

		func (d Duration) Hours() float64
		func (d Duration) Minutes() float64
		func (d Duration) Seconds() float64
		func (d Duration) Milliseconds() int64
		func (d Duration) Microseconds() int64
		func (d Duration) Nanoseconds() int64

		так же модуль time содержит константы типа Duration
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)
	*/

	fmt.Println("----------------------------------------------------------------------------------------------")

}

package main

import (
	"errors"
	"fmt"
)

type DatePublic struct {
	Year  int
	Month int
	Day   int
}

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year //автоматически получает значение по указателю
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid year")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

type Event struct {
	Date
	Title string
}

func main() {

	date1 := DatePublic{Year: 2019, Month: 5, Day: 27}
	fmt.Println(date1) // {2019 5 27}

	//1) инкапсуляция
	/*
		при создании структуры мы хотим предохранить от того, чтобы юзер ввел некорректные значения - для этого
		есть инкапсуляция. Инкапсуляция в GO реализована посредством импортируемых/неимпортируемых полей. То есть
		прямого механизма - нет (все, что начинается с маленькой буквы, в том числе и поля структуры - не экспорт).

		FILE "date.go" (за пределами рабочей области):

		1) Создаем структуру, названия инкапсулированных полей должны быть обязательно с маленькой буквы, так они
		не могут быть экспортированы. Но сама структура - с большой буквы, так как она должна быть экспортирована.

		type Person struct {
			name string
		}

		2) Добавляем для этого инкапсулированного поля GET и SET методы (геттеры и сеттеры).
		SET - название: [Set][Название_поля_с_большой_буквы]
		GET - название: [Название_поля_с_большой_буквы]

		func (p *Person) SetName (name string) error {
			if len(name) < 10 {
				return errors.New("invalid name")
			}
			p.name = name
			return nil
		}

		func (p *Person) Name () {
			return p.name
		}

		3) Теперь, когда мы обратимся в пакет с структурой Person - мы сможем взаимодействовать с ней только через
		(с полем name) геттеры и сеттеры.


		FILE "main.go":

		import (
			"fmt"
			"github.com/headfirstgo/date"
		)

		func main(){
			date := date.Date{}
			date.SetYear(2019)
			date.SetMonth(11)
			date.SetDay(3)

			fmt.Println(date.Year(), date.Month(), date.Day()) //2019 11 3

			date.SetDay(60) //не сможем установить невалидную дату, вернет ошибку!
		}


	*/

	date2 := Date{}
	err1 := date2.SetYear(2019)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(date2.Year()) //2019
	}

	date3 := Date{}
	err2 := date3.SetYear(-1)
	if err2 != nil {
		fmt.Println(err2) //invalid year
	} else {
		fmt.Println(date3.Year())
	}

	//2) встраивание
	/*
		Если типу структуры потребуются методы, которые уже существуют у другого типа, можно встроить другой тип
		в свой тип структуры, а затем воспользоваться методами встроенного типа так, если бы они были определены
		для вашего собственного типа.

		Встроить можно через анонимное поле -
		type Event struct {
			Date 			// это и есть анонимное поле, получить к нему доступ можно через d.Date.v
			Title string
		}
	*/

	event1 := Event{Title: "event1"}
	//можно обращаться к методам встроенного типа как к собственному, либо через встроенный тип (сработает и так и так):
	event1.SetYear(2019)
	event1.Date.SetDay(2019)

}

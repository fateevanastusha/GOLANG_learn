package main

import (
	"fmt"
	"math"
)

/*
	Стуктуры - это тип данных, который позволяет объединить несколько значений (полей) разных типов
	в один логический объект.


	1) определить
	type Circle struct {
		x, y, r float64
	}



	2) использовать

	в переменной:
	var c Circle
	c := Circle{}
	c := new(Circle)
	ИЛИ С ПРИСВОЕННЫМИ ЗНАЧЕНИЯМИ:
	c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}

	как и с другими типами данных, будет создана локальная переменная типа Circle, и ее поля по
	умолчанию будут равны нулю (0 для int, 0.0 для float, "" для string, nil для указателей, …),
	если мы их не определили.


	в функции, как аргумент:
	func (c Circle)



	3) получить доступ к полям:
	//прочитать
	fmt.Println(c.x, c.y, c.r)
	//переопределить
	c.x = 10
	c.y = 5

	func circleArea(c Circle) float64 {
		//здесь мы используем копии, поэтому оригинальная переменная не изменится
		return math.Pi * c.r * c.r
	}



	4) методы:
	//создать (может быть доступна только для Circle)
	func ([получатель]) [название функции] [тип возвращаемого значения] {}
	func (c *Circle) area() float64 {
		return math.Pi * c.r * c.r //таким образом он сразу имеет доступ к свойствам
	}

	//использовать, go автоматически предоставляет доступ к указателю на Circle для этого метода
	fmt.Println(c.area())



	5) встраиваемые типы (анонимные поля) (наследование):

	допустим, у нас есть ЛИЧНОСТЬ:
	type Person struct {
		Name string
	}
	func (p *Person) Talk() {
		fmt.Println("Hi, my name is", p.Name)
	}

	и мы хотим создать АНДРОИДА, у которого тоже есть ЛИЧНОСТЬ:
	type Android struct {
		Person Person (переменная Person типа Person)
		Model string
	}

	- это будет работать, но мы хотим описать другое отношение. Сейчас - это у АНДРОИДА "есть"
	ЛИЧНОСТЬ, а мы хотим чтобы было АНДРОИД "является" ЛИЧНОСТЬЮ. Такое отношение можно описать
	через встраиваемые типы (анонимные поля):

	type Android struct {
		Person //использовали тип (Person) и не написали его имя
		Model string
	}

	- теперь у АНРОИДА мы имеем доступ ко всем полям и методам ЛИЧНОСТИ:
	a := new(Android)
	ИЛИ
	var a = Android{
		Model: "model",
		Person: Person {
			Name: "name",
		},
	}

	a.Talk() или a.Person.Talk()


*/

func main() {
	rect := Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}
	circ := Circle{x: 0, y: 0, r: 5}

	fmt.Println("Rectangle area:", rect.area())
	fmt.Println("Circle area:", circ.area())
}

type Circle struct {
	x float64
	y float64
	r float64
}

/*
	можно заменить на:
	type Circle struct {
		x, y, r float64
	}
*/

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
	l := r.x2 - r.x1
	w := r.y2 - r.y1
	return l * w
}

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/*

		ФУНКЦИЯ

		Функция представляет блок операторов, которые все вместе выполняют какую-то
		определенную задачу. С помощью функций можно многократно вызывать ее блок
		операторов как единое целое в других частях программы.


		1) объявить
		func имя_функции (список_параметров) (типы_возвращаемых_значений) {
			выполняемые_операторы
		}

		func main() {
		  hello()
		}

		func hello() {
			fmt.Println("Hello World")
		}



		2) параметры функции
		- для каждого параметра задается его имя и тип (аналогично переменным).

		func multiply(a int, b int) {
			result := a * b
			fmt.Println("Результат:", result)
		}

		multiply(3, 7)

		если несколько параметров имеют одинаковый тип, можно указать его только для последнего

		func calculate(x, y int, p, q, r float64) {
			sumInt := x + y
			sumFloat := p + q + r
			fmt.Println("Сумма целых чисел:", sumInt)
			fmt.Println("Сумма дробных чисел:", sumFloat)
		}

		x, y - int
		p, q, r - float64



		3) в go аргументы в функцию передаются по значению, то есть передается их копия. Изменения
		внутри функции не влияет на оригинальную переменную.

		func main() {
			num := 15
			fmt.Println("До вызова функции:", num) //15
			increase(num)
			fmt.Println("После вызова функции:", num) //15
		}

		func increase(n int) {
			fmt.Println("Внутри функции, до увеличения:", n) //15
			n += 10
			fmt.Println("Внутри функции, после увеличения:", n) //25
		}

		До вызова функции: 15
		Внутри функции, до увеличения: 15
		Внутри функции, после увеличения: 25
		После вызова функции: 15



		4) чтобы в функцию передать само значение, а не его копию - нужно передавать указатель
		(адрес этой переменной) через &

		func main() {
			num := 15
			fmt.Println("До вызова функции:", num) //15
			increase(&num)
			fmt.Println("После вызова функции:", num) //25
		}

		func increase(n *int) {
			fmt.Println("Внутри функции, до увеличения:", *n) //15
			*n += 10
			fmt.Println("Внутри функции, после увеличения:", *n) //25
		}

		До вызова функции: 15
		Внутри функции, до увеличения: 15
		Внутри функции, после увеличения: 25
		После вызова функции: 25



		5) возвращение результата из функции:
		используем ключевое слово return

		func имя_функции(параметры) тип_возвращаемого_значения {
			// тело функции
			return возвращаемое_значение
		}

		func subtract(a, b int) int {
			return a - b
		}

		result1 := subtract(15, 8)  // 7
		result2 := subtract(50, 30) // 20



		6) функция может возвращать несколько значений одновременно

		func calculate(a, b int, firstName, lastName string) (int, string) {
			total := a + b
			fullName := firstName + " " + lastName
			return total, fullName
		}

		sum, fullName := calculate(3, 7, "John", "Doe")
		fmt.Println(sum)       // 10
		fmt.Println(fullName)  // John Doe



		7) можно игнорировать возращаемые значения:

		process()
		result, _ := process()
		_, err := process()



		8) аргументов может быть неопределенное количество, например fmt.Print()

		func Print(a ...interface{}) (n int, err error)

		... - указывает, что функция может принять любое количество аргументов данного типа. Все
		переданные аргументы рассматриваются как срез.

		func myPrint(a ...interface{}) {
			for _, elem := range a {
				fmt.Printf("%v ", elem)
			}
		}

		func ExampleMyPrint() {
			myPrint(1, "Hello", 3.14, true)

			// Output:
			// 1 Hello 3.14 true
		}

		myPrint принимает любое количество значений любого типа и выводит их через цикл range для
		обхода среза a.
		Если функция имеет обязательные параметры, за которыми следует переменное количество аргументов,
		то такие аргументы должны быть указаны в конце списка параметров:

		func Fprint(w io.Writer, a ...interface{}) (n int, err error)
		- принимает обязательный параметр типа io.Writer и переменное количество аргументов типа
		interface{} (это пустой интерфейс, позволяющий передавать в функцию аргументы любого типа,
		будь то строки, числа или даже структуры)
	*/

	/*
		Анонимные функции
		- не имеют имени и могут быть использованы прямо в выражениях.
	*/

	//1) использовать
	src := "aBcDeF"

	src = strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)

	}, src)

	fmt.Println("Map result - ", src) //"AbCdEf"

	//2) присвоить переменной
	fn := func(a, b int) int {
		return a + b
	}

	fmt.Println("Fn result - ", fn(45, 32)) //77

	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Println(x)
}

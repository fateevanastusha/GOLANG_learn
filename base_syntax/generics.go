package main

import "fmt"

// -----------------------------------------------------------------------------------------

func main() {
	/*
		Дженерики - "общие типы"
	*/

	showSum()
	showContains()
	showAny()
	unionInterfaceAndType()
	typeApproximation()
}

// -----------------------------------------------------------------------------------------

//ДЖЕНЕРИКИ В ФУНКЦИЯХ
/*
	что делать, если мы хотим написать функцию, складывающую элементы в слайсе, при этом
	она должна уметь работать с любым float и int? Использование дженериков поможет решить
	эту проблему.
*/

func showSum() {
	floats := []float64{1.0, 2.0, 3.0}
	ints := []int64{1, 2, 3}
	//если функция принимаетд дженерик, то НЕЛЬЗЯ передать в этом параметр interface{} (any).
	//wrongFloats := []interface{}{"string", struct{}{}, true}

	fmt.Println(sum(floats))
	fmt.Println(sum[int64](ints))
	//fmt.Println(sum(wrongFloats))
}

/*
Для функции дженерики пишутся через []. Обозначается имя какого-то типа (создается какой-то
тип), который доступен только в этой функции. В данном случае это тип либо int64, либо float64,
и функция может работать с обоими случаями (через union оператор).
*/
func sum[V int64 | float64](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// -----------------------------------------------------------------------------------------
// ВСТРОЕННЫЕ ИНТЕРФЕЙСЫ, КОТОРЫЕ МОЖНО ИСПОЛЬЗОВАТЬ В ДЖЕНЕРИКАХ
// 1) comparable
func showContains() {
	type Person struct {
		name     string
		age      int64
		jobTitle string
	}

	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("int:", contains(ints, 4)) //true

	strings := []string{"Vasya", "Dima", "Katya"}
	fmt.Println("strings:", contains(strings, "Katya")) //true

	people := []Person{
		{
			name:     "Vasya",
			age:      20,
			jobTitle: "Programmer",
		},
		{
			name:     "Dasha",
			age:      23,
			jobTitle: "Designer",
		},
		{
			name:     "Pasha",
			age:      30,
			jobTitle: "Admin",
		},
	}

	fmt.Println("structs:", contains(people, Person{
		name:     "Vasya",
		age:      21,
		jobTitle: "Programmer",
	})) //false
}

/*
	Функция, которая должна найти значение в любой структуре данных и вернуть bool в зависимости
	от того, нашла ли она значение.
	comparable - это тип, который означает что значение этого типа можно сравнивать (в go таковыми
	являются все типы, кроме map/slice/structure(у которых в качестве полей используются map/slice).
	Все остальные имплементируют интерфейс comparable).
*/

func contains[T comparable](elements []T, searchEl T) bool {
	//ищем элемент в структуре переданных нам данных
	for _, el := range elements {
		if searchEl == el {
			return true
		}
	}

	return false
}

//2) any

func showAny() {
	show(1, 2, 3)
	show("test1", "test2", "test3")
	show([]int64{1, 2, 3}, []int64{4, 5, 6})
	show(map[string]int64{
		"first":  1,
		"second": 2,
	})
	show(interface{}(1), interface{}("string"), any(struct{ name string }{name: "Vasya"}))
}

/*
	any - то же самое, что и interface{} (то есть любое значение). Может пригодиться в специфических
	случаях.
*/

func show[T any](entities ...T) {
	fmt.Println(entities)
}

// -----------------------------------------------------------------------------------------
// КАК ИСПОЛЬЗОВАТЬ ИНТЕРФЕЙСЫ В ДЖЕНЕРИКАХ
// ДЖЕНЕРИК КАК ТИП

type Number interface {
	~int64 | float64 //приближение типа - то есть пример даже кастомный тип, основанный на int64.
}

/*
	Дженерик можно создать отдельно от функции - через ключевое слово type.
*/

type Numbers[T Number] []T

func unionInterfaceAndType() {
	var ints Numbers[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...)

	floats := Numbers[float64]{1.0, 2, 5, 3, 5}

	fmt.Println(sumUnionInterface(ints))
	fmt.Println(sumUnionInterface(floats))
}

/*
	Можно использовать интерфейс в дженерике, но при этом у этого интерфейса НЕ должно быть
	НИКАКИХ методов. Иначе нельзя.
*/

func sumUnionInterface[V Number](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// -----------------------------------------------------------------------------------------
// ПРИБЛИЖЕНИЕ ТИПА (type approximation)

type CustomInt int64

// метод нужен, чтобы показать, что кастомный тип чем-то расширен
func (ci CustomInt) IsPositive() bool {
	return ci > 0
}

func typeApproximation() {
	customInts := []CustomInt{1, 2, 3, 5, 6}

	/*
		так передавать не стоит, так как CustomInt является кастомным типом, и из-за этого будет ошибка.
		НО если в типе Number указать ~int64, то можно - это и есть приближение типа. Это говорит о том,
		что мы можем использовать не только int64, но и те типы, что основаны на типе int64.
	*/
	fmt.Println(sumUnionInterface(customInts))
}

// -----------------------------------------------------------------------------------------

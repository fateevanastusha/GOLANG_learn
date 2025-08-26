package main

import "fmt"

func main() {
	/*
		Указатели - это тип данных, которые в качестве значения хранят адрес ячейки
		памяти, где и хранится нужное значение. Дефолнтое значение - nil.

		1) обозначение *Type
		Указатели не существуют отдельно, они могут быть только в связке с другими
		типами: *int, *string, *bool

		2) какие адреса могут хранить?
		- адреса конкретных значений (7, "nastya")
		- адрес другого указателя (указатель на указатель)
	*/

	//объявить указатель
	var intPointer *int                             //default value = nil
	fmt.Printf("%T %#v \n", intPointer, intPointer) //*int (*int)(nil) - указывает на указатель int, значение nil
	// fmt.Printf("%T %#v %#v\n", intPointer, intPointer, *intPointer) //panic! invalid memory address or nil pointer dereference

	//получение not-nil указателей
	//1)
	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)                             //int64 7
	var pointerA *int64 = &a                                  //получить указатель на &a
	fmt.Printf("%T %#v %#v\n", pointerA, pointerA, *pointerA) //*int64 (*int64)(0x1400000e0c8) 7s

	//получение по значению (чтобы из указателя получить значение (по адресу), нужно поставить * - 0x14000098028 => 7)
	fmt.Printf("%T %#v \n", *pointerA, *pointerA) //int64 7

	//2)
	var newPointer = new(float64)                                   //будет ссылаться на default значение указанного типа, в данном случае = 0
	fmt.Printf("%T %#v %#v\n", newPointer, newPointer, *newPointer) //*float64 (*float64)(0x14000090040) 0
	*newPointer = 3                                                 //так как тип у этой переменной - указатель на float64, то нужно его разыменовать
	fmt.Printf("%T %#v %#v\n", newPointer, newPointer, *newPointer) ////*float64 (*float64)(0x14000090040) 3

	/*
		3) где используют указатели?
	*/

	//1) side-effect
	/*
		в функции передаются КОПИИ, а не сама переменная в аргументы, поэтому надо
		передавать УКАЗАТЕЛЬ, чтобы изменить значение
	*/
	b := 15
	f1 := func(v int) {
		v += 10
	}
	f1(b)
	fmt.Println("This is B before =", b) //15 - значение не изменилось!
	f2 := func(v *int) {
		*v += 10
	}
	f2(&b)                              //& - получить указатель на переменную
	fmt.Println("This is B after =", b) //25 - значение изменилось!

	//2) признак пустого значения (позволит отличить дефолтное значение типа неопределенной переменной от отсутствия значения)

	hasWallet := func(money *int) bool {
		return money != nil
	}

	var wallet1 *int                                  // defalut value = nil
	fmt.Println("HAS WALLET1 :", hasWallet(wallet1))  // FALSE
	var wallet2 int                                   // default value of int = 0
	fmt.Println("HAS WALLET2 :", hasWallet(&wallet2)) // TRUE
	wallet3 := 100
	fmt.Println("HAS WALLET3 :", hasWallet(&wallet3)) // TRUE

	//3) экономия памяти (чтобы не копировать значения)
	/*
		когда у нас есть огромные данные, которые нам нужно куда-то передать, то,
		чтобы их не копировать - можно передавать указатель!
	*/
}

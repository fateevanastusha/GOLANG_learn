package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	/*
		GO
		- НЕ ПОДДЕРЖИВАЕТ неявное преобразование
		- статически типизированный язык (типы данных в Go (и других статически типизированных языках)
		связаны с переменной, а не с ее значением)
	*/

	fmt.Println("\n-----------------------------------------------------------------------\n")

	/*
		1) приведение целочисленных типов
	*/

	//с бОльшим количеством бит
	var a int8
	var b int32
	var c uint64
	a = 10
	b = int32(a)
	c = uint64(b)
	fmt.Println("This is C = ", c)

	//с меньшим количеством биты

	var d int64
	var e int8
	d = 15
	e = int8(d)
	fmt.Println("This is E = ", e)

	var f int64
	var g int8
	f = 129
	g = int8(f) //-127, произойдет потеря данных!
	fmt.Println("This is G = ", g)

	//получить максимальные значения для различных целочисленных типов:
	fmt.Println("This is max value for type Int8 = ", math.MaxInt8)     // 127
	fmt.Println("This is max value for type Uint8 = ", math.MaxUint8)   // 255
	fmt.Println("This is max value for type Int16 = ", math.MaxInt16)   // 32767
	fmt.Println("This is max value for type Uint16 = ", math.MaxUint16) // 65535

	fmt.Println("\n-----------------------------------------------------------------------\n")

	/*
		2) приведение целых чисел и чисел с плавающей точкой
	*/

	//преобразование целого числа в число с плавающей точкой
	var h int64
	var i float64
	h = 57
	i = float64(h)

	fmt.Println("This is I = ", i)

	//преобразование чисел с плавающей точкой в целые числа
	var j float64
	var k int
	j = 56.7742
	k = int(j)

	fmt.Println("This is K = ", k) //56 - теряем точность (после запятой)

	//числа, конвертируемые с помощью деления
	l := 5 / 2
	fmt.Println("This is L = ", l) //2

	m := 5.0 / 2
	fmt.Println("This is M = ", m) //2.5

	fmt.Println("\n-----------------------------------------------------------------------\n")

	/*
		3) конвертация строк в байты/rune и обратно
		строка в Go это срез байтов, поэтому мы можем конвертировать байты в строку и наоборот.
	*/

	//побайтовый срез
	var n string
	var o []byte
	var p string
	n = "str"
	o = []byte(n)
	p = string(o)

	fmt.Println("This is N = ", n) //str
	fmt.Println("This is O = ", o) //[115 116 114] - побайтовый срез
	fmt.Println("This is P = ", p) //str

	//руны
	var q string
	var r []rune
	var s string
	q = "str"
	r = []rune(q)
	s = string(r)

	fmt.Println("This is Q = ", q) //str
	fmt.Println("This is R = ", r) //[115 116 114] - срез рун
	fmt.Println("This is S = ", s) //str

	fmt.Println("\n-----------------------------------------------------------------------\n")

	/*
		4) конвертация в строки
	*/

	//конвертация целых чисел в строки

	/*
		Golang - язык со статической и строгой типизацией. Он не позволит вам сложить строку и число.
		user := "ученик"
		steps := 4

		fmt.Println("Поздравляю, " + user + "! Ты прошел " + steps + " шага по приведению типов.")

		Мы получим ошибку во время компиляции:
		invalid operation: ("Поздравляю, " + user + "! Ты прошел ") + steps (mismatched types string
		and int)
	*/

	var t int
	var u string
	t = 2000
	u = strconv.Itoa(t)

	fmt.Println("This is U = ", u) //"2000"

	//конвертация чисел с плавающей запятой в строку
	var v float64
	var w string
	var x string
	v = 1.00003426763444322
	// 1 параметр - число для конвертации
	// fmt - форматирование
	// prec - кол-во знаков после запятой (если мы хотим учесть все цифры после запятой, то можем в
	// prec передать -1)
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	w = strconv.FormatFloat(v, 'f', 5, 64)
	x = strconv.FormatFloat(v, 'f', -1, 64)

	fmt.Println("This is W = ", w) //"1.00003"
	fmt.Println("This is X = ", x) //"1.0000342676344431"

	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent),
	// or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).

	var y float64
	var z string
	// Использовать fmt для конвертации нежелательно из-за того что производительность ниже по
	// сравнению с strconv.
	y = 563.3456772334457580958892389489
	z = fmt.Sprintf("%f", y)

	fmt.Println("This is Z = ", z) //"563.345677"

	fmt.Println("\n-----------------------------------------------------------------------\n")

	/*
		5) конвертация bool в string
	*/

	var aa bool
	var ab string
	aa = true
	ab = strconv.FormatBool(aa)

	fmt.Println("This is AA = ", ab) //"true"

	fmt.Println("\n-----------------------------------------------------------------------\n")

	/*
		6) конвертация строк в другие типы
	*/

	/*
		foo := "10"
		bar := "15"
		baz := foo - bar
		fmt.Println(baz)

		ОШИБКА: invalid operation: foo - bar (operator - not defined on string)
		Операнд вычитания не является действительным для строк
	*/

	var ac string = "10"
	var ad string = "15"

	// при конвертации строки, которая не содержит в себе число - ваша программа выдаст вам ошибку
	//strconv.Atoi: parsing "not a number": invalid syntax

	ae, err1 := strconv.Atoi(ac) //или ParseInt
	if err1 != nil {
		panic(err1)
	}
	af, err2 := strconv.Atoi(ad)
	if err2 != nil {
		panic(err2)
	}

	ag := af - ae

	fmt.Println("This is AG = ", ag) //5

	var ah string = "23.23456"
	ai, err3 := strconv.ParseFloat(ah, 64) //float64 вернет в любом случае
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("This is AI = ", ai) //"23.23456"

	aj := "1.0000000012345678"

	ak, _ := strconv.ParseFloat(aj, 32)
	al, _ := strconv.ParseFloat(aj, 64)

	fmt.Println("This is AK = ", ak) //"1" - не уместились и потеряли всю дробную часть!
	fmt.Println("This is AL = ", al) //"1.0000000012345678"

	fmt.Println("\n-----------------------------------------------------------------------\n")

}

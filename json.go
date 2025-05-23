package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	/*
		Marshal и Unmarshal (кодирование и декодирование) данных в формате JSON в стандартной
		библиотеке Go реализовано в пакете encoding/json.
	*/

	//1) кодирование

	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	s1 := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
	// и возвращает байтовый срез с данными, кодированными в формат JSON.
	data1, _ := json.Marshal(s1)
	fmt.Printf("%s\n", data1) // {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}

	//чтобы получить человекочитаемый формат (с отступами)
	data2, _ := json.MarshalIndent(s1, "", "\t")
	fmt.Printf("%s\n", data2)

	fmt.Println("----------------------------------------------------------------------------------------------")

	//2) декодирование

	data3 := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)
	var s2 myStruct
	json.Unmarshal(data3, &s2)
	fmt.Printf("%v\n", s2)

	fmt.Println("----------------------------------------------------------------------------------------------")

	//3) чек json на валидность

	type user struct {
		Name     string
		Email    string
		Status   bool
		Language []byte
	}

	m := user{Name: "John Connor", Email: "test email"}

	data, _ := json.Marshal(m)

	data = bytes.Trim(data, "{") // испортим json удалив '{'

	// функция json.Valid возвращает bool, true - если json правильный
	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}

	fmt.Printf("%s\n", data) // вывод: "Name":"John Connor","Email":"test email","Status":false,"Language":null}

	fmt.Println("----------------------------------------------------------------------------------------------")

	//4) поля, начинающиеся с маленькой буквы в кодировании/декодировании не участвуют! (их называют неэкспортируемыми)

	type test struct {
		Name  string
		email string
	}

	t := test{Name: "John Connor", email: "test email"}
	data4, _ := json.MarshalIndent(t, "", "\t")
	fmt.Printf("%s\n", data4) //{"Name": "John Connor"}

	data5 := []byte(`{"Name": "John Connor", "email": "test email"}`)
	var s3 test
	json.Unmarshal(data5, &s3)
	fmt.Printf("%v\n", s3) //{John Connor }

	fmt.Println("----------------------------------------------------------------------------------------------")

	//5) аннотирование структур
	/*
		если мы не хотим, чтобы какие-то поля принимали участие в кодировании/декодировании, или делали это под
		другим названием, можно использовать теги для полей структуры
	*/

	// в общем виде аннотация выглядит так: `json:"используемое_имя,опция,опция"`

	type Person struct {
		// при кодировании / декодировании будет использовано имя name, а не Name
		Name string `json:"name"`

		// при кодировании / декодировании будет использовано то же имя (Age),
		// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
		// то при кодировании оно будет опущено
		Age int `json:",omitempty"`

		// при кодировании / декодировании поле всегда игнорируется
		Status bool `json:"-"`
	}

	p := Person{Name: "John Connor", Age: 0, Status: true}

	data6, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", data6) // {"name":"John Connor"}

	fmt.Println("----------------------------------------------------------------------------------------------")

	//6) Encoder и Decoder
	/*
		Данные типы помимо методов Encode() и Decode() предоставляют нам ряд дополнительных методов, которые могут быть
		использованы в определенных случаях.
	*/

	// Определяем структуру с JSON-тегами
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var (
		src = testStruct{Name: "John Connor", Age: 35} // исходная структура с данными
		dst = testStruct{}                             // структура-приёмник (будет заполнена позже)
		buf = new(bytes.Buffer)                        // буфер для промежуточного хранения JSON
	)

	// Создаём JSON-энкодер, который будет писать в буфер
	enc := json.NewEncoder(buf)

	// Создаём JSON-декодер, который будет читать из того же буфера
	dec := json.NewDecoder(buf)

	// Кодируем структуру src в JSON и записываем в буфер
	enc.Encode(src)

	// Декодируем JSON из буфера обратно в структуру dst
	dec.Decode(&dst)

	// Выводим результат — структура dst теперь содержит те же данные, что и src
	fmt.Print(dst) // {John Connor 35}

}

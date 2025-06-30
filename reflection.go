package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

/*
	Пакет reflection дает инструментарий для любых взаимодействий с любой структурой. Основная фича reflection
	заключается в способности работать с типами и значениями в режиме выполнения. Основные концепции, которые
	есть в reflect - это:
	1) Type - представляет описание типа в Go. С помощью reflect.Type можноузнать о характеристиках типа,
	таких как его имя, размер, количество полей (если это структура), и т.п.
	2) Value - представляет собой значение переменной. С помощью reflect.Value можно получить и изменять
	данные, хранящиеся в переменной.

	Функции из пакета reflect:
	1) reflect.TypeOf(): возвращает reflect.Type, представляющий тип переменной.
	2) reflect.ValueOf(): возвращает reflect.Value, представляющий значение переменной.
	3) Interface(): возвращает interface{}, представляющий текущее значение. Это позволяет извлекать оригинальные
	данные из reflect.Value.
	4) Kind(): возвращает reflect.Kind, представляющий конкретный тип данных, например, Int, Float64, Struct, и т.д.
	5) NumField() и Field(i int): используются для работы со структурами, позволяют получить количество полей и доступ
	к каждому полю соответственно.
	6) NumMethod() и Method(i int): используются для работы с методами, предоставляя доступ к методу и его характеристикам.

*/

// -----------------------------------------------------------------------------------------
// пример использования:

func main() {
	fmt.Println("-----------------------------------------------------------------------------------------")
	f_1()
	fmt.Println("-----------------------------------------------------------------------------------------")
	f_2()
	fmt.Println("-----------------------------------------------------------------------------------------")
	f_3()
	fmt.Println("-----------------------------------------------------------------------------------------")
	f_4()
	fmt.Println("-----------------------------------------------------------------------------------------")
	f_5()
	fmt.Println("-----------------------------------------------------------------------------------------")
	f_6()
	fmt.Println("-----------------------------------------------------------------------------------------")
}

type Person struct {
	Name string
	Age  int
}

func f_1() {
	fmt.Println("F1\n")
	p := Person{Name: "Alice", Age: 30}

	// получаем тип переменной
	t := reflect.TypeOf(p)
	fmt.Println("Тип:", t.Name()) // вывод: Тип: Person

	// получаем значение переменной
	v := reflect.ValueOf(p)
	fmt.Println("Значение:", v) // вывод: Значение: {Alice 30}

	// получаем количество полей
	numFields := v.NumField()
	fmt.Println("Количество полей:", numFields) // вывод: Количество полей: 2

	// итерация по полям структуры
	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		fmt.Printf("Поле %d: %v\n", i, field)
		// вывод: поле 0: Alice
		//        поле 1: 30
	}
}

// -----------------------------------------------------------------------------------------
//Автоматическая сериализация и десериализация структур

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func f_2() {
	fmt.Println("F2\n")
	user := User{Name: "John Doe", Email: "john@example.com", Age: 30}
	jsonData, err := MarshalStructToJSON(user)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}

	fmt.Println("JSON:", string(jsonData))
}

// MarshalStructToJSON принимает любую структуру и возвращает ее JSON-представление
func MarshalStructToJSON(s interface{}) ([]byte, error) {
	// Получаем значение из интерфейса
	v := reflect.ValueOf(s)

	// Проверяем, что переданный параметр - структура
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct but got %s", v.Kind())
	}

	// Используем стандартную библиотеку для сериализации в JSON
	return json.Marshal(s)
}

// -----------------------------------------------------------------------------------------
//Валидатор структур с использованием тэгов

type Product struct {
	Name  string  `validate:"required"`
	Price float64 `validate:"min=0"`
}

func f_3() {
	fmt.Println("F3\n")

	product := Product{Name: "", Price: -15.0}

	err := ValidateStruct(product)
	if err != nil {
		fmt.Println("Ошибка валидации:", err)
	} else {
		fmt.Println("Структура валидна")
	}

}

// ValidateStruct принимает структуру и проверяет её поля на соответствие тэгам валидации
func ValidateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		return errors.New("expected struct")
	}

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("validate")

		if strings.Contains(tag, "required") && fieldValue.IsZero() {
			return fmt.Errorf("поле %s обязательно", fieldType.Name)
		}

		if strings.Contains(tag, "min=0") && fieldValue.Kind() == reflect.Float64 && fieldValue.Float() < 0 {
			return fmt.Errorf("значение поля %s должно быть неотрицательным", fieldType.Name)
		}
	}

	return nil
}

// -----------------------------------------------------------------------------------------
// Динамическое создание экземпляров типов

type Order struct {
	ID    int
	Total float64
}

func f_4() {
	fmt.Println("F4\n")
	orderType := reflect.TypeOf(Order{})
	order := CreateInstanceOfType(orderType).(Order)

	fmt.Printf("Созданный экземпляр: %+v\n", order)
}

func CreateInstanceOfType(t reflect.Type) interface{} {
	return reflect.New(t).Elem().Interface()
}

// -----------------------------------------------------------------------------------------
// Вызов методов с помощью reflection

type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

func f_5() {
	fmt.Println("F5\n")
	calc := Calculator{}

	sum, err := CallMethod(calc, "Add", 5, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Сумма:", sum)

	product, err := CallMethod(calc, "Multiply", 5, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Произведение:", product)
}

func CallMethod(obj interface{}, methodName string, args ...interface{}) (interface{}, error) {
	v := reflect.ValueOf(obj)
	method := v.MethodByName(methodName)

	if !method.IsValid() {
		return nil, fmt.Errorf("метод %s не найден", methodName)
	}

	methodArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		methodArgs[i] = reflect.ValueOf(arg)
	}

	result := method.Call(methodArgs)
	if len(result) > 0 {
		return result[0].Interface(), nil
	}

	return nil, nil
}

// -----------------------------------------------------------------------------------------
// Инспекция и манипуляция с полями структур

type Employee struct {
	Name   string
	Salary float64
}

func f_6() {
	fmt.Println("F6\n")
	emp := Employee{Name: "Alice", Salary: 50000.0}
	fmt.Println("До изменения:", emp)

	err := SetFieldValue(&emp, "Salary", 60000.0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("После изменения:", emp)
}

func SetFieldValue(obj interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(obj).Elem()
	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		return fmt.Errorf("поле %s не найдено", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("поле %s не может быть изменено", fieldName)
	}

	field.Set(reflect.ValueOf(value))
	return nil
}

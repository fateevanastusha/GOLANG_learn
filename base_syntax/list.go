package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*
		Двусвязные списки (двунаправленные) - структура данных в виде последовательности, где каждый элемент хранит
		собственное значение, а так же ссылку на предыдущий и на следующий элемент.

	*/

	//1) создать

	mylist := list.New()

	//2) добавить элемент
	mylist.PushBack(1) //добавить в конец
	mylist.PushBack(2)
	mylist.PushFront(3) //добавить в начало
	mylist.PushFront(4)
	//4312

	//3) итерироваться
	for t := mylist.Front(); t != nil; t = t.Next() {
		fmt.Println("Iter list, value - ", t.Value)
	}
	//4 3 1 2

	//4) получить значение
	fmt.Println("Back value - ", mylist.Back().Value)   //2
	fmt.Println("Front value - ", mylist.Front().Value) //4

	//5) удалить элемент
	/*
		удалить можно по указателю на элемент
	*/
	mylist.Remove(mylist.Front()) //останется 312
	mylist.Remove(mylist.Back())  //останется 31

	mylist.PushBack(5)
	mylist.PushBack(6)
	mylist.PushBack(7)
	mylist.PushBack(8)

	//до - 315678

	for t := mylist.Front(); t != nil; t = t.Next() {
		if t.Value == 5 {
			mylist.Remove(t)
		}
	}

	for t := mylist.Front(); t != nil; t = t.Next() {
		fmt.Println("Iter list after deleting, value - ", t.Value)
	}
	//после - 31678

}

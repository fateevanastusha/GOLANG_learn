package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	/*
		os
		- содержит огромный набор высокоуровневых инструментов для работы с файлами (и не только).
	*/
	// создаем файл
	os.Create("text.txt")
	// переименовываем файл
	os.Rename("text.txt", "new_text.txt")
	// удаляем файл
	os.Remove("new_text.txt")
	// кстати, os позволяет работать не только с файлами
	// выходим из программы:
	os.Exit(0)

	file1, _ := os.Create("text.txt")
	file2, _ := os.Create("text.txt")
	info1, _ := file1.Stat() // функция Stat возвращает информацию о файле и ошибку
	info2, _ := file2.Stat()
	fmt.Println(os.SameFile(info1, info2)) // true

	// вот что мы можем получить из FileInfo:
	// A FileInfo describes a file and is returned by Stat and Lstat.
	type FileInfo interface {
		Name() string       // base name of the file
		Size() int64        // length in bytes for regular files; system-dependent for others
		Mode() FileMode     // file mode bits
		ModTime() time.Time // modification time
		IsDir() bool        // abbreviation for Mode().IsDir()
		Sys() interface{}   // underlying data source (can return nil)
	}

	file3, _ := os.Create("text.txt")
	file3.WriteString("1 строка \n")
	file3.WriteString("2 строка \n")
	file3.Close()

	// внутри файла будет:
	// 1 строка
	// 2 строка
}

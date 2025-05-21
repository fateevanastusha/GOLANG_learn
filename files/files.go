package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	dataForFile := []byte("Тестовая строка, предназначенная для записи в файл")

	// Создаем новый файл и записываем в него данные dataForFile (если такой файл уже есть - он его перезапишет)
	if err := ioutil.WriteFile("test.txt", dataForFile, 0600); err != nil {
		panic(err)
	}

	// Читаем данные из того же файла
	dataFromFile, err := ioutil.ReadFile("test.txt")
	if err != nil {

	}

	// Сравниваем исходные данные с записанными в файл и прочитанными из него
	fmt.Printf("dataForFile == dataFromFile: %v\n", bytes.Equal(dataFromFile, dataForFile))

	// Изучаем содержимое директории
	filesFromDir, err := ioutil.ReadDir(".")
	if err != nil {

	}

	for _, file := range filesFromDir {
		// Проходим по всем найденным файлам и печатаем их имя и размер
		fmt.Printf("name: %s, size: %d, mode: %d, isdir: %t", file.Name(), file.Size(), file.Mode(), file.IsDir())
	}

	// Output:
	// dataForFile == dataFromFile: true
	// name: main.go, size: 727
	// name: test.txt, size: 93
}

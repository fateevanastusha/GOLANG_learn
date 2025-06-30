package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	/*
		io.Reader/io.Writer
	*/
	simpleReader() // простое чтение
	fmt.Println("-----------------------------------------------------------------------")
	rowsReader() // чтение по строке
	fmt.Println("-----------------------------------------------------------------------")
	simpleWriter() // простая запись
	fmt.Println("-----------------------------------------------------------------------")
	rowsWriter() // запись по строке
	fmt.Println("-----------------------------------------------------------------------")
	osFile()

}

func simpleReader() {
	buf := make([]byte, 10)

	reader := NumsReader{nums: "1,2,3,4,5,6.7=,g8"}

	count, err := reader.Read(buf)
	if err != nil && err != io.EOF { //обязательно нужно проверять end of file
		log.Fatal(err)
	}
	fmt.Println(string(buf), count)

}

type NumsReader struct {
	nums string
}

// ТУТ ПОЛУЧАЕМ "ФЛЕШКУ", НА КОТОРУЮ ЗАПИСЫВАЕМ ДАННЫЕ С ДИСКА (r.nums)
// берет буфер, в который нужно класть прочитанные данные
// возвращает количество прочитанных символов и ошибку, если она есть
func (r NumsReader) Read(p []byte) (n int, err error) {
	var cnt int
	for i := 0; i < len(r.nums); i++ {
		if r.nums[i] >= '0' && r.nums[i] <= '9' {
			p[cnt] = r.nums[i]
			cnt++
		}
	}

	return cnt, io.EOF
}

//------------------------------------------------------------------------------------------------------------------

func rowsReader() {
	rowsReader := RowsReader{text: "first_row\nsecond_row\nthird_row"}
	var (
		err   error
		count int
	)
	row := make([]byte, 100)

	for err != io.EOF { //читаем все строки
		count, err = rowsReader.Read(row)
		fmt.Println(string(row[0:count]), count)
	}
}

type RowsReader struct {
	text string
}

func (r *RowsReader) Read(p []byte /* пишем данные сюда */) (n int, err error) {
	var i int
	for i = 0; i < len(r.text); i++ {
		/*
			как только находим перенос строки - обрезаем строку до этого переноса и заканчиваем работу
		*/
		if r.text[i] == '\n' {
			r.text = r.text[i+1:]
			break
		}
		p[i] = r.text[i]

		/*
			если мы уже прошлись по всей строке - то просто возвращаем ее, а читаемый текст делаем пустой строкой
		*/
		if i == len(r.text)-1 {
			r.text = ""
			return i + 1, io.EOF
		}
	}
	return i + 1, nil
}

//------------------------------------------------------------------------------------------------------------------

func simpleWriter() {
	nums := []byte{'1', ',', '2', '=', '3', '.', '9'}

	writer := NumsWriter{storedNums: make([]byte, 10)}
	count, err := writer.Write(nums)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println(string(writer.storedNums), count)
}

type NumsWriter struct {
	storedNums []byte
}

// ПЕРЕДАЕМ "ФЛЕШКУ", ЧТОБЫ ПЕРЕПИСАТЬ С НЕЕ ДАННЫЕ НА ДИСК
// берет буфер, из котрого нужно брать данные для записи, записывает их в w
// возвращает количество записанных символов и ошибку, если она есть
func (w NumsWriter) Write(p []byte) (n int, err error) {
	var cnt int
	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			w.storedNums[cnt] = p[i]
			cnt++
		}
	}

	return cnt, io.EOF
}

//------------------------------------------------------------------------------------------------------------------

func rowsWriter() {
	text := "first_row\nsecond_row\nthird_row"
	var (
		err   error
		count int
	)
	rowsWriter := RowsWriter{storedRow: make([]byte, 100)}

	for err != io.EOF {
		count, err = rowsWriter.Write(&text)
		fmt.Println(string(rowsWriter.storedRow[0:count]), count)
	}

}

type RowsWriter struct {
	storedRow []byte
}

func (w *RowsWriter) Write(text *string /* читаем данные отсюда */) (n int, err error) {
	var i int
	for i = 0; i < len(*text); i++ {
		if (*text)[i] == '\n' {
			*text = (*text)[i+1:]
			w.storedRow = w.storedRow[0 : i+1]
			break
		}
		w.storedRow[i] = (*text)[i]

		if i == len(*text)-1 {
			*text = ""
			return i + 1, io.EOF
		}
	}

	return i + 1, nil
}

//------------------------------------------------------------------------------------------------------------------

func osFile() {
	newFile, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}

	n, err := newFile.WriteString("hello world")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n) //11

	//Close
	/*
		Когда закончили работу с файлом - нужно обязательно его закрыть
		1) мы больше не сможем случайно с ним что-то сделать
		2) не засорять память, в некоторых ос ограничение на количество открытых файлов
	*/
	err = newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	//Open
	/*
		может не получиться записать в открытый файл, так как пермишенны подразумевают только
		чтение (readonly)
	*/
	file, err := os.Open("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 100)
	n, err = file.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf), n) //hello world 11
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	//OpenFile
	/*
		необходимо указать пермишенны
	*/
	file, err = os.OpenFile("new.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	n, err = file.WriteString("\nhello world2")
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	//есть сокращения, позволяющие не писать дополнительно open/close
	//ReadFile
	res, err := os.ReadFile("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("START - ", string(res), " - END")
	//WriteFile
	err = os.WriteFile("new.txt", []byte("helloworld1\nhelloworld2\nhelloworld3"), 0666)
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	createFile()
	fmt.Println("-----------------------------------------------------------------------")
	// writeTo()
	fmt.Println("-----------------------------------------------------------------------")
	// readFrom()
	fmt.Println("-----------------------------------------------------------------------")
	// ioCopy()
	fmt.Println("-----------------------------------------------------------------------")
	bufioWrite()
	fmt.Println("-----------------------------------------------------------------------")
	// createFileBufio()

}

func createFile() {
	fmt.Println("FUNC CREATEFILE\n")

	start := time.Now()

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Create("test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var i int
	for i < 100 {
		/*
			в данном случае мы занимаем меньше памяти, но время выполнения увеличится
		*/
		if _, err = file.WriteString(fmt.Sprintf("row number %d\n", i+1)); err != nil { //вызывает 100 syscall'ов
			log.Fatal(err)
		}
		i++
	}

	fmt.Println("end time -", time.Since(start))
}
func writeTo() {
	fmt.Println("FUNC WRITETO\n")

	//io.WriteTo (чтение в )
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	//передаем сюда что-то, что реализует интерфейс io.WriteTo
	//переписываем данные из файла в консоль
	count, err := file.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("writeTo count -", count) //количество записанных байтов

}
func readFrom() {
	fmt.Println("FUNC READFROM\n")

	file2, err := os.OpenFile("test2.txt", os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file2.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count, err := file2.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("readFrom count -", count) //количество записанных байтов

}
func ioCopy() {
	fmt.Println("FUNC IOCOPY\n")

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// count, err := io.Copy(os.Stdout, file)
	count, err := io.CopyBuffer(os.Stdout, file, make([]byte, 1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ioCopy count -", count) //количество записанных байтов

}
func bufioWrite() {
	fmt.Println("FUNC BUFIOWRITE\n")

	/*
		ситуация, нам нужно чтобы записалось либо все, либо ничего. тут нам поможет writer.
		принцип: все, что нужно записать, кладется в буфер (через WriteString), затем, если
		что-то пошло не так - то мы просто их не аппендим. если все ок - записываем данные
		из буфера и очищаем его.
	*/

	var err error

	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		if err == nil {
			/*
				если ошибки нет - то мы флашим (что значит записать в указанный ранее os.Stdout),
				если ошибка есть - то не флашим.
			*/
			if err := writer.Flush(); err != nil {
				log.Fatal(err)
			}
		}
	}()

	// do some logic 1
	// write log 1
	if _, err := writer.WriteString("first action done\n"); err != nil {
		log.Fatal(err)
	}

	// do some logic 2
	// write log 2
	if _, err := writer.WriteString("second action done\n"); err != nil {
		log.Fatal(err)
	}

	// err = errors.New("some error")

}
func createFileBufio() {
	fmt.Println("FUNC CREATEFILEBUFIO\n")

	start := time.Now()

	file, err := os.Create("test3.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	writer := bufio.NewWriter(file) //сущность, чтобы писать данные в переданный итем
	defer func() {
		//очистить буфер
		if err = writer.Flush(); err != nil {
			log.Fatal(err)
		}
	}()

	var i int

	for i < 100 {
		/*
			пишет данные сначала в буфер, пока он не заполнится, а только потом вызывает
			syscall, там уже и пишет данные (то есть в итоге мы жертвуем оперативной памятью)
			получается, в данном случае мы занимаем больше памяти, но сокращаем время выполнения
		*/
		if _, err = writer.WriteString(fmt.Sprintf("%d\n", i)); err != nil { // check syscalls
			log.Fatal(err)
		}

		i++
	}

	fmt.Println("end time -", time.Since(start))
}

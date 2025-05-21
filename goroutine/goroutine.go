// package main

// import (
// 	"fmt"
// 	"time"
// )

// func downloadFile(filename string, done chan bool) {
// 	fmt.Printf("Starting download: %s\n", filename)
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("Finished download: %s\n", filename)

// 	done <- true // отправляем сигнал о завершении
// }

// func main() {
// 	fmt.Println("Starting downloads...")

// 	startTime := time.Now()

// 	// создаем канал для отслеживания статуса горутин
// 	done := make(chan bool)

// 	go downloadFile("file1.txt", done)
// 	go downloadFile("file2.txt", done)
// 	go downloadFile("file3.txt", done)

// 	// Ждем пока все горутины сигнализируют о закрытии
// 	for i := 0; i < 2; i++ {
// 		<-done // Получаем сигнал от каждой завершенной горутины
// 	}

// 	elapsedTime := time.Since(startTime)
// 	fmt.Printf("All downloads completed! Time elapsed: %s\n", elapsedTime)
// }

package main

import (
	"fmt"
	"time"
)

func sender(ch chan string, done chan bool) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("message %d", i)
		time.Sleep(100 * time.Millisecond)
	}
	ch <- "nastya"
	close(ch) // Закрываем канал после завершения отправки
	done <- true
}

func receiver(ch chan string, done chan bool) {
	// range позволяет читать канал пока он не закроется(главное его закрыть)
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
	done <- true
}

func main() {
	ch := make(chan string)
	senderDone := make(chan bool)
	receiverDone := make(chan bool)

	go sender(ch, senderDone)
	go receiver(ch, receiverDone)

	// блокируемся до тех пор пока не отработают функции
	<-senderDone
	<-receiverDone

	fmt.Println("All operations completed!")
}

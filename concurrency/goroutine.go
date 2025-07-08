package main

import (
	"fmt"
	"sync"
)

func main() {

	//1) зачем нужны горутины?

	// downloadFile1 := func(filename string) {
	// 	fmt.Printf("Starting download: %s\n", filename)
	// 	// Simulate file download with sleep
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Printf("Finished download: %s\n", filename)
	// }

	// fmt.Println("Starting downloads...")

	// startTime := time.Now()

	// downloadFile1("file1.txt")
	// downloadFile1("file2.txt")
	// downloadFile1("file3.txt")

	// elapsedTime := time.Since(startTime)

	// fmt.Printf("All downloads completed! Time elapsed: %s\n", elapsedTime)

	// Starting downloads...
	// Starting download: filel.txt
	// Finished download: filel.txt
	// Starting download: file2.txt
	// Finished download: file2.txt
	// Starting download: file3.txt
	// Finished download: file3.txt
	// All downloads completed! Time elapsed: 6s
	// Program exited.

	/*

		0s        2s        4s        6s
		|---------|---------|---------|
		| file1   | file2   | file3   |
		| .txt    | .txt    | .txt    |
		| (2s)    | (2s)    | (2s)    |
		|---------|---------|---------|

		выполнение очень медленное, так как все задачи выполнялись последовательно и выполнились за t*n время (t - время
		выполнения одной задачи, n - количество задач). Можно это ускорить с помощью горутин, запустив их параллельно.
	*/

	fmt.Println("----------------------------------------------------------------------------------------------")

	// fmt.Println("Starting downloads...")

	// startTime := time.Now()

	// go downloadFile1("file1.txt")
	// go downloadFile1("file2.txt")
	// go downloadFile1("file3.txt")

	// elapsedTime := time.Since(startTime)

	// fmt.Printf("All downloads completed! Time elapsed: %s\n", elapsedTime)

	// Starting downloads...
	// All downloads completed!
	// Program exited.

	/*
		Вывелось не все, потому что main завершилась до того, как горутины отработали. Фукнция main - тоже горутина.
	*/

	fmt.Println("----------------------------------------------------------------------------------------------")

	//2) wait group
	/*
		для управления конкурентностью в го применяется `sync.WaitGroup`, благодаря wg можно ожидать завершение горутин.

		1) WaitGroup инициализирует внутренний счетчик
		2) wg.Add(n) увеличивает счетчик на n
		3) wg.Done() уменьшает счетчик на 1
		4) wg.Wait() блокирует main до тех пор пока счетчик не станет 0


		эти две строчки эквивалентны:
		var wg sync.WaitGroup
		wg := sync.WaitGroup{}

		так как когда в go мы объявляем переменную с каким-то типом, она становится равна default значению для этого типа
		("" для string и тд), поэтому в первом случае он инициализирует его нулевым состоянием, готовым к использованию.

	*/

	// downloadFile2 := func(filename string, wg *sync.WaitGroup) {
	// 	// Сообщим wg о том что мы закончили перед выходом из функции
	// 	defer wg.Done()

	// 	fmt.Printf("Starting download: %s\n", filename)
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Printf("Finished download: %s\n", filename)
	// }

	// fmt.Println("Starting downloads...")

	// var wg sync.WaitGroup

	// // Сообщим wg что мы собираемся запускать 3 горутины
	// wg.Add(3)

	// go downloadFile2("file1.txt", &wg)
	// go downloadFile2("file2.txt", &wg)
	// go downloadFile2("file3.txt", &wg)

	// // Ждем завершения всех горутин
	// wg.Wait()

	// fmt.Println("All downloads completed!")

	// Starting downloads...
	// Starting download: file3.txt
	// Starting download: filel.txt
	// Starting download: file2.txt
	// Finished download: file3.txt
	// Finished download: filel.txt
	// Finished download: file2.txt
	// All downloads completed!
	// Program exited.

	/*
		1) main вызывает add(3) до запуска горутин
		2) Каждая горутина вызывает Done() по завершению функции (defer wg.Done())
		3) main блокируется до тех пор пока счетчик wait() не станет 0
		4) Когда счетчик становится равным 0, блокировка main снимается и программа может завершаться
	*/

	/*
		// ТАК ДЕЛАТЬ НЕЛЬЗЯ
		go downloadFile("file1.txt", &wg)
		wg.Add(1)  // Нарушен порядок!

		// ТАК ДЕЛАТЬ НЕЛЬЗЯ
		wg.Add(2)  // Нарушен счетчик
		go downloadFile("file1.txt", &wg)
		go downloadFile("file2.txt", &wg)
		go downloadFile("file3.txt", &wg)

		// ТАК ДЕЛАТЬ НЕЛЬЗЯ
		func downloadFile(filename string, wg *sync.WaitGroup) {
			// не указан wg.Done()
			fmt.Printf("Downloading: %s\n", filename)
		}
	*/

	fmt.Println("----------------------------------------------------------------------------------------------")

	//3) как горутинам обмениваться информацией?
	/*
		горутины могут обмениваться информацией через каналы. Каналы это мощные примитивы взаимодействия между горутинами
		предоставляющий возможность безопасного обмена данными.

		вот некоторые свойства каналов

		1) Каналы это блокировки по своей сути
		2) Запись в канал ch <- value  блокирует main пока другая горутина не прочтет из канала.
		3) Чтение из канала <-ch блокирует main пока другая горутина не запишет из канал
	*/

	// downloadFile3 := func(filename string, done chan bool) {
	// 	fmt.Printf("Starting download: %s\n", filename)
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Printf("Finished download: %s\n", filename)

	// 	done <- true // отправляем сигнал о завершении
	// }

	// fmt.Println("Starting downloads...")

	// startTime := time.Now()

	// // создаем канал для отслеживания статуса горутин
	// done := make(chan bool)

	// go downloadFile3("file1.txt", done)
	// go downloadFile3("file2.txt", done)
	// go downloadFile3("file3.txt", done)

	// // Ждем пока все горутины сигнализируют о закрытии
	// for i := 0; i < 3; i++ {
	// 	<-done // Получаем сигнал от каждой завершенной горутины
	// }

	// elapsedTime := time.Since(startTime)
	// fmt.Printf("All downloads completed! Time elapsed: %s\n", elapsedTime)

	fmt.Println("----------------------------------------------------------------------------------------------")

	//4) mutex
	/*
		Что делать, если у нас есть несколько экземпляров горутины, влияющие на одну переменную?
	*/

	var x int
	wg1 := new(sync.WaitGroup)

	for i := 0; i < 1000; i++ {
		// Запускаем 1000 экземпляров горутины, увеличивающей счетчик на 1
		wg1.Add(1)
		go func(wg1 *sync.WaitGroup) {
			defer wg1.Done()
			x++
		}(wg1)
	}

	wg1.Wait()

	// По идее значение счетчика должно быть 1000, но крайне вероятно, что этого не произойдет
	fmt.Println(x) //случайной значение, < 1000

	/*
		Первая горутина получает значение переменной x, а вторая одновременно с этим делает то же самое. В результате
		обе горутины считают, что x = 0, и обе присваивают ему значение = 1 (0 + 1). Получается, что мы наткнулись на
		эффект "грязного чтения". Чтобы заблокировать это значение существуют мьютексы.

		В Go мьютексы (sync.Mutex) используются для синхронизации доступа к разделяемым данным, чтобы избежать гонки
		данных (data race) при выполнении конкурентного кода (например, с использованием горутин). В Go (и вообще в
		многопоточном программировании), мьютексы блокируют не доступ к переменной или функции напрямую, а управляют
		доступом к критической секции кода, в которой работает с этими переменными.

		Он блокирует код, который ты сам обернул в mu.Lock() и mu.Unlock(). Если другая горутина зашла в область кода,
		которая залочена - он будет ждать, пока она разлочится.

		Кстати, под капотом у каналов - мьютексы, так что мьютексы работают быстрее каналов в любом случае.
	*/

	var y int
	wg2 := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		// Запускаем 1000 экземпляров горутины, увеличивающей счетчик на 1
		wg2.Add(1)
		go func(wg2 *sync.WaitGroup, mu *sync.Mutex) {
			defer wg2.Done()
			mu.Lock() //заблокировали
			y++       // теперь они не смогут одновременно получить доступ к этой переменной
			mu.Unlock()
		}(wg2, mu)
	}

	wg2.Wait()
	fmt.Println(y) //1000

	//реализовать то же самое с использованием каналов
	var z int
	wg3 := new(sync.WaitGroup)
	channel := make(chan bool, 1)

	channel <- true

	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go func(wg3 *sync.WaitGroup) {
			defer wg3.Done()
			<-channel
			z++
			channel <- true
		}(wg3)
	}

	wg3.Wait()
	fmt.Println(z) //1000

	/*
		у нас забитый канал размером 1. При заходе в первую горутину он может прочитать значение, забирает его
		(параллельные другие горутины пока не могут прочитать, так как канал пустой, и ждут пока там что-то
		появится). Далее мы инкрементируем значение (или выполняем любой другой конкуретный код), затем кладем
		значение в канал снова. Любая другая горутина его ловит (одна) и все повторяется по кругу.
	*/

	fmt.Println("----------------------------------------------------------------------------------------------")

	//5) с циклами
	N := 10

	//неправильный вариант, когда берем значение, которое УЖЕ поменялось
	first := func() {
		m := make(map[int]int)

		wg := &sync.WaitGroup{}
		mu := &sync.Mutex{}
		wg.Add(N)

		for i := 0; i < N; i++ {
			go func() {
				defer wg.Done()
				mu.Lock()
				m[i] = i //берет ссылку на i, а не копию. К этому моменту i может уже измениться, поэтому многие i мы пропустим!
				mu.Unlock()
			}()
		}

		wg.Wait()
		fmt.Println(len(m)) //выведет < 10
	}

	//правильный вариант, берем копию значения, поэтому нам уже неважно, поменялось оно или нет
	second := func() {
		m := make(map[int]int)

		wg := &sync.WaitGroup{}
		mu := &sync.Mutex{}
		wg.Add(N)

		for i := 0; i < N; i++ {
			go func(i int) {
				defer wg.Done()
				mu.Lock()
				m[i] = i
				mu.Unlock()
			}(i)
		}

		wg.Wait()
		fmt.Println(len(m)) //все верно - 10
	}

	first()
	second()

	fmt.Println("----------------------------------------------------------------------------------------------")

}

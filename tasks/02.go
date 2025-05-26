/*
	Напишите функцию download.

	Функция download:
		•	на вход получает адреса для скачивания — urls
		•	конкурентно скачивает информацию из каждого url (для скачивания используйте функцию fakeDownload)
		•	если вызовы fakeDownload возвращают ошибки, то нужно вернуть их все (см. errors.Join)

	// timeoutLimit - вероятность, с которой не будет возвращаться ошибка от fakeDownload():
	// timeoutLimit = 100 - ошибок не будет;
	// timeoutLimit = 0 - всегда будет возвращаться ошибка.
	// Можете "поиграть" с этим параметром, для проверки случаев с возвращением ошибки.
	const timeoutLimit = 90

	type Result struct {
		msg string
		err error
	}

	// fakeDownload - имитирует разное время скачивания для разных адресов
	func fakeDownload(url string) Result {
		r := rand.Intn(100)
		time.Sleep(time.Duration(r) * time.Millisecond)
		if r > timeoutLimit {
			return Result{
				err: errors.New(fmt.Sprintf("failed to download data from %s: timeout", url)),
			}
		}

		return Result{
			msg: fmt.Sprintf("downloaded data from %s\n", url),
		}
	}

	// download - параллельно скачивает данные из urls
	func download(urls []string) ([]string, error) {
		// напишите ваш код здесь
	}

	func main() {
		msgs, err := download([]string{
			"https://example.com/1.xml",
			"https://example.com/1.xml",
			"https://example.com/1.xml",
			"https://example.com/1.xml",
			"https://example.com/1.xml",
		})

		if err != nil {
			panic(err)
		}

		fmt.Println(msgs)
	}
*/

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const timeoutLimit = 90

type Result struct {
	msg string
	err error
}

func fakeDownload(url string) Result {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
	if r > timeoutLimit {
		return Result{
			err: errors.New(fmt.Sprintf("failed to download data from %s: timeout", url)),
		}
	}

	return Result{
		msg: fmt.Sprintf("downloaded data from %s\n", url),
	}
}

func download(urls []string) ([]string, error) {
	var msgs []string
	var errs []error
	out := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			out <- fakeDownload(url)
			defer wg.Done()
		}(url)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for v := range out {
		if v.err != nil {
			errs = append(errs, v.err)
		} else {
			msgs = append(msgs, v.msg)
		}
	}

	return msgs, errors.Join(errs...)

}

func main() {
	msgs, err := download([]string{
		"https://example.com/1.xml",
		"https://example.com/2.xml",
		"https://example.com/3.xml",
		"https://example.com/4.xml",
		"https://example.com/5.xml",
		"https://example.com/6.xml",
		"https://example.com/7.xml",
		"https://example.com/8.xml",
		"https://example.com/9.xml",
		"https://example.com/10.xml",
		"https://example.com/11.xml",
		"https://example.com/12.xml",
		"https://example.com/13.xml",
		"https://example.com/14.xml",
		"https://example.com/15.xml",
	})
	fmt.Println(msgs)

	if err != nil {
		panic(err)
	}

}

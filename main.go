package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		results := make([]chan int, n)
		for i := range n {
			resultChan := make(chan int, 1)
			results[i] = resultChan

			x1 := <-in1
			x2 := <-in2
			go func(resultChan chan int, x1, x2 int) {
				ch1 := make(chan int, 1)
				ch2 := make(chan int, 1)
				go func() { ch1 <- fn(x1) }()
				go func() { ch2 <- fn(x2) }()
				result := <-ch1 + <-ch2
				resultChan <- result
			}(resultChan, x1, x2)
		}
		for _, channel := range results {
			out <- <-channel
			close(channel)
		}
	}()
}

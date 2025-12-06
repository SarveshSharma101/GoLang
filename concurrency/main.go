package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// WaitGroup()
	// BlockingUsingChannels()
	// BufferedChan()
	// LockUnlock()
	// FanOut()
	FanIn()
}

func WaitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go count(i+1, &wg)
	}
	wg.Wait()
}

func count(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("goroutine %v counting %v\n", id, i)
	}

}

func BlockingUsingChannels() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go RecieveGoR(ch, &wg)
	go SenderGoR(ch, true, &wg)
	wg.Wait()
}

func RecieveGoR(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		x, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("->", x)
	}
}

func SenderGoR(ch chan<- int, wait bool, wg *sync.WaitGroup) {
	defer wg.Done()
	if wait {
		time.Sleep(10 * time.Second)
	}
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func BufferedChan() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(2)
	go SenderGoR(ch, false, &wg)
	time.Sleep(5 * time.Second)
	go RecieveGoR(ch, &wg)
	wg.Wait()
}

func LockUnlock() {
	c := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			c++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("->", c)
}

func Job(job <-chan int, res chan<- int, wg *sync.WaitGroup) {
	// defer wg.Done()
	for j := range job {
		res <- j * 2
	}

}

func FanOut() {
	job := make(chan int)
	res := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		// wg.Add(1)
		go Job(job, res, &wg)
	}

	for i := 0; i < 10; i++ {
		job <- i + 1
	}
	close(job)
	for i := 0; i < 10; i++ {
		fmt.Println(<-res)
	}
	// wg.Wait()

}

func FanIn() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 10)

	go MultiBy(2, 5, ch1)
	go MultiBy(3, 10, ch2)

	op1, op2 := true, true

	for op1 || op2 {
		select {
		case v, ok := <-ch1:
			if !ok {
				op1 = false
			} else {
				fmt.Println("from ch1: ", v)
			}
		case v, ok := <-ch2:
			if !ok {
				op2 = false
			} else {
				fmt.Println("from ch2: ", v)
			}
		}

	}
}

func MultiBy(x, y int, ch chan<- int) {
	for i := 0; i <= y; i++ {
		ch <- x * i
	}
	close(ch)
}

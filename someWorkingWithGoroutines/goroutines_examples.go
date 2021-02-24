package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("You shouldn't do so")
	wrongUse()

	fmt.Println("Use this function!")
	rightUse()

	fmt.Println("Or this one!")
	rightUseRightOrder()


	fmt.Println("Or this one too!")
	anotherRight()

}


func wrongUse() {

	var i int
	var wg sync.WaitGroup

	fmt.Println("Function result: ")
	for i = 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func rightUse() {

	var i int
	var wg sync.WaitGroup

	fmt.Println("Function result: ")
	for i = 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func rightUseRightOrder() {
	var cnt int32 = 0
	var wg sync.WaitGroup

	fmt.Println("Function result: ")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&cnt, 1)
			fmt.Println(cnt)
			wg.Done()
		}()
	}
	wg.Wait()
}

func anotherRight() {
	ch := make(chan int)
	waitG := sync.WaitGroup{}

	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			waitG.Add(1)
			go func(i int, ch chan int) {
				ch <- i
				waitG.Done()
			}(i, ch)
		}
		waitG.Wait()

		done <- true
	}()
	defer close(ch)
	defer close(done)
	for {
		select {
		case val := <- ch:
			fmt.Println(val)

		case <-done:
			fmt.Println("Done")
			//os.Exit(1)
			return
		}
	}

}

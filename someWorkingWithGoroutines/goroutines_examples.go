package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("You shouldn't do so")
	wrongUse()

	fmt.Println("Use this function!")
	rightUse()
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


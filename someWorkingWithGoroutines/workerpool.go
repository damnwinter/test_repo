package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Worker pool")
	jobsCount := 10
	jobs := make(chan int, jobsCount)
	jobsOut := make(chan int, jobsCount)

	for i := 0; i < jobsCount; i++ {
		jobs <- i
	}
	close(jobs)


	wg := sync.WaitGroup{}
	go func() {
		workerCount := 3
		for i := 0; i < workerCount; i ++ {
			wg.Add(1)
			go worker(i + 1, jobs, jobsOut, &wg)
		}
		wg.Wait()
		close(jobsOut)
	}()

	for out := range jobsOut {
		fmt.Println("Jobs output: ", out)
	}

}

func worker(id int, jobs chan int, jobsOut chan int, group *sync.WaitGroup) {
	for val := range jobs {
		fmt.Printf("Worker %d start job\n", id)
		jobsOut <- val
		fmt.Printf("Worker %d done job\n", id)
	}
	group.Done()
}


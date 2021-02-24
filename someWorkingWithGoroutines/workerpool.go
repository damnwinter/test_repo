package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Worker pool 1")
	workerPoolAllFixed()

	fmt.Println("Worker pool 2")
	workerPoolAnyJobsFixedWorker(4)
}


// worker live, while execute jobs (one worker -> done only one job)
func workerPoolAnyJobsFixedWorker(workerCount int) {

	jobsCount := 100

	jobs := make(chan int, workerCount)
	jobsOut := make(chan int, workerCount)

	go func() {
		for i := 0; i < jobsCount; i++ {
			jobs <- i
		}
	}()

	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < jobsCount; i++ {
			wg.Add(1)
			go workerAnyJobsFixedWorker(i + 1, jobs, jobsOut, &wg)
		}
		wg.Wait()
		close(jobs)
		close(jobsOut)
	}()

	for out := range jobsOut {
		fmt.Printf("Jobs %d output\n", out)
	}

}

func workerAnyJobsFixedWorker(id int, jobs <-chan int, jobsOut chan<- int, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Printf("Worker %d started\n", id)

	job := <- jobs
	jobsOut <- job
	fmt.Printf("Worker %d finished\n", id)
}




// Fixed worker count, fixed jobs (all started know)
func workerPoolAllFixed() {
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
			go workerAllFixed(i + 1, jobs, jobsOut, &wg)
		}
		wg.Wait()
		close(jobsOut)
	}()

	for out := range jobsOut {
		fmt.Println("Jobs output: ", out)
	}
}

func workerAllFixed(id int, jobs chan int, jobsOut chan int, group *sync.WaitGroup) {
	for val := range jobs {
		fmt.Printf("Worker %d start job\n", id)
		jobsOut <- val
		fmt.Printf("Worker %d done job\n", id)
	}
	group.Done()
}


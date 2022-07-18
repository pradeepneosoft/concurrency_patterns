package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	numberGoroutine = 4
	taskLoad        = 10
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutine)
	for gr := 1; gr <= numberGoroutine; gr++ {
		go Worker(tasks, gr)
	}
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("(task %d )", post)
	}
	close(tasks)
	wg.Wait()
}
func Worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Println("shutting down worker", worker)
			return
		}
		fmt.Println("worker started task ", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Println("completed task", worker, task)
	}
}

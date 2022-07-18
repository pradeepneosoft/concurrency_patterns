package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1
	wg.Wait()

}
func Runner(baton chan int) {
	var newRunner int
	runner := <-baton

	fmt.Println("runner running with baton ", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Println("runner to the line ", newRunner)
		go Runner(baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Println("runner fineshed ,race over ", runner)
		wg.Done()
		return
	}
	fmt.Println("runner xechanged with ", runner, newRunner)
	baton <- newRunner

}

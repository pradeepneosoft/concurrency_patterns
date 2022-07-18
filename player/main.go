package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//natural no. print
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	court := make(chan int)
	wg.Add(2)
	go Player("saina ", court)
	go Player("rafel", court)
	court <- 1
	wg.Wait()
}

func Player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Println("player won ", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println("player missed", name)
			close(court)
			return
		}
		fmt.Println("player hit ", name, ball)
		ball++
		court <- ball

	}
}

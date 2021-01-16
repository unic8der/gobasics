package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go jobs(c1)

	go workerJobs(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Goshhh I'm tired!!!")
}

func jobs(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func workerJobs(c1, c2 chan int) {
	var wg sync.WaitGroup
	for ci := range c1 {
		wg.Add(1)
		go func(i int) {
			c2 <- doSomething(i)
			wg.Done()
		}(ci)
	}
	wg.Wait()
	close(c2)
}

func doSomething(n int) int {
	r := rand.Intn(n + 1*10)
	time.Sleep(time.Second * time.Duration(r))
	return n + r
}

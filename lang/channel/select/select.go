package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1000)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	//var c1, c2 chan int // c1 and c2 = nil

	var c1, c2 = generator(), generator()

	var worker = createWorker(0)

	var values []int

	tm := time.After(time.Second * 10)
	tick := time.Tick(time.Second)
	//n := 0
	//hasValue := false
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			//w <- n
			values = append(values, n)
			//hasValue = true
		case n := <-c2:
			//w <- n
			values = append(values, n)
			//hasValue = true
		case activeWorker <- activeValue:
			values = values[1:]
			//hasValue = false
			//fmt.Println("Received from c2: ", n)
		case <-time.After(time.Millisecond * 600):
			fmt.Println("timeout")
		case <-tm:
			fmt.Println("Bye")
			return
		case <-tick:
			fmt.Println("queue len =", len(values))
		}
	}
}

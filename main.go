package main

import (
	"fmt"
	"sync"
)

type atomicCounter struct {
	mut   sync.Mutex
	count int
}

func (ac *atomicCounter) incrementBy(n int) {
	ac.mut.Lock()
	ac.count += n
	ac.mut.Unlock()
}

func dataRace() {
	counter := atomicCounter{}
	done1 := make(chan bool)
	done2 := make(chan bool)
	go func() {
		for i := 0; i < 1000; i++ {
			counter.incrementBy(1)
		}
		done1 <- true
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			counter.incrementBy(1)
		}
		done2 <- true
	}()
	<-done1
	<-done2
	fmt.Println("The counter value is: ", counter.count)
}

func main() {
	for i := 0; i < 20; i++ {
		dataRace() //This will produce the same output each run
	}
}

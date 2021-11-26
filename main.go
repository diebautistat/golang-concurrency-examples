package main

import (
	"fmt"
)

func dataRace() {
    var count int
	done1 := make(chan bool)
    done2 := make(chan bool)
	go func() {
		for i := 0; i < 1000; i++ {
			count++ //non atomic operation
		}
		done1 <- true
	}()
    go func() {
		for i := 0; i < 1000; i++ {
			count++ //non atomic operation
		}
		done2 <- true
	}()
	<-done1
    <-done2
	fmt.Println("The counter value is: ", count)
}

func main() {
	for i := 0; i < 20; i++{
        dataRace() //This might produce different outputs each run
    }
}

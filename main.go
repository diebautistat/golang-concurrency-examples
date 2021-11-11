package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=0; i<10; i++ {
		go expensiveTask(i)
	}
	time.Sleep(2 * time.Second)
}

func expensiveTask(value int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Ran using value", value)
}




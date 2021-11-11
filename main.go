package main

import (
	"fmt"
)

func sum(numbers []int, channel chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	channel <- sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	channel := make(chan int)
	go sum(numbers[:len(numbers)/2], channel)
	go sum(numbers[len(numbers)/2:], channel)
	fmt.Println("First gotten result", <- channel)
	secondResult := <- channel
	fmt.Println("Second gotten result", secondResult)
}


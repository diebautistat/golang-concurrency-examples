package main

import (
	"fmt"
	"time"
)

func twoSecondsTask(results chan string){
    time.Sleep(2 * time.Second)
    results <- "result"
}

func main() {
    results := make(chan string)
    go twoSecondsTask(results)
    select {
    case <- results:
        fmt.Println("2 seconds task done")
    case <- time.After(time.Second):
        fmt.Println("timeout!") //Happens first
    }
}



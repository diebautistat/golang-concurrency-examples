package main

import (
    "fmt"
    "time"
)

func oneSecondTask(results chan string){
    time.Sleep(time.Second)
    results <- "Done"
}

func twoSecondsTask(results chan string){
    time.Sleep(2 * time.Second)
    results <- "Done"
}

func main() {
    channel1 := make(chan string)
    channel2 := make(chan string)
    go oneSecondTask(channel1)
    go twoSecondsTask(channel2)
    select {
    case <- channel1:
        fmt.Println("Received from first channel") //Happens first
    case res := <- channel2:
        fmt.Println("Received from second channel", res)
    } 
}



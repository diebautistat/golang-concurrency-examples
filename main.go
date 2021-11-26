package main

import (
    "fmt"
)

func randomDigits(results chan int, limit int){
    for i := 0; i < limit; i++ {
        select{
        case results <- 0:
        case results <- 1:
        case results <- 2:
        case results <- 3:
        case results <- 4:
        case results <- 5:
        case results <- 6:
        case results <- 7:
        case results <- 8:
        case results <- 9:
        }
    }
    close(results)
}

func main() {
    series := make(chan int)
    go randomDigits(series, 10)
    for number := range series {
        fmt.Println(number)
    } 
}



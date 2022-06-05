
package main

import (
    "fmt"
    "time"
)

func printMessage(message string, wait time.Duration, c chan int) {
    time.Sleep(wait)
    fmt.Println(message)
    c <-0
}

func main() {

    channel := make(chan string)

    go func() { channel <- "message"}()

    fmt.Println(<-channel)

    c := make(chan int)

    go printMessage("start", 3*time.Second, c)
    <- c

    for i := 0; i <= 5; i++ {
        time.Sleep(time.Second)
        fmt.Println("time:", i)
    }

    go printMessage("end", 3*time.Second, c)
    <- c
}

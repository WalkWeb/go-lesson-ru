
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for {
            time.Sleep(time.Second)
            c1 <- "message"
        }
    }()

    go func() {
        for {
            time.Sleep(time.Second * 2)
            c2 <- "message"
        }
    }()

    for {
        select {
            case message1 := <-c1:
                fmt.Println("received: ", message1)
            case message2 := <-c2:
                fmt.Println("received: ", message2)
            default:
                time.Sleep(time.Microsecond * 50)
        }
    }

}


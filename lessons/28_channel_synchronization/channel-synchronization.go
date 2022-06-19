
package main

import (
    "fmt"
    "time"
)

func curl(channel chan bool) {
    fmt.Println("make curl request...")
    time.Sleep(time.Second)
    fmt.Println("response received")

    channel <- true
}

func main() {

    channel := make(chan bool)
    for i := 0; i < 3; i++ {
        go curl(channel)
    }

    for i := 0; i < 3; i++ {
        <- channel
    }

    fmt.Println("end")
}

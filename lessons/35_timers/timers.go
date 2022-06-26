
package main

import (
    "fmt"
    "time"
)

func main() {

    timer1 := time.NewTimer(2 * time.Second)

    <-timer1.C

    fmt.Println("Timer1 worked")

    timer2 := time.NewTimer(1 * time.Second)

    go func() {
       <-timer2.C
       fmt.Println("Timer2 worked")
    }()

    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer2 stopped")
    }

    time.Sleep(3 * time.Second)
}


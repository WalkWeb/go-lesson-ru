
package main

import (
    "fmt"
    "time"
)

func main() {

    // Example #1

    requests := make(chan int, 15)

    for i := 0; i < 5; i++ {
        requests <- 1
    }

    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <- limiter
        fmt.Println("request", req, time.Now())
    }

    // Example #2

    plosiveRequests := make(chan int, 5)

    for i := 0; i < 5; i++ {
        plosiveRequests <- i
    }

    close(plosiveRequests)

    plosiveLimiter := make(chan time.Time, 3)

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            plosiveLimiter <- t
        }
    }()

    for req := range plosiveRequests {
        <-plosiveLimiter
        fmt.Println("plosive request", req, time.Now())
    }
}

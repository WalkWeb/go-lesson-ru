
package main

import "fmt"

func main() {

    channel := make(chan int, 2)

    channel <- 1
    channel <- 2

    close(channel)

    for element := range channel {
        fmt.Println(element)
    }

}

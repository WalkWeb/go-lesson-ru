
package main

import "fmt"

func send(channel chan<- string) {
    channel <- "message"
}

func received(channel <-chan string) {
    fmt.Println(<-channel)
}

func main() {

    channel := make(chan string, 1)

    send(channel)
    received(channel)

}

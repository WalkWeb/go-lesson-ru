
package main

import "fmt"

func main() {

    channel := make(chan string, 2)

    channel <- "message 1"
    channel <- "message 2"

    fmt.Println(<- channel)
    fmt.Println(<- channel)

}

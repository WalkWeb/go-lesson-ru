
package main

import "fmt"

func intClosures() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    anonymous := func() {
        fmt.Println("run anonymous function")
    }

    anonymous()

    xClosures := intClosures()

    fmt.Println(xClosures())
    fmt.Println(xClosures())
    fmt.Println(xClosures())

    yClosures := intClosures()

    fmt.Println(yClosures())
}

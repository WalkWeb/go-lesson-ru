
package main

import "fmt"

func recovery() {
    if recover := recover(); recover != nil {
        fmt.Println("Recovered error:", recover)
    }
}

func someProcess() {
    defer recovery()
    fmt.Println("some process...")
    panic("Panic!")
}

func main() {
    someProcess()
    fmt.Println("After panic")
}

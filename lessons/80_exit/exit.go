
package main

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("This code will not be executed")

    os.Exit(3)
}

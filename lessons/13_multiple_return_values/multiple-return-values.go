
package main

import "fmt"

func numeric() (int, string) {
    return 5, "five"
}

func main() {

    number, name := numeric()

    fmt.Println("number:", number)
    fmt.Println("name:", name)

    _, n := numeric()
    fmt.Println("n:", n)

}

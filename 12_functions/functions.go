
package main

import "fmt"

func addition(a int, b int) int {
    return a + b
}

func subtraction(a, b int) int {
    return a - b
}

func main() {
    fmt.Println("10 + 5 =", addition(10, 5))
    fmt.Println("10 - 5 =", subtraction(10, 5))
}

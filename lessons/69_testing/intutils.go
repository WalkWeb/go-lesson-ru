
package main

import (
    "fmt"
)

func MinInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {

    fmt.Println(MinInt(10, 5))

}

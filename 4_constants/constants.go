
package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {

    fmt.Println(s)

    const n = 500000000 // Equivalent: const n float64 = 500000000
    const d = 3e20 / n

    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))

    const (
        pi float64 = 3.1415
        e float64 = 2.7182
    )

    const k, l = 20, 2000

    fmt.Println(pi)
    fmt.Println(e)
    fmt.Println(k)
    fmt.Println(l)

    const (
        a = 1
        b
        c
    )

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

//     var p = "page"
//     const badConst = p // Error: const initializer p is not a constant
}

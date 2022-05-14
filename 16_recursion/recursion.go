
package main

import "fmt"

func fact(n int) int {
    if n == 0 {
    fmt.Println("Recursion end")
        return 1
    }
    fmt.Println("Make recursion. n: ", n)
    return n * fact(n - 1)
}

func main() {
    fmt.Println((fact(7)))

    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n - 1) + fib(n - 2)
    }

    fmt.Println("fin(7) =", fib(7))
}

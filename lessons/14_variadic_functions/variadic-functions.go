
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " \n")

    sum := 0
    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum)
}

func main() {
    sum(1, 2, 3)

    slice := []int{10, 20, 30}
    sum(slice...)
}

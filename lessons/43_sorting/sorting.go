
package main

import (
    "fmt"
    "sort"
)

func main() {

    messages := []string{"c", "a", "b"}
    sort.Strings(messages)
    fmt.Println("Messages:", messages)

    integers := []int{10, 4, 6}
    sort.Ints(integers)
    fmt.Println("Integers:", integers)

    isSorted := sort.IntsAreSorted(integers)
    fmt.Println("Is sorted:", isSorted)
}

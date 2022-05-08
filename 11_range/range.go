
package main

import "fmt"

func main() {

    // Это срез, если бы было написано nums := [3]int{2, 3, 4} - то это был бы массив
    nums := []int{2, 3, 4}
    sum := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum)

    for index, num := range nums {
        fmt.Println("index:", index, "num:", num)
    }

    languages := map[string]string{"p": "php", "g": "golang"}
    for k, v := range languages {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range languages {
        fmt.Println("key:", k)
    }

    for i, c := range "golang" {
        fmt.Println(i, c)
    }

}

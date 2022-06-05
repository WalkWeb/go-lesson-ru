
package main

import "fmt"

func main() {

    if (3 > 5) {
        fmt.Println("3 больше 5")
    } else {
        fmt.Println("3 не больше 5")
    }

    if (8%4 == 0) {
        fmt.Println("8 делится без остатка на 4")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "отрицательное")
    } else if (num < 10) {
        fmt.Println(num, "это 1 цифра")
    } else {
        fmt.Println(num, "имеет несколько цифр")
    }
}


package main

import "fmt"

func no_link(i int) {
    fmt.Println("Start no_link() function - value:", i)
    i = 0
    fmt.Println("End no_link() function - value:", i)
    fmt.Println("no_link() function i pointer:", &i)
}

func is_link(i *int) {
    fmt.Println("Start is_link() function - value:", i)
    *i = 0
    fmt.Println("End is_link() function - value:", i)
    fmt.Println("is_link() function i pointer:", &i)
}

func main() {
    i := 5

    fmt.Println("initial i:", i)
    fmt.Println("pointer:", &i)

//     no_link(i)
//     fmt.Println("no_link(i):", i)

    is_link(&i)

    fmt.Println("is_link(i):", i)
    fmt.Println("pointer:", &i)
}

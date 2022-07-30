
package main

import (
    "fmt"
)

func main() {

    panic("Do panic")

    fmt.Println("This code will not be executed")

    // Panic caused by accessing a non-existent array index, uncomment the code
    //arr := [5]int{}
    //arr[10] = 100

}

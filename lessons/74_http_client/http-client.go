
package main

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {

    response, err := http.Get("https://example.com/")
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

    fmt.Println("Response status code:", response.Status)

    scanner := bufio.NewScanner(response.Body)

    for i := 0; scanner.Scan() && i < 4; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

}

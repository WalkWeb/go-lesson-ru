
package main

import "fmt"

func main() {
    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("s append:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("c copy:", c)

    l := s[2:5]
    fmt.Println("l s:", l)

    l = s[:5]
    fmt.Println("s[:5]:", l)

    l = s[2:]
    fmt.Println("s[2:]:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("t:", t)

    twoS := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoS[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoS[i][j] = i + j
        }
    }
    fmt.Println("2d slice: ", twoS)
}

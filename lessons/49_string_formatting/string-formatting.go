
package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
    fmt.Printf("struct1: %v\n", p)

    fmt.Printf("struct2: %+v\n", p)

    fmt.Printf("struct3: %#v\n", p)

    fmt.Printf("type: %T\n", p)

    fmt.Printf("bool: %t\n", true)

    fmt.Printf("int: %d\n", 10)

    fmt.Printf("bin: %b\n", 10)

    fmt.Printf("hex: %x\n", 10)

    fmt.Printf("float1: %f\n", 10.5)

    fmt.Printf("float2: %e\n", 10.5)

    fmt.Printf("float3: %E\n", 10.5)

    fmt.Printf("char: %c\n", 33)

    fmt.Printf("string1: %s\n", "apple")

    fmt.Printf("string2: %q\n", "apple")

    fmt.Printf("string3: %x\n", "apple")

    fmt.Printf("pointer: %p\n", &p)

    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

    s := fmt.Sprintf("sprintf: %s", "string")
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}

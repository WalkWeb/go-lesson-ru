
package main

import (
    "fmt"
    "os"
)

func createFile(filePath string) *os.File {
    fmt.Println("Create file:", filePath)
    file, err := os.Create(filePath)
    if err != nil {
        panic(err)
    }
    return file
}

func writeFile(file *os.File, content string) {
    fmt.Println("Write file:", content)
    _, err := file.WriteString(content)
    if err != nil {
        panic(err)
    }
}

func closeFile(file *os.File) {
    fmt.Println("Close file")
    err := file.Close()
    if err != nil {
        panic(err)
    }
}

func main() {
    file := createFile("/tmp/defer")
    defer closeFile(file)
    writeFile(file, "file content\n")
}

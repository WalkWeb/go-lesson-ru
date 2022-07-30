
package main

import (
    "fmt"
    "sort"
)

type byLength []string

func (s byLength) Len() int {
    return len(s)
}

func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {

    messages := []string{"aaa", "cc", "b"}
    sort.Sort(byLength(messages))
    fmt.Println(messages)

}

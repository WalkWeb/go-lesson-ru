
package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {

    p("Contains: ", s.Contains("test", "es"))
    p("Count:    ", s.Count("test", "t"))
    p("HasPrefix:", s.HasPrefix("test", "te"))
    p("HasSuffix:", s.HasSuffix("test", "xxx"))
    p("Index:    ", s.Index("task", "a"))
    p("Join:     ", s.Join([]string{"a", "b", "c"}, "---"))
    p("Repeat:   ", s.Repeat("a", 5))
    p("Replace:  ", s.Replace("xxaaxx", "a", "1", 1))
    p("Split:    ", s.Split("a-b-c", "-"))
    p("ToLower:  ", s.ToLower("Name"))
    p("ToUpper:  ", s.ToUpper("Name"))

}

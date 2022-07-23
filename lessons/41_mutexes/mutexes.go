
package main

import (
    "fmt"
    "sync"
)

type Container struct {
    mutex    sync.Mutex
    counters map[string]int
}

func (c *Container) inc(name string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.counters[name]++
}

func main() {

    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }

    var wg sync.WaitGroup

    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }

    wg.Add(3)

    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    wg.Wait()

    fmt.Println(c.counters)
}

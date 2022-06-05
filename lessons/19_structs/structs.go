
package main

import "fmt"

type person struct {
    name string
    age  int
}

func createPerson(name string, age int) *person {
    return &person {name, age}
}

func main() {
    fmt.Println(person{"Bob", 20})

    fmt.Println(person{age: 30, name: "Alice"})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{"Anna", 25})

    fmt.Println(createPerson("Olga", 40))

    p := person{"Nikolai", 33}
    fmt.Println(p.name)

    p.age = 50
    fmt.Println(p.age)
}

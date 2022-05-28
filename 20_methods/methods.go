
package main

import "fmt"

type square struct {
    width, height int
}

type user struct {
    name string
    year int
}

func (r square) area() int {
    return r.width * r.height;
}

func (r square) perimeter() int {
    return 2*r.width + 2*r.height;
}

func (u user) birthday_no_link() int {
    u.year = u.year + 1;
    return u.year
}

func (u *user) birthday_link() int {
    u.year = u.year + 1;
    return u.year
}

func main() {

    s := square{10, 5}

    fmt.Println(s)
    fmt.Println("square area:", s.area())
    fmt.Println("square perimeter:", s.perimeter())

    u1 := user{"Vasya", 30}

    fmt.Println(u1)
    fmt.Println(u1.birthday_no_link())
    fmt.Println(u1)

    u2 := user{"Masha", 50}

    fmt.Println(u2)
    fmt.Println(u2.birthday_link())
    fmt.Println(u2)

}

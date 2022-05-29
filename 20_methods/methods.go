
package main

import "fmt"

type rectangle struct {
    width, height int
}

type user struct {
    name string
    year int
}

func (r rectangle) area() int {
    return r.width * r.height;
}

func (r rectangle) perimeter() int {
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

    r := rectangle{10, 5}

    fmt.Println(r)
    fmt.Println("rectangle area:", r.area())
    fmt.Println("rectangle perimeter:", r.perimeter())

    u1 := user{"Vasya", 30}

    fmt.Println(u1)
    fmt.Println(u1.birthday_no_link())
    fmt.Println(u1)

    u2 := user{"Masha", 50}

    fmt.Println(u2)
    fmt.Println(u2.birthday_link())
    fmt.Println(u2)

}

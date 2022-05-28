
# Изучаем язык Go #21 – Интерфейсы

Интерфейсы позволяют абстрагироваться от конкретных реализаций, работая с некой абстракцией, которая реализует 
необходимые нам [методы](https://github.com/WalkWeb/go-lesson-ru/tree/master/20_methods), при этом нам не важны детали, 
с чем именно мы работаем.

В этом уроке мы рассмотрим как создавать и работать с интерфейсами в Go.

Объявляются интерфейсы следующим образом:

```
type geometry interface {
    area() float64
    perimeter() float64
}
```

Создадим две структуры (объекта) прямоугольник `rectangle` и круг `circle`, реализующих данный интерфейс:

```
type rectangle struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rectangle) area() float64 {
    return r.width * r.height
}

func (r rectangle) perimeter() float64 {
    return 2 * r.width + 2 * r.height
}

func (c circle) area() float64 {
    return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```

Теперь добавим функцию, которая работает с интерфейсом, а не с конкретной реализацией:

```
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perimeter())
}
```

И осталось добавить код для демонстрации того, как оно работает:

```
    r := rectangle{3, 4}
    c := circle{5}

    measure(r)
    measure(c)
```

Выполнив этот код мы получим:

```
$ go run interfaces.go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```

Т.е. одна функция работает с разными структурами – потому что обе они соответствуют необходимому интерфейсу.

Обратите внимание, что в отличие от многих других языков, реализующих 
[объектно-ориентированная парадигму](https://ru.wikipedia.org/wiki/%D0%9E%D0%B1%D1%8A%D0%B5%D0%BA%D1%82%D0%BD%D0%BE-%D0%BE%D1%80%D0%B8%D0%B5%D0%BD%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%BD%D0%BE%D0%B5_%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5), 
в Go не нужно указывать, что такая-то структура (класс) реализует такой-то интерфейс, эта привязка происходит 
автоматически, если реализованы все необходимые [методы](https://github.com/WalkWeb/go-lesson-ru/tree/master/20_methods).

____

В этом уроке вы узнали:

1. Как создавать и использовать интерфейсы

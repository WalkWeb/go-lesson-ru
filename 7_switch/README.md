
# Изучаем язык Go #7 - Оператор switch

Оператор switch по своему действию аналогичен операторам [if/else if/else](https://github.com/WalkWeb/go-lesson-ru/tree/master/6_if_else), 
который позволяет писать длинные цепочки проверок более красивым и удобным для восприятия вида.

Базовый вариант использования switch выглядит так:

```
    i := 2

    fmt.Println("Write ", i, " as ")
    switch i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
    }
```

Который сделает одно из 4 вариантов действий (3 описаны, плюс один вариант когда `i` не равен ничему из перечисленного - 
в этом случае ничего не будет сделано).

Можно использовать несколько значений в одной проверке, через запятую, а также использовать `default`, который
будет выполняться если никаких совпадений не произошло:

```
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
```

Еще один способ использовать switch - когда в самом `case` происходит вычисление значения:

```
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
```

Но стоит отметить, что в этом случае обычный if/else при той же функциональности будет выглядеть лучше:

```
t := time.Now()
if (t.Hour() < 12) {
    fmt.Println("It's before noon")
} else {
    fmt.Println("It's after noon")
}
```

И еще один пример использования switch, который сравнивает типы, а не значение:

```
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }

    whatAmI(true) // I'm a bool
    whatAmI(1) // I'm an int
    whatAmI("hey") // Don't know type string
```
____


В этом уроке вы узнали:

1. Как использовать оператор switch

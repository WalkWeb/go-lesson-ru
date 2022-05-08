
# Изучаем язык Go #15 - Замыкания

Прежде чем переходить к замыканиям, нужно рассмотреть анонимные функции. Анонимные функции это функции без названия.

Например, можно сделать переменную как функцию:

```
func main() {

    anonymous := func() {
        fmt.Println("run anonymous function")
    }

    anonymous()
}
```

Выполнив код получим:

```
$ go run closures.go 
run anonymous function
```

Теперь перейдем к замыканию. Замыкание это особая анонимная функции, которая замыкается на свое же возвращаемое 
значение:

```
func intClosures() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

Вызывая её мы можем убедиться, что состояние `i` передается из одного вызова в другой:

```
    xClosures := intClosures()

    fmt.Println(xClosures())
    fmt.Println(xClosures())
    fmt.Println(xClosures())

    yClosures := intClosures()

    fmt.Println(yClosures())
```

Выполнив код получим:

```
$ go run closures.go 
// ...
1
2
3
1
```
____

В этом уроке вы узнали:

1. Что такое анонимные функции
2. Что такое замыкания

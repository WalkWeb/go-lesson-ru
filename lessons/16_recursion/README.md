
# Изучаем язык Go #16 – Рекурсия

Рекурсия – это функция, которая вызывает себя саму, а чтобы рекурсия не была бесконечной – отдельным условием функция 
прерывается.

Вот простой пример рекурсии в Go:

```
func fact(n int) int {
    if n == 0 {
    fmt.Println("Recursion end")
        return 1
    }
    fmt.Println("Make recursion. n: ", n)
    return n * fact(n - 1)
}
```

И её вызов:

```
func main() {
    fmt.Println((fact(7)))
}
```

Выполнив этот код получим:

```
$ go run recursion.go 
Make recursion. n:  7
Make recursion. n:  6
Make recursion. n:  5
Make recursion. n:  4
Make recursion. n:  3
Make recursion. n:  2
Make recursion. n:  1
Recursion end
5040
```

Рекурсию можно сделать и на [замыкании](https://github.com/WalkWeb/go-lesson-ru/tree/master/15_closures), но для этого
замыкание должно быть объявлено перед созданием:

```
    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n - 1) + fib(n - 2)
    }

    fmt.Println("fin(7) =", fib(7))
```

Если не написать `var fib func(n int) int`, а попытаться сразу сделать `fib := func ...`, то получим ошибку:

```
$ go run recursion.go 
# command-line-arguments
./recursion.go:24:16: undefined: fib
```

Это происходит из-за того, что компилятор Go проверяя код внутри анонимной функции, и видя вызов функции `fib` уже 
требует того, чтобы она была известна, т.е. была заранее объявлена.

Ну а если все написано правильно, при выполнении кода мы увидим:

```
$ go run recursion.go 
// ...
fin(7) = 13
```
____

В этом уроке вы узнали:

1. Что такое рекурсия и как их создавать в Go.

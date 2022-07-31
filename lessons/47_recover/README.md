
# Изучаем язык Go #47 – Recover

В отличие от ошибки уровня `log.Fatal()` ошибку уровня [panic](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/45_panic) 
можно перехватить, обработать и продолжить работу программы.

Рассмотрим как это сделать.

Начнем с функции, которая будет перехватывать панику, если она произошла:

```
func recovery() {
    if recover := recover(); recover != nil {
        fmt.Println("Recovered error:", recover)
    }
}
```

В `recover` мы получим то, что было передано в панику, например `panic("This is a panic")` вернет в `recover` – 
`This is a panic`

Далее напишем функцию, которая отвечает за основную логику программы, в которой может произойти паника. Для простоты 
создаем панику вручную:

```
func someProcess() {
    defer recovery()
    fmt.Println("some process...")
    panic("Panic!")
}
```

Обратите внимание, что функция отвечающая за перехват паники указана с ключевым словом 
[defer](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/46_defer)

Теперь осталось только вызвать функцию `someProcess()` и сделать что-то после неё, чтобы проверить, что программа не
завершается при панике:

```
func main() {
    someProcess()
    fmt.Println("After panic")
}
```

Выполнив программу получим:

```
$ go run recover.go 
some process...
Recovered error: Panic!
After panic
```

Мы получили именно то, что хотели – несмотря на панику, ошибка была перехвачена, а программа продолжила работать.

____

В этом уроке вы узнали:

- Как с помощью `recover()` перехватывать ошибку типа Panic

Читайте также:

- Официальную документацию по [Recover](https://go.dev/blog/defer-panic-and-recover)

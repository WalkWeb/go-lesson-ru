# Изучаем язык Go #25 – Горутины

Горутины - это специальный встроенный функционал в Go, который позволяет создавать легковесные и независимые (т.е. 
выполняющиеся параллельно) процессы. Которые, в свою очередь, позволяют делать многопоточную обработку данных, сразу на 
всех имеющихся ядрах в процессоре и за счет этого повышать производительность. К слову, если ядро в процессоре одно - то 
и никакой выгоды от многопоточности не будет.

Чтобы понять и на практике посмотреть разницу между выполнением кода в одном потоке, и параллельно, мы пойдем поэтапными
шагами, от простого к более сложному.

Для начала, напишем простую программу, которая считает до 5, делая на каждой итерации (перед выводом 0 тоже) задержку на 
секунду:

```
package main

import (
    "fmt"
    "time"
)

func main() {

    for i := 0; i <= 5; i++ {
        time.Sleep(time.Second)
        fmt.Println("time:", i)
    }

}
```

Выполнив эту программу мы получим:

```
$ go run goroutines.go 
time: 0
time: 1
time: 2
time: 3
time: 4
time: 5
```

Затем добавим функцию `printMessage`:

```
func printMessage(message string, wait time.Duration) {
    time.Sleep(wait)
    fmt.Println(message)
}
```

И добавим её использование до и после цикла с задержкой на 1 секунду:

```
func main() {

    go printMessage("start", 1*time.Second)

    for i := 0; i <= 5; i++ {
        time.Sleep(time.Second)
        fmt.Println("time:", i)
    }

    go printMessage("end", 1*time.Second)

}
```

Выполнив этот код получим:

```
$ go run goroutines.go 
start
time: 0
time: 1
time: 2
time: 3
time: 4
time: 5
end
```

Все очевидно и просто - наш код выполняется поэтапно. А теперь давайте подключим функционал горутин. Это делается просто - 
через добавление ключевого слова `go` перед вызовом метода. А также увеличим задержку до 3 секунд в методах
`printMessage`:

```
    go printMessage("start", 3*time.Second)

    for i := 0; i <= 5; i++ {
        time.Sleep(time.Second)
        fmt.Println("time:", i)
    }

    go printMessage("end", 3*time.Second)
```

Выполнив этот код получим:

```
$ go run goroutines.go 
time: 0
time: 1
start
time: 2
time: 3
time: 4
time: 5
```

Как видно - результат получился совершенно иным. Теперь все, что мы написали в `printMessage` выполняется в отдельном
независимом процессе и никак не задерживает выполнение кода из основной функции `main`. По этому сообщение start 
теперь выводится не в начале, а после "time: 1". Хотя иногда оно может выводиться и после "time: 2" - потому что оба
сообщения выводится с одинаковой задержкой, соответственно результат может быть неоднозначным.

А сообщение "end" пропало вовсе. Почему? Потому что выполнение функции `main` происходит тоже в горутине, а когда
родительская горутина выполняет все свои действия, она завершается, и также завершаются все её дочерние горутины. Проще 
говоря - сообщение "end" не было выведено, потому что горутина, которая отвечала за вывод этого сообщения была завершена
до того, как успела отработать.

Убедиться в этом (что она не исчезла, а просто не успела отработать) можно путем добавления в `main` задержки на 4 
секунды:

```
    time.Sleep(4*time.Second)
```

Выполнив этот код получим:

```
$ go run goroutines.go 
time: 0
time: 1
time: 2
start
time: 3
time: 4
time: 5
end
```

Во-первых, обратите внимание, что "start" на этот раз был выведен после "time: 2", такая ситуация случается у меня
где-то 1 раз на 10 запусков, а также на этот раз было выведено сообщение "end" - за счет добавленной задержки на 4
секунды вторая горутина успела отработать.

В завершение стоит добавить, что написание и понимание многопоточных программ требует большего опыта и абстрактного 
мышления, чем обычные программы, где компьютер просто поочередно выполняет указанные нами команды. 

В тоже время, многопоточное программирование в Go намного, НАМНОГО более простое в понимании, написании и использовании,
чем в JavaScript, где каждая команда (т.е. каждая строчка, в некотором упрощении) выполняется независимо от команд 
написанных ниже или выше её. И если кто-то вам говорит, что между nodejs (серверный вариант использования JavaScript) и 
Go нет разницы, и что nodejs также хороша, как Go, знайте, что этот человек неопытен. Потому что сам разработчик nodejs 
[говорит](https://habr.com/ru/post/413187/) о том, что Go имеет все плюсы nodejs, но не имеет недостатков JavaScript.

____

В этом уроке вы узнали:

1. Что такое многопоточное программирование
2. Как использовать горутины

# Изучаем язык Go #28 – Синхронизация каналов

Сложной частью в работе с асинхронным и параллельным выполнением кода (т.е. использованием 
[горутин](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/25_goroutines)) является синхронизация выполнения 
(там, где она нужна).

Для этого в го используются каналы, работа которых по-умолчанию сделана так, что она синхронизирует выполнение кода. 
Давайте рассмотрим это на примере.

Представим, что нам нужно делать http запросы к какому-то API в большом количестве. Разумеется, чтобы ускорить этот 
процесс имеет смысл распараллелить его.

Вначале сделаем этот функционал без синхронизации через каналы, а затем добавим каналы.

Итак, начнем с функции которая как-бы будет делать http-запрос:

```
func curl() {
    fmt.Println("make curl request...")
    time.Sleep(time.Second)
    fmt.Println("response received")
}
```

Его суть – это какая-то задержка, потому что http-запрос требует времени. Причем мы никогда не знаем, как быстро нам
ответит внешний сервер – он может ответить за 50 миллисекунд, а может за 5 секунд.

Теперь используем эту функцию в цикле. Представим, что нам нужно сделать 3 запроса, хотя это может быть и 300 и 3000:

```
func main() {
    for i := 0; i < 3; i++ {
        go curl()
    }

    fmt.Println("end")
}
```

Выполнив этот код мы получим...

```
$ go run channel-synchronization.go 
end
```

Согласитесь – совсем не то что мы ожидали. Почему так произошло? Потому что Go увидев, что функция `curl()` выполняется
в горутине, не ждет её завершения и сразу переходит к следующей операции – двум аналогичным запускам и финальной 
`fmt.Println("end")`, а выполнив последнюю инструкцию завершает текущую горутину `main()` и автоматически завершаются 
все дочерние горутины – которые даже не успевают начать выполнение.

Как можно исправить? Можно добавить задержку:

```
func main() {
    for i := 0; i < 3; i++ {
        go curl()
    }

    time.Sleep(time.Second * 2)
    fmt.Println("end")
}
```

Выполнив этот код мы получим:

```
$ go run channel-synchronization.go 
make curl request...
make curl request...
make curl request...
response received
response received
response received
end
```

Казалось бы, решение найдено. Но, как уже говорилось выше – мы не можем заранее знать, за сколько выполнился http-запрос.
И если он будет выполняться больше 2-х секунд, на которую мы поставили задержку – ответ обработан не будет.

В этом случае на помощь и приходят каналы с их встроенной механикой синхронизации.

Во-первых, усовершенствуем функцию `curl()` добавив в неё функционал передачи сообщения в канал при выполнении запроса.
Тем самым функция будет «говорить» каналу, что она закончила свою работу.

```
func curl(channel chan bool) {
    fmt.Println("make curl request...")
    time.Sleep(time.Second)
    fmt.Println("response received")

    channel <- true
}
```

А в функции `main()` добавим создание канала и чтение сообщений из него, а также удалим задержку на 2 секунды:

```
func main() {

    channel := make(chan bool)
    for i := 0; i < 3; i++ {
        go curl(channel)
    }

    for i := 0; i < 3; i++ {
        <- channel
    }

    fmt.Println("end")
}
```

Выполнив этот код получим:

```
$ go run channel-synchronization.go 
make curl request...
make curl request...
make curl request...
response received
response received
response received
end
```

Казалось бы то же самое, но это не так: во-первых, программа завершает работу сразу после того, как все запросы 
отработали (в варианте с `time.Sleep(time.Second * 2)` была еще секундная задержка), во вторых, вы можете изменить 
задержку в функции `curl()` на большую, и убедиться, что все также работает – т.е. теперь код точно дождется ответа
от внешнего API. За исключением, когда API вообще недоступно, и запрос оборвется по timeout, который по умолчанию в curl
составляет 60 секунд.

____

В этом уроке вы узнали:

1. Как синхронизировать работу горутин через каналы
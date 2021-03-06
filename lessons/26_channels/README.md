# Изучаем язык Go #26 – Каналы

При работе с [горутинами](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/25_goroutines) возникает несколько 
задач. Две самых базовых: это передача данных между горутинами и их синхронизация. Для этой задачи в Go созданы каналы и
в этом уроке мы начнем с ними знакомиться.

Каналы создаются с помощью функции `make()` и ключевого слова `chan`, указывающего, что мы создаем канал (channel):

```
    channel := make(chan string)
```

В каналы можно передавать и получать сообщения. Делается это простым синтаксисом через стрелочки `channel <-` и 
`<-channel`:

```
    go func() { channel <- "message"}()

    fmt.Println(<-channel)
```

В первой строчке мы создаем горутину, в ней передаем сообщение в канал, а во второй строчке получаем сообщение из канала
и выводим его. Выполнив этот код мы получим:

```
$ go run channels.go 
message
```

Как мы рассмотрели в [предыдущем уроке](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/25_goroutines) – 
горутины работают независимо друг от друга. Но передача и получение сообщений в каналах блокирует горутины. Убедиться в 
этом, а заодно сломать наш код легко - для этого просто уберем вызов `func() { channel <- "message"}()` через горутину, 
т.е. удалим перед неё ключевое слово `go`:

```
    channel := make(chan string)

    func() { channel <- "message"}()

    fmt.Println(<-channel)
```

Попробовав выполнить этот код получим ошибку:

```
$ go run channels.go 
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main.func1(...)
        /var/www/github_go/lessons/26_channels/channels.go:10
main.main()
        /var/www/github_go/lessons/26_channels/channels.go:10 +0x5a
exit status 2
```

Что происходит: на выполнение кода в функции `main()` Go создает горутину. Внутри неё мы передаем сообщение в канал. Эта
операция блокирует текущую горутину. Т.к. других горутин, которые могли бы разблокировать её нет – получаем `deadlock` –
блокировку которая не может завершиться, и чтобы компьютер не завис при выполнении такой программы – компилятор, видя
ошибку в коде, сообщает нам об этом и не дает скомпилировать программу.

Теперь посмотрим как можно синхронизировать работу программы, использующую горутины через каналы. Возьмем код из 
[предыдущего урока](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/25_goroutines) и несколько улучшим его.

В метод `printMessage` добавим отправку сообщения в канал в конце – таким образом мы будем сообщать другой горутине, что
эта горутина отработала:

```
func printMessage(message string, wait time.Duration, c chan int) {
    time.Sleep(wait)
    fmt.Println(message)
    c <-0
}
```

А после вызова `printMessage` добавим получение сообщения, с которым ничего делать не будем. Также, уберем в конце
задержку на 4 секунды:

```
    c := make(chan int)

    go printMessage("start", 3*time.Second, c)
    <- c

    for i := 0; i <= 5; i++ {
        time.Sleep(time.Second)
        fmt.Println("time:", i)
    }

    go printMessage("end", 3*time.Second, c)
    <- c
```

Выполнив этот код получим:

```
$ go run channels.go 
start
time: 0
time: 1
time: 2
time: 3
time: 4
time: 5
end
```

Все работает предсказуемым образом, несмотря на то, что в коде используется несколько горутин, со своими задержками. 
Вместо искусственных задержек, которые мы делали, это может быть обращение к внешнему API, когда мы не знаем, получим ли 
мы ответ сразу, или через 5 секунд.

Конечно, функционал каналов намного больше, чем показанный в этом уроке. В следующих уроках мы рассмотрим другие их
возможности.
____

В этом уроке вы узнали:

1. Как создавать и работать с каналами
2. Как передача и получение сообщений в каналах блокирует работу горутин

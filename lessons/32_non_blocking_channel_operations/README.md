
# Изучаем язык Go #32 – Неблокирующие операции с каналами

В этом уроке рассмотрим работу `default` в контексте получения сообщений из каналов через 
[select](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/30_select).

Для начала давайте создадим канал и попробуем получить из него сообщение (не отправляя его), напишем код:

```
    messages := make(chan string)

    <- messages
```

Если попытаться выполнить его получим ошибку:

```
$ go run non-blocking-channel-operations.go 
fatal error: all goroutines are asleep - deadlock!
```

Мы бы получили бесконечное «зависание» при запуске, потому что `<- messages` блокирует дальнейшее выполнение кода, а
сообщение в канал никто не отправляет. Но Go понимает ошибку написанного кода и сам завершает программу.

Чтобы сделать неблокирующую операцию чтения из канала необходимо использовать `select` и `default`:

```
    messages := make(chan string)

    select {
        case msg := <- messages:
            fmt.Println("received message", msg)
        default:
            fmt.Println("no message received")
    }
```

Выполнив этот код получим:

```
$ go run non-blocking-channel-operations.go 
no message received
```

Все как и ожидалось – никто не отправлял сообщение в канал, и выполнилась команда по умолчанию, написанная в блоке 
`default`.

Теперь рассмотрим неблокирующий вариант отправки сообщения. Добавьте код:

```
    msg := "hi"

    select {
        case messages <- msg:
            fmt.Println("sent message", msg)
        default:
            fmt.Println("no message sent")
    }
```

Выполнив этот код получим:

```
$ go run non-blocking-channel-operations.go 
no message received
no message sent
```

Почему так произошло? Потому что сообщение не может быть отправлено в канал, потому что в канале нет буфера и нет 
получателя.

Если добавить буфер:

```
messages := make(chan string, 1)
```

И выполнить код еще раз, сообщение отправится:

```
$ go run non-blocking-channel-operations.go 
no message received
sent message hi
```

И рассмотрим третий вариант, с неблокирующим получением сообщений сразу из нескольких каналов. Уберите буферизацию в 
канале, которую мы добавляли на предыдущем шаге и добавьте код:

```
    signals := make(chan bool)

    select {
        case msg := <-messages:
            fmt.Println("received message", msg)
        case sig := <-signals:
            fmt.Println("received signal", sig)
        default:
            fmt.Println("no activity")
    }
```

Выполнив этот код получим:

```
$ go run non-blocking-channel-operations.go 
no message received
no message sent
no activity
```

Сообщений нет ни в одном из каналов, и как и прежде отработал блок команд написанных в `default`.

Но если опять добавить буферизацию канала `messages`:

```
messages := make(chan string, 1)
```

И выполнить код, то получим:

```
$ go run non-blocking-channel-operations.go 
no message received
sent message hi
received message hi
```

____

В этом уроке вы узнали:

1. Как делать неблокирующее чтение из каналов

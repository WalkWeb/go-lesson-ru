
# Изучаем язык Go #29 – Направления каналов

При использовании каналов как параметров функции можно указывать, предназначен ли канал только для записи в рамках
данной функции, или наоборот, только для получения сообщений. Это позволяет писать более строго-типизированный код, 
который поможет избежать случайных ошибок.

Так указывается канал только на отправку сообщений в канал:

```
func send(channel chan<- string) {
    channel <- "message"
}
```

А так – только на получение сообщений из канала:

```
func received(channel <-chan string) {
    fmt.Println(<-channel)
}
```

Используем эти функции:

```
func main() {

    channel := make(chan string, 1)

    send(channel)
    received(channel)

}
```

Выполнив этот код получим:

```
$ go run channel-directions.go 
message
```

Вы можете проверить, и попробовать получить сообщение в функции `send()`:

```
func send(channel chan<- string) {
    channel <- "message"
    fmt.Println(<-channel)
}
``` 

Попробовав выполнить такой код получим ожидаемую ошибку:

```
$ go run channel-directions.go 
# command-line-arguments
./channel-directions.go:8:17: invalid operation: <-channel (receive from send-only type chan<- string)
```

____

В этом уроке вы узнали:

1. Как в параметрах функции/метода указывать каналы только на отправку/получение сообщений.

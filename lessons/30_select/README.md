
# Изучаем язык Go #30 – Select

Внимательный читатель заметил такую проблему с каналами, которые использовались в предыдущих уроках, что количество
получаемых сообщений в них строго определено. А как получать сообщения из каналов, если мы заранее не знаем сколько их 
будет? И как получать сообщения сразу из нескольких каналов?

Для этих задач в Go существует специальная конструкция `Select` и в этом уроке мы познакомимся с её работой.

Создадим два канала, и отправим в них сообщения:

```
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()
```

Рассмотрим пример, как слушать сообщения из двух каналов одновременно с помощью `Select`. Добавьте код ниже:

```
    for i := 0; i < 2; i++ {
        select {
            case message1 := <-c1:
                fmt.Println("received: ", message1)
            case message2 := <-c2:
                fmt.Println("received: ", message2)
        }
    }
```

Выполнив этот код получим:

```
$ go run select.go 
received:  one
received:  two
```

Теперь рассмотрим вариант, как слушать сообщения в канале дольше. Расширим функционал `select` добавив в него стадию
`default` а также увеличим количество итераций до 5:

```
    for i := 0; i < 5; i++ {
        select {
            case message1 := <-c1:
                fmt.Println("received: ", message1)
            case message2 := <-c2:
                fmt.Println("received: ", message2)
            default:
                time.Sleep(time.Second)
                fmt.Println("wait")
        }
    }
```

Выполнив этот код получим:

```
$ go run select.go 
wait
received:  one
wait
received:  two
wait
```

Ну а теперь перейдем к более реалистичному варианту, когда наша программа просто слушает два канала и обрабатывает
сообщения в них бесконечно долго:

```
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for {
            time.Sleep(time.Second)
            c1 <- "message"
        }
    }()

    go func() {
        for {
            time.Sleep(time.Second * 2)
            c2 <- "message"
        }
    }()

    for {
        select {
            case message1 := <-c1:
                fmt.Println("received: ", message1)
            case message2 := <-c2:
                fmt.Println("received: ", message2)
            default:
                time.Sleep(time.Microsecond * 50)
        }
    }
```

Выполнив программу получим:

```
$ go run select.go 
received:  message
received:  message
received:  message
received:  message
received:  message
```

При этом программа продолжит работать и получать сообщения, потому что мы написали код так, что она бесконечно 
генерирует и отправляет сообщения в каналы, и бесконечно слушает их. Чтобы принудительно остановить её нажмите `Ctrl+C`

____

В этом уроке вы узнали:

1. Как использовать `Select` для получения сообщений из каналов.

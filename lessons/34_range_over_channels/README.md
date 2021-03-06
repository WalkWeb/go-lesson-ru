
# Изучаем язык Go #34 – Range по каналам

В [предыдущем](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/33_closing_channels) уроке мы подробно 
разобарли особенность закрытия каналов, и как это можно использовать для корректного получения всех сообщений из канала
с помощью цикла.

В этом уроке мы рассмотрим аналогичную механику только через `range`.

Обход выглядит просто и практически схожим образом, за исключением того, что `range` сам проверяет, что канал закрыт и
все сообщения уже получены и сам остановит свою работу:

```
    channel := make(chan int, 2)

    channel <- 1
    channel <- 2

    close(channel)

    for element := range channel {
        fmt.Println(element)
    }
```

Выполнив этот код получим:

```
$ go run range-over-channels.go 
1
2
```

При этом, если убрать буферизацию и/или закрытие канала получим ошибку `deadlock`. Если вам непонятно из-за чего она 
возникла – еще раз подробно изучите предыдущий урок, в нем это подробно разбиралось.

____

В этом уроке вы узнали:

1. Как с помощью range корректно обойти канал

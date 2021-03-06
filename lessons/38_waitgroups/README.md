
# Изучаем язык Go #38 – WaitGroup

В этом уроке мы рассмотрим работу типа `sync.WaitGroup` который позволяет дожидаться окончания работы n-количества
горутин.

Для начала напишем функцию `worker()` которая будет эмулировать какую-то работу через задержку на 1 секунду:

```
func worker(id int) {
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}
```

Теперь в функции `main()` пишем основную логику. Начинаем с инициализации типа `sync.WaitGroup`:

```
    var wg sync.WaitGroup
```

Затем в цикле создаем 5 горутин, в которых будет запускаться воркер:

```
    for i := 1; i <= 5; i++ {

        wg.Add(1)

        i := i

        go func() {
            defer wg.Done()
            worker(i)
        }()

    }
```

Многое из этого кода может быть непонятным. Рассмотрим подробнее, что происходит.

```
        wg.Add(1)
```

Через метод `Add()` мы указываем `sync.WaitGroup` работу скольких горутин нужно будет ожидать. На каждой итерации цикла
мы запускаем одну горутину, соответственно цифру 1 указываем в `Add()`.

```
        i := i
```

Механика этой непонятной на первый взгляд строки подробно описана 
[здесь](https://go.dev/doc/faq#closures_and_goroutines). Если коротко – это пересоздание переменной `i` нужно для 
корректной работы [замыканий](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/15_closures) и 
[горутин](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/25_goroutines). Попробуйте удалить её и посмотреть 
как отработает код.

Рассмотрим последний блок в цикле:

```
        go func() {
            defer wg.Done()
            worker(i)
        }()
```

Здесь все должно быть понятно (функция `worker()` запускается в горутине) кроме ключевого слова `defer` – команда,
перед которой указано это ключевое слово будет выполнена последней в функции. Т.е. несмотря на то, что вначале идет
`defer wg.Done()`, а потом `worker(i)` очередность выполнения будет наоборот.

И в завершение, уже после цикла, нужно указать:

```
    wg.Wait()
```

Этой командой мы указываем `sync.WaitGroup`, что нужно дождаться завершения всех горутин.

Выполнив код программы получим:

```
$ go run waitgroups.go 
Worker 5 started
Worker 1 started
Worker 4 started
Worker 2 started
Worker 3 started
Worker 2 done
Worker 5 done
Worker 3 done
Worker 4 done
Worker 1 done
```

____

В этом уроке вы узнали:

1. Как использовать `sync.WaitGroup`

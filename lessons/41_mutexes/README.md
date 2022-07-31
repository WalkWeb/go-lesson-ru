
# Изучаем язык Go #41 – Mutexes

В [прошлом уроке](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/40_atomic_counters) мы рассмотрели как 
управлять состоянием счетчика через из нескольких горутин через `sync/atomic`. Но это лишь один из вариантов. Другим, 
более безопасным вариантом является вариант через `sync.Mutex`.

В этом примере мы также пойдем чуть дальше и сделаем сразу несколько счетчиков. Все они будут храниться в одной 
структуре, но в разных строковых ключах:

```
type Container struct {
    mutex    sync.Mutex
    counters map[string]int
}
```

Функция, которая будет увеличивать счетчик на единицу:

```
func (c *Container) inc(name string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.counters[name]++
}
```

Обратите внимание, что вначале мы блокируем мьютекс `mutex.Lock()`, затем увеличиваем его на 1, и затем разблокируем 
`mutex.Unlock()`, за счет механики блокировки и разблокировки гарантируется атомарность изменения которая не будет 
затронута другими горутинами. Также напоминаем, что операции отмеченные ключевым словом `defer` будут выполняться 
последними в функции.

Теперь переходим к функции `main()`. Вначале создаем контейнер:

```
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }
```

Затем создаем `sync.WaitGroup` с помощью которого будем отслеживать завершение работы горутин:

```
    var wg sync.WaitGroup
```

Создаем замыкание, в котором будет вызываться метод, увеличивающий счетчик 10 000 раз:

```
    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }
```

И собственно сам вызов:

```
    wg.Add(3)

    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)

    wg.Wait()

    fmt.Println(c.counters)
```

Контейнер со счетчиками будет увеличиваться одновременно 3 горутинами, при этом две из них будут увеличивать один и тот 
же счетчик.

Выполнив этот код получим:

```
$ go run mutexes.go 
map[a:20000 b:10000]

```

____

В этом уроке вы узнали:

1. Как организовать атомарность операций, при одновременной работе нескольких горутин с одними и теми же данными через `sync.Mutex`
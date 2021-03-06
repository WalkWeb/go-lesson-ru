
# Изучаем язык Go #40 – sync/atomic

Прежде чем переходить к уроку нужно освятить два момента:

- Атомарная операция — операция, которая либо выполняется целиком, либо не выполняется вовсе; операция, которая не может быть частично выполнена и частично не выполнена (с) Википедия
- Организовать целостность данных при одновременной работе с данными нескольких десятков, сотен и тысяч процессов – одна из сложнейших задач в программировании. 
- Пакет `sync/atomic` позволяет совершать низкоуровневые атомарные операции с памятью для синхронизации алгоритмов. Пакет требует особой осторожности в работе и не рекомендуется в использовании там, где без него можно обойтись

Основным способом управления состояния в Golang является связь через 
[каналы](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/26_channels). В то же время есть и другие способы
управлениям состоянием. В этом уроке мы рассмотрим использование пакета `sync/atomic` для реализации атомарного 
счетчика, к которому обращается несколько [горутин](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/25_goroutines).

Для хранения значения счетчика мы будем использовать беззнаковый (префикс `u`) int размером 64 бита: 

```
    var ops uint64
```

Для ожидания завершения работы всех горутин используем пакет 
[sync.WaitGroup](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/38_waitgroups):

```
    var wg sync.WaitGroup
```

Для поверки корректности нашего решения мы создадим 50 горутин каждая из которых будет 1000 раз изменять значение 
счетчика на +1:

```
    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {
                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }
```

Обратите внимание, что счетчик изменяется через специальный метод `AddUint64()` при этом ему передается не значение
счетчика, а ссылка на область памяти и значение, на которое необходимо увеличить хранящееся там число.

После цикла нам остается дождаться, пока все горутины отработают и вывести получившееся значение счетчика:

```
    wg.Wait()

    fmt.Println("ops:", ops)
```

Выполнив код получим:

```
$ go run atomic-counters.go 
ops: 50000
```

____

В этом уроке вы узнали:

1. Как выполнять низкоуровневые операции с памятью через `sync/atomic`

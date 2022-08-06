
# Изучаем язык Go #69 – Тестирование и бенчмаркинг

Важной частью хорошего кода и проекта являются авто-тесты. В этом уроке мы рассмотрим примеры, как их писать.

Рабочий код и код с авто-тестами принято хранить в разных файлах. По этому для начала создадим файл `intutils.go` и
напишем там код, который будет реализовывать какую-то логику. Например, выбирать из двух чисел наименьшее:

```

package main

import (
    "fmt"
)

func MinInt(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {

    fmt.Println(MinInt(10, 5))

}
```

Функция `main()` для тестирования не нужна, но в ней мы, условно говоря, проверяем код вручную – передавая нужные
параметры, а потом, запуская программу смотрим, что получили именно то, что ожидали:

```
$ go run intutils.go 
5
```

Но как раз для того, чтобы каждый раз не проверять весь функционал руками придумали авто-тестирование.

Создаем отдельный файл `intutils_test.go`. В Go принято писать тесты в тех же файлах, что и тестируемый код, добавляя 
суффикс `_test`.

В этом файле создадим вначале простой тест, делающий одну проверку:

```
func TestMinIntBasic(t *testing.T) {
    result := MinInt(2, -2)

    if result != -2 {
        t.Errorf("MinInt(2, -2) = %d; expected -2", result)
    }
}
```

Запустим наш тест, для этого выполним команду `go test -v`:

```
$ go test -v
=== RUN   TestMinIntBasic
--- PASS: TestMinIntBasic (0.00s)
PASS
ok      _/var/www/github_go/lessons/69_testing  0.001s
```

Тест успешно пройден.

Но часто бывает недостаточно просто сделать одну проверку. Чтобы протестировать функционал по-хорошему – нужно проверить
его с самыми разными вариантами данных, при этом, чем более необычные и неочевидные (но подходящие) данные прогоняются
в тестах – тем лучше. Так мы будем уверены, что программа не сломается, если из внешнего источника придут какие-то
данные, о варианте которых программист не подумал.

Чтобы написать тест одной функции/метода с разными вариантами данных, напишем:

```
func TestMinIntTableDriven(t *testing.T) {
    var tests = []struct{
        a, b     int
        expected int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {3, 2, 2},
        {0, -1, -1},
    }

    for _, test := range tests {

        testname := fmt.Sprintf("%d,%d", test.a, test.b)
        t.Run(testname, func(t *testing.T) {

            result := MinInt(test.a, test.b)

            if result != test.expected {
                t.Errorf("got %d, want %d", result, test.expected)
            }

        })

    }
}
```

Выполнив тесты еще раз получим:

```
$ go test -v
=== RUN   TestMinIntBasic
--- PASS: TestMinIntBasic (0.00s)
=== RUN   TestMinIntTableDriven
=== RUN   TestMinIntTableDriven/0,1
=== RUN   TestMinIntTableDriven/1,0
=== RUN   TestMinIntTableDriven/3,2
=== RUN   TestMinIntTableDriven/0,-1
--- PASS: TestMinIntTableDriven (0.00s)
    --- PASS: TestMinIntTableDriven/0,1 (0.00s)
    --- PASS: TestMinIntTableDriven/1,0 (0.00s)
    --- PASS: TestMinIntTableDriven/3,2 (0.00s)
    --- PASS: TestMinIntTableDriven/0,-1 (0.00s)
PASS
ok      _/var/www/github_go/lessons/69_testing  0.002s
```

В завершение рассмотрим нагрузочное тестирование (или тестирование производительности). Добавим еще код в 
`intutils_test.go`:

```
func BenchmarkMinInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MinInt(1, 2)
    }
}
```

И выполним команду `go test -bench=.`:

```
$ go test -bench=.
goos: linux
goarch: amd64
BenchmarkMinInt-12      1000000000               0.258 ns/op
PASS
ok      _/var/www/github_go/lessons/69_testing  0.293s
```

Что произошло: Go запустил написанный тест производительности `BenchmarkMinInt` и ответил, что запустил проверку 
1000000000 раз, при этом на одну итерацию в среднем потребовалось 0.258 наносекунд.

Но это не все, еще можно протестировать потребление памяти и количество аллокаций памяти. Для этого нужно добавить флаг
`-benchmem`:

```
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
BenchmarkMinInt-12      1000000000               0.265 ns/op           0 B/op          0 allocs/op
PASS
ok      _/var/www/github_go/lessons/69_testing  0.303s
```

В дополнение ко времени на одну итерацию мы видим количество потребляемой памяти `0 B/op` и аллокаций памяти 
`0 allocs/op`. Все по нулям, потому что тестируемая функция очень простая.

В реальном проекте такие проверки будут очень полезны. При этом, за счет использования git вы можете откатываться к
предыдущим версиям кода, и смотреть, как менялась производительность программы по мере её изменения и усложнения.

____

В этом уроке вы узнали:

- Как покрывать код Go авто-тестами

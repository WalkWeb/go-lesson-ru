
# Изучаем язык Go #4 - Константы

Константы, как и [переменные](https://github.com/WalkWeb/go-lesson-ru/tree/master/3_variables) хранят какие-то данные,
но в отличие от переменных константы после своего объявления уже не могут быть изменены в процессе выполнения программы.

Константы могут быть объявлены там же, где и переменные: глобально или внутри функций

Для начала, как и раньше, создаем новый файл `constants.go` и заполняем его базовым кодом:

```

package main

import (
    "fmt"
)

func main() {

}

```

После этого, перед функцией `main` объявим константу:

```
const s string = "constant"
```

А внутри функции, её вывод:

```
    fmt.Println(s)
```

Выполнив программу получим:

```
$ go run constants.go 
constant
```

Все как и в случае работы с переменными.

Далее внутри функции `main` объявим новую две новых константы, при этом вторая будет рассчитываться из первой:

```
    const n = 500000000
    const d = 3e20 / n
```

Выполнив получим:

```
$ go run constants.go 
constant
6e+11
```

Числовая константа не будет иметь типа до тех пор, пока он не будет задан, например, явным преобразованием. Для 
демонстрации добавим строку:

```
fmt.Println(int64(d))
```

И выполним команду еще раз, получим:

```
$ go run constants.go 
constant
6e+11
600000000000
```

Такая механика позволяет присвоить константе тот тип, который нужен в данном контексте, например, при использовании в
`math.Sin()` константа будет приведена к типу float64.

Для иллюстрации этого добавляем еще две строки - в импорт (`import (...)`) добавляем `"math"`, а в функции строку:

```
fmt.Println(math.Sin(n))
```

Константы можно объявлять сразу группами, используя два варианта синтаксиса:

```
    const (
        pi float64 = 3.1415
        e float64 = 2.7182
    )

    const k, l = 20, 2000

    fmt.Println(pi)
    fmt.Println(e)
    fmt.Println(k)
    fmt.Println(l)
```

Если несколько констант должны иметь одно значение, можно указать так:

```
    const (
        a = 1
        b
        c
    )

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
```

В этом случае константа `a` будет равна `1`, а все остальные приравнены к `a`.

При этом нужно помнить, что константа не может приравниваться к переменной. Такой код приведет к ошибке:

```

    var p = "page"
    const badConst = p // Error
```

При попытке выполнения получим ошибку:

```
$ go run constants.go 
# command-line-arguments
./constants.go:47:11: const initializer p is not a constant
```
____

В этом уроке вы узнали:

1. Различные способы объявления констант

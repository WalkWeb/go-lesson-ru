
# Изучаем язык Go #11 - Range

Range - это специальное ключевое слово, с помощью которого можно делать обход 
[массивов](https://github.com/WalkWeb/go-lesson-ru/tree/master/8_arrays), 
[срезов](https://github.com/WalkWeb/go-lesson-ru/tree/master/9_slices), 
[карт](https://github.com/WalkWeb/go-lesson-ru/tree/master/10_maps), и channel. При этом с помощью range можно сделать
обход и строки, потому что строка в go это срез байтов доступных только для чтения.

Рассмотрим пример обхода среза с помощью range:

```
    nums := []int{2, 3, 4}
    sum := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum) // sum: 9
```

Range возвращает как индексы так и значения итерируемой структуры. В примере выше индексы нам были не нужны, и мы их
проигнорировали с помощью символа нижнего подчеркивания `_`, а в примере ниже они используются:

```
    for index, num := range nums {
        fmt.Println("index:", index, "num:", num)
    }
```

Выполнив этот код получим:

```
$ go run range.go 
// ...
index: 0 num: 2
index: 1 num: 3
index: 2 num: 4
```

Рассмотрим пример обхода карты с помощью range:

```
    languages := map[string]string{"p": "php", "g": "golang"}
    for k, v := range languages {
        fmt.Printf("%s -> %s\n", k, v)
    }
```

Выполнив этот код получим:

```
$ go run range.go 
// ...
g -> golang
p -> php
```

В обходе можно задействовать только ключи:

```
    for k := range languages {
        fmt.Println("key:", k)
    }
```

Выполнив этот код получим:

```
$ go run range.go 
// ...
key: p
key: g
```

В заключение рассмотрим экзотический вариант использования range для обхода строки (которая на самом деле является 
срезом):

```
    for i, c := range "golang" {
        fmt.Println(i, c)
    }
```

Выполнив этот код получим:

```
$ go run range.go 
// ...
0 103
1 111
2 108
3 97
4 110
5 103
```

Полученный результат тяжело описать в двух словах, если вам стало интересно почему получился такой результат почитайте 
самостоятельно про строки в go.
____


В этом уроке вы узнали:

1. Что такое range и как его использовать


# Изучаем язык Go #43 – Сортировка

Иногда нужно отсортировать что-либо по порядку. Например срез из чисел. Для этого в Golang существует пакет `sort`, 
который предоставляет все нужны методы для сортировки.

Рассмотрим несколько примеров.

Сортировка строкового среза делается так:

```
    messages := []string{"c", "a", "b"}
    sort.Strings(messages)
    fmt.Println("Messages:", messages)
```

Выполнив этот код получим:

```
$ go run sorting.go 
Messages: [a b c]
```

Сортировка числового среза делается аналогично, только через другой метод:

```
    integers := []int{10, 4, 6}
    sort.Ints(integers)
    fmt.Println("Integers:", integers)
```

Выполнив этот код получим:

```
$ go run sorting.go 
Messages: [a b c]
Integers: [4 6 10]
```

Также можно проверить, например, срез, отсортирован ли он уже:

```
    isSorted := sort.IntsAreSorted(integers)
    fmt.Println("Is sorted:", isSorted)
```

Выполнив этот код получим:

```
$ go run sorting.go 
Messages: [a b c]
Integers: [4 6 10]
Is sorted: true
```

Посмотреть все методы пакета `sort` можно в [официальной документации](https://pkg.go.dev/sort)

____

В этом уроке вы узнали:

1. Как делать сортировку с помощью встроенного пакета `sort`

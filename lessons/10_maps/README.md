
# Изучаем язык Go #10 - Карты

Карты - это специальная структура данных, которую обычно называю ассоциативными массивами.

Карта создается следующим образом:

```
    m := make(map[string]int)
```

В квадратных скобках указывается, каким типом будет индекс (ключи ассоциативного массива). В данном случае это будут 
строки.

Запишем в карту два значения и посмотрим что получится:

```
    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m) // map[k1:7 k2:13]
    fmt.Println("map [k1]:", m["k1"]) // map [k1]: 7
```

Все просто и интуитивно понятно.

Так же как и с [массивами](https://github.com/WalkWeb/go-lesson-ru/tree/master/8_arrays) и 
[срезами](https://github.com/WalkWeb/go-lesson-ru/tree/master/9_slices) длину карты (т.е. количество записей в ней) 
можно узнать с помощью функции `len()`:

```
    fmt.Println("len:", len(m)) // len: 2
```

Можно удалить одно из значений в карте:

```
    delete(m, "k2")
    fmt.Println("len:", len(m)) // len: 1
```

И замерив размер карты еще раз увидим, что теперь в ней один элемент ("k1")

Обращение к несуществующему ключу вернет имя ключа `0` со значением `false`:

```
    key, value := m["k3"]
    fmt.Println("key:", key, "value", value) // key: 0 value false
```

Примечательно, что никакой ошибки не происходит.

В заключение рассмотрим вариант как создать и заполнить карту сразу:

```
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n) // map: map[bar:2 foo:1]
```
____


В этом уроке вы узнали:

1. Как работать с картами

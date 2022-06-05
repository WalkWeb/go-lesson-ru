
# Изучаем язык Go #19 – Структуры

Структуры в Go позволяют группировать данные, например, относящиеся к одной логической сущности.

Объявим структуру `person` с двумя полями: имя и возраст:

```
type person struct {
    name string
    age  int
}
```

И создадим её:

```
    fmt.Println(person{"Bob", 20}) // {Bob 20}
```

Соблюдать порядок передаваемых значений не обязательно, но для этого нужно отдельно указывать имя каждого параметра:

```
    fmt.Println(person{age: 30, name: "Alice"}) // {Alice 30}
```

Какое-то из полей может быть пропущено, в этом случае в него будет записано `nil`-значение для данного типа данных. Для
`integer` это 0:

```
    fmt.Println(person{name: "Fred"}) // {Fred 0}
```

Как и с [переменными](https://github.com/WalkWeb/go-lesson-ru/tree/master/3_variables), можно передать 
[указатель](https://github.com/WalkWeb/go-lesson-ru/tree/master/17_pointers) на структуру используя специальный символ
`&`:

```
    fmt.Println(&person{"Anna", 25}) // &{Anna 25}
```

Обычно ответственность создания структуры назначают какой-то 
[функции](https://github.com/WalkWeb/go-lesson-ru/tree/master/12_functions):

```
func createPerson(name string, age int) *person {
    return &person {name, age}
}
```
```
    fmt.Println(createPerson("Olga", 40)) // &{Olga 40}
```

Доступ к отдельным полям структуры осуществляется через точку:

```
    p := person{"Nikolai", 33}
    fmt.Println(p.name) // Nikolai
```

Данные в структурах можно изменять:

```
    p.age = 50
    fmt.Println(p.age)
```
____

В этом уроке вы узнали:

1. Как создавать и работать со структурами

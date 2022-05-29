
# Изучаем язык Go #22 – Встраивание структур

Прежде чем переходить к самому встраиванию структур друг в друга (в этом в общем-то нет ничего сложного), нужно 
объяснить, где это используется. Представим, что нам нужно создать структуры под два объекта: машину и мотоцикл. При этом
в каждой структуре нам нужно: указать параметры двигателя, производителя руля и материал изготовления кресла плюс 
какие-то отдельные параметры только для машины, и отдельные только для мотоцикла.

Можно сделать полностью отдельные структуры для машины и мотоцикла, но если немного подумать над задачей, становится
очевидно, что параметры двигателя, руля и кресел будут одинаковыми, и чтобы их не дублировать - лучше вынести в 
отдельные сущности, и уже машину и мотоцикл собирать из них + добавить те уникальные параметры, которые необходимы для 
них.

Теперь, когда стало понятно, где используется встраивание структур, переходим к коду. Чтобы не запутать читателя обилием
структур, параметров и кода, рассмотрим встраивание на двух простых сущностях, никак не связанных с примером выше.

Создадим структуру возраста с одним параметром, указывающую собственно возраст:

```
type age struct {
    amount int
}
```

И структуру персоны, содержащую в себе имя и структуру возраста:

```
type person struct {
    name string
    age
}
```

Как видно, встраивание одной структуры в другую производится элементарно – мы просто указываем её название в списке.

Теперь создадим объекты. Их можно создать двумя вариантами:

```
    a := age{30}
    p := person{"Max", a}

    fmt.Println(p)

    p2 := person{
        "Maria",
        age{26},
    }

    fmt.Println(p2)
```

В первом случае мы структуру возраста вначале присвоили отдельной переменной, и потом использовали её, во втором случае 
вся структура создается с нуля.

Обратите внимание, что запятая в первом случае после последнего параметра не пишется, хотя её можно и добавить
`person{"Max", a,}` и это ничего не меняет. А во втором случае запятая после последнего параметра обязательна.

Выполнив этот код получим:

```
$ go run struct_embedding.go 
{Max {30}}
{Maria {26}}
```

Обращение к параметрам структуры, как мы уже рассматривали ранее, происходит через точку:

```
    fmt.Printf("name: %v, age: %v\n", p.name, p.amount)

    fmt.Println("also age:", p.age.amount)
```

Обратите внимание, что в первом случае Go сам находит amount внутри дочерней структуры. И только во втором случае мы 
руками указываем полный путь к параметру.

Выполнив этот код получаем:

```
$ go run struct_embedding.go 
...
name: Max, age: 30
also age: 30
```

Теперь добавим метод `tellAge()` структуре `age` и интерфейс `speaker`:

```
func (a age) tellAge() string {
    return fmt.Sprintf("my age %v", a.amount)
}

type speaker interface {
    tellAge() string
}
```

И вызовем его. Обратите внимание, что как и в случае с параметром `amount`, мы не указываем полный путь до метода, 
оставляя эту задачу для Go:

```
    fmt.Println("p.tellAge():", p.tellAge()) // p.tellAge(): my age 30
```

Это в свою очередь позволяет интерфейс дочерней структуры присваивать родительской структуре. Это наглядно видно на 
примере ниже, когда структура `person` присваивается переменной `s` реализующей интерфейс `speaker`:

```
    var s speaker = p2
    fmt.Println("speaker:", s.tellAge()) // speaker: my age 26
```

____

В этом уроке вы узнали:

1. Как встраивать одни структуры в другие
2. Встроенную механику поиска нужного параметра/метода в дочерних структурах, когда не обязательно указывать полный путь
   до неё
3. Как родительская структура автоматически реализует интерфейсы, которые реализуют её дочерние структуры
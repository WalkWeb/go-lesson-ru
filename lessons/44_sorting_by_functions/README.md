
# Изучаем язык Go #44 – Сортировка по функциям

Помимо базовой сортировки, которую мы рассмотрели в [предыдущем уроке](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/43_sorting), 
иногда нужно сделать сортировку по какой-то особой логике. Например отсортировать строки не по алфавиту, а по длине слов.

Для создания такой сортировки нам нужен соответствующий тип, реализующий интерфейс [sort.Interface](https://pkg.go.dev/sort#Interface)

Создаем тип `byLength`, который будет аналогом `[]string`:

```
type byLength []string
```

Добавляем ему необходимые методы для реализации интерфейса `sort.Interface`, а именно: `Len() int`, 
`Less(i, j int) bool`, `Swap(i, j int)`:

```
func (s byLength) Len() int {
    return len(s)
}

func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
```

Основная логика, как видно, находится в методе `Less(i, j int) bool`, в котором мы сравниваем две строки по их длине.

Теперь осталось применить созданную нами сортировку:

```
func main() {

    messages := []string{"aaa", "cc", "b"}
    sort.Sort(byLength(messages))
    fmt.Println(messages)

}
```

Выполнив этот код получим:

```
$ go run sorting-by-functions.go 
[b cc aaa]
```

Результат получился таким, какой и ожидался: строки отсортированы по их длине (вначале короткие, затем длинные), а не по 
алфавитному порядку.

____

В этом уроке вы узнали:

- Как делать сортировку по нужной вам логике

Читайте также:

- Официальную документацию по пакету [sort](https://pkg.go.dev/sort)


# Изучаем язык Go #48 – Строковые функции 

В этом уроке мы рассмотрим основные строковые функции, которые содержит пакет `strings`

Для более удобной работы сделаем сокращения для пакета и функции `fmt.Println`:

```
import (
    "fmt"
    s "strings"
)

var p = fmt.Println
```

Теперь переходим к основным строковым функциям:

### strings.Contains(s, substr string) bool

Проверяет наличие `substr` в строке `s`:

```
    p("Contains:", s.Contains("test", "es"))
```

Вернет `true` – `es` присутствует в строке `test`. Если быть точнее вернет `Contains:  true`, но подразумевается именно
результат функции.

### strings.Count(s, substr string) int

Считает, сколько раз `substr` встречается в строке `s`:

```
    p("Count:", s.Count("test", "t"))
```

Вернет `2` – строка `t` встречается в строке `test` два раза

### strings.HasPrefix(s, prefix string) bool

Проверяет, начинается ли строка `s` с `prefix`:

```
    p("HasPrefix:", s.HasPrefix("test", "te"))
```

Вернет `true` – строка `test` начинается с `te`

### strings.HasSuffix(s, suffix string) bool

Проверяет, заканчивается ли строка `s` на `suffix`:

```
    p("HasSuffix:", s.HasSuffix("test", "xxx"))
```

Вернет `false` – строка `test` не заканчивается на `xxx`

### strings.Index(s, substr string) int

Возвращает номер первого совпадения `substr` в строке `s`. Если строка не найдена вернет `-1`:

```
    p("Index:", s.Index("task", "a"))
```

Вернет `1` (счет идет с 0). Если указать `x` или `aa` вернется `-1` – совпадения не найдено.

### strings.Join(elems []string, sep string) string

Объединяет строки в срезе `elems`, добавляя между каждым строку `sep`:

```
    p("Join:", s.Join([]string{"a", "b", "c"}, "---"))
```

Вернет `a---b---c`

### strings.Repeat(s string, count int) string

Повторяет строку `s` `count`-раз:

```
    p("Repeat:   ", s.Repeat("a", 5))
```

Вернет: `aaaaa`

### strings.Replace(s, old, new string, n int) string

Поиск и замена символов в строке. Заменяет строку `old` в строке `s` на строку `new` `n`-раз. Если в `n` передать `-1` 
то будет заменены все нахождения:

```
    p("Replace:  ", s.Replace("xxaaxx", "a", "1", 1))
```

Вернет: `xx1axx` – строка `a` была заменена на строку `1` в `xxaaxx` при одном нахождении.

### strings.Split(s, sep string) []string

Разбивает строку `s` на срез из строк по символу `sep`:

```
    p("Split:    ", s.Split("a-b-c", "-"))
```

Вернет: `[a b c]`

### strings.ToLower(s string) string

Приводит все символы переданной строки `s` в нижний регистр:

```
    p("ToLower:  ", s.ToLower("Name"))
```

Вернет `name` – т.е. большая `N` заменена на маленькую `n`

### strings.ToUpper(s string) string

Аналогично `ToLower()` только наоборот – приводит все символы переданной строки `s` в верхний регистр:

```
    p("ToUpper:  ", s.ToUpper("Name"))
```

Вернет: `NAME`


Мы рассмотрели основные, наиболее часто использующиеся функции для работы со строками. Полный список функций смотрите в
[официальной документации](https://pkg.go.dev/strings)

____

В этом уроке вы узнали:

- Основные строковые функции

Читайте также:

- Официальную документацию по пакету [strings](https://pkg.go.dev/strings#Contains)


# Изучаем язык Go #46 – Defer

Defer – это специальное ключевое слово в Golang, которым отмечается инструкция, которая должна выполниться в функции
последней.

Для наглядности напишем код:

```
    defer fmt.Println("Defer")
    fmt.Println("Hello")
```

Выполнив его получим:

```
$ go run defer.go 
Hello
Defer
```

Также отмеченная defer инструкция будет выполнена перед выполнением [panic](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/45_panic).

Несмотря на кажущуюся простоту, понять очередность вызовов может оказаться не так просто. Напишите такой код и 
попробуйте перед выполнением понять, какой будет очередность вызовов:

```
func doPanic() {
    defer fmt.Println("doPanic(): Defer") 
    panic("doPanic(): panic")             
}

func main() {

    defer fmt.Println("Defer")             
    fmt.Println("Hello")                   

    doPanic()
}
```

Правильной очередность будет такой:

```
func doPanic() {
    defer fmt.Println("doPanic(): Defer") // 2
    panic("doPanic(): panic")             // 4
}

func main() {

    defer fmt.Println("Defer")             // 3
    fmt.Println("Hello")                   // 1

    doPanic()
}
```

Теперь рассмотрим более практичный пример. Представим классическую работу с файлом, который нужно создать и что-то в
него записать. Чем нужно завершить такую операцию, независимо от того, будет ли операция завершена успешно, или 
сломается на середине? Закрыть работу с файлом.

Напишем такую программу.

Во-первых, нам нужны три отдельных функции, которые будут 1) создавать файл 2) записывать какое-то содержимое в файл 3)
закрывать работу с файлом

```
func createFile(filePath string) *os.File {
    fmt.Println("Create file:", filePath)
    file, err := os.Create(filePath)
    if err != nil {
        panic(err)
    }
    return file
}

func writeFile(file *os.File, content string) {
    fmt.Println("Write file:", content)
    _, err := file.WriteString(content)
    if err != nil {
        panic(err)
    }
}

func closeFile(file *os.File) {
    fmt.Println("Close file")
    err := file.Close()
    if err != nil {
        panic(err)
    }
}
```

Теперь создадим и запишем какое-то содержимое в файл с помощью этих функций, при этом закрытие функции отметим ключевым
словом `defer`:

```
func main() {
    file := createFile("/tmp/defer")
    defer closeFile(file)
    writeFile(file, "file content\n")
}
```

Выполним написанный код:

```
$ go run defer.go 
Create file: /tmp/defer
Write file: file content

Close file
```

Чтобы проверить, что файл действительно был создан и содержимое там действительно появилось, выполните `cat /tmp/defer`,
вы должны увидеть такой результат:

```
$ cat /tmp/defer
file content
```

Также нужно отдельно обговорить, почему `defer` идет второй строкой: указать `defer closeFile(file)` первой строкой мы
не можем, потому что файла вначале еще не существует. А указать третьей также не можем, потому что в этом случае, если
при записи в файл произойдет ошибка, программа не дойдет до указаний, написанных в `defer closeFile(file)`.

Проверить это легко, замените содержимое функции `writeFile()`:

```
func writeFile(file *os.File, content string) {
    panic("writeFile panic!")
}
```

И поставьте `defer closeFile(file)` последней инструкцией:

```
func main() {
    file := createFile("/tmp/defer")
    writeFile(file, "file content\n")
    defer closeFile(file)
}
```

Выполнив этот код получим:

```
$ go run defer.go 
Create file: /tmp/defer
panic: writeFile panic!

goroutine 1 [running]:
main.writeFile(0xc0000ba020, 0x4be337, 0xd)
        /var/www/github_go/lessons/46_defer/defer.go:19 +0x39
main.main()
        /var/www/github_go/lessons/46_defer/defer.go:37 +0x6f
exit status 2
```

Т.е. до вызова инструкции `defer closeFile(file)` код не дошел. Теперь вернем `defer closeFile(file)` на вторую строку и
выполним программу еще раз:

```
$ go run defer.go 
Create file: /tmp/defer
Close file
panic: writeFile panic!

goroutine 1 [running]:
main.writeFile(...)
        /var/www/github_go/lessons/46_defer/defer.go:19
main.main()
        /var/www/github_go/lessons/46_defer/defer.go:33 +0x7f
exit status 2
```

На этот раз закрытие файла отработало. Момент с расположения инструкции с `defer` важно понимать. Обычно её делают как
можно ближе к началу тела функции, в идеале первой.

____

В этом уроке вы узнали:

- Как работает Defer в Golang

Читайте также:

- Официальную документацию по [Defer](https://go.dev/blog/defer-panic-and-recover)

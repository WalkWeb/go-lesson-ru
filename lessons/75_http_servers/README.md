
# Изучаем язык Go #75 – HTTP-серверы

Как и в случае с [http-клиентом](https://github.com/WalkWeb/go-lesson-ru/tree/master/lessons/74_http_client) для 
написания http-сервер используется базовая библиотека `net/http`

Основной концепцией net/http-серверов являются хандлеры. Хандлер – это объект реализующий интерфейс `http.Handler`,
обычно для этого пишется функция с соответствующими параметрами.

Пример такой функции, которая на http-запрос вернет ответ с телом `hello`. Она обязательно должна иметь в принимаемых 
аргументах `writer http.ResponseWriter, request *http.Request`. `http.ResponseWriter` используется для записи ответа:

```go
func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello\n")
}
```

Напишем еще одну функцию, которая в ответе напишет, с какими заголовками был сделан запрос:

```go
func headers(writer http.ResponseWriter, request *http.Request) {
	for name, headers := range request.Header {
		for _, header := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, header)
		}
	}
}
```

Пришло время создать сам http-сервер и зарегистрировать функции-хандлеры, которые будут обрабатывать http-запросы:

```go
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		fmt.Println("Listen error:", err)
	}
```

`8090` – это номер порта, по которому наш http-сервер будет слушать запросы.

Теперь пришло время проверить его в работе. Вначале запускаем его *в фоне* с помощью команды:

```
$ go run http-servers.go &
[1] 13018
```

Цифра `13018` – это ID запущенного процесса, по которому его можно «убить», т.е. остановить. В вашем случае цифра будет 
другой.

Теперь сделаем http-запрос на метод `/hello` с помощью `curl` и проверим ответ:

```
$ curl localhost:8090/hello
hello
```

Как и ожидалось, получен ответ `hello`. Теперь посмотрим ответ на запрос на метод `/headers`:

```
$ curl localhost:8090/headers
User-Agent: curl/7.70.0
Accept: */*
```

Так как запрос делается через curl – количество заголовков будет минимальным. Вы можете ради интереса сделать запрос
через браузер, набрав в адресе `localhost:8090/headers`. В моем случае ответ получился таким:

```
Sec-Ch-Ua-Platform: "Linux"
Upgrade-Insecure-Requests: 1
Sec-Fetch-Dest: document
User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36
Sec-Fetch-Mode: navigate
Sec-Ch-Ua-Mobile: ?0
Sec-Fetch-Site: none
Accept-Encoding: gzip, deflate, br
Accept-Language: ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,lb;q=0.6,uk;q=0.5
Connection: keep-alive
Sec-Ch-Ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-Fetch-User: ?1
```

В завершение «убьем» запущенный http-сервер по ID процесса:

```
$ sudo kill 13018
[1]+  Завершено      go run http-servers.go
```

Потому что если не освободить его, и попробовать запустить http-сервер еще раз, получим ошибку:

```
$ go run http-servers.go 
Listen error: listen tcp :8090: bind: address already in use
```

Это происходит из-за того, что один порт может слушать (т.е. занимать) только один процесс.

____

В этом уроке вы узнали:

- Как сделать http-сервер

Читайте также:

- [Официальную документацию по пакету net/http](https://pkg.go.dev/net/http)

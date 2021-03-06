
# Изучаем язык Go #37 – Рабочие пулы

Преимущество Golang перед другими языками – это простая реализация быстрой и параллельной обработки каких-то данных. И
в этом уроке мы рассмотрим конкретный пример как это сделать.

Начнем с функции `worker()`. Вообще под воркерами подразумевают какие-то процессы, которые запущены на сервере (часто в 
отдельном docker-контейнере) и что-то делают. Например, слушают данные в канале RabbitMQ, делают с полученными данными 
необходимые манипуляции и отправляют результат в другой канал RabbitMQ.

Для эмуляции такого воркера пишем функцию с простой задержкой на 1 секунду (эмулируем какую-то обработку), и выводом 
отладочной информации, чтобы было понятно, когда и какой воркер начал и закончил работу:

```
func worker(id int, jobs <-chan int, result chan<- string) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- "abstract worker result"
	}
}
```

А теперь параллельная обработка произвольного количества данных на произвольном количестве процессов:

```
func main() {

	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

    close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
```

Выполнив этот код получим:

```
$ go run worker-pools.go 
worker 3 started job 1
worker 2 started job 3
worker 1 started job 2
worker 3 finished job 1
worker 3 started job 4
worker 1 finished job 2
worker 2 finished job 3
worker 1 started job 5
worker 3 finished job 4
worker 1 finished job 5
```

Рекомендую поменять разные значения в константах `numJobs` и `numWorkers` и посмотреть как будет менять работа 
программы.

Код программы не должен требовать каких-то отдельных пояснений, единственное что стоит отдельно отметить, это строчку 
`close(jobs)` – все будет работать и без неё, а закрывается канал исключительно на всякий случай, чтобы с обходом канала
через `range` не возникло каких-либо потенциальных проблем.

В заключение стоит добавить несколько слов об оптимальном количестве воркеров. Если речь идет о вычислениях нагружающих 
процессор, то количество воркеров не имеет смысла задавать больше, чем количество ядер в процессоре. А если речь идет о 
запросах к внешнему API, то в этом случае воркеров можно быть значительно больше. Тут скорее стоит заботиться о том, чтобы 
внешний API не упал под нагрузкой.

____

В этом уроке вы узнали:

1. Как организовать параллельную обработку данных в Golang

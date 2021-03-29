## Домашнее задание №11 «Клиент TELNET»

Необходимо реализовать крайне примитивный TELNET клиент
(без поддержки команд, опций и протокола в целом).

Примеры вызовов:
```bash
$ go-telnet --timeout=10s host port
$ go-telnet mysite.ru 8080
$ go-telnet --timeout=3s 1.1.1.1 123
```

* Программа должна подключаться к указанному хосту (IP или доменное имя) и порту по протоколу TCP.
* После подключения STDIN программы должен записываться в сокет,
а данные, полученные из сокета, должны выводиться в STDOUT - всё происходит конкурентно.
* Опционально в программу можно передать таймаут на подключение к серверу
(через аргумент `--timeout`) - по умолчанию `10s`.
* При нажатии `Ctrl+D` программа должна закрывать сокет и завершаться с сообщением.
* При получении `SIGINT` программа должна завершать свою работу.
* Если сокет закрылся со стороны сервера, то при следующей попытке отправить сообщение программа
должна завершаться (допускается завершать программу после "неудачной" отправки нескольких сообщений).
* При подключении к несуществующему серверу, программа должна завершаться с ошибкой соединения/таймаута.

При необходимости можно выделять дополнительные функции / ошибки.

Примеры работы:

1) Сервер обрывает соединение
```bash
$ nc -l localhost 4242
Hello from NC
I'm telnet client
Bye, client!
^C
```

```bash
$ go-telnet --timeout=5s localhost 4242
...Connected to localhost:4242
Hello from NC
I'm telnet client
Bye, client!
Bye-bye
...Connection was closed by peer
```

Здесь сообщения
```
Hello from NC
Bye, client!
```
и операция Ctrl+C (Unix) относятся к `nc`,

а сообщения
```
I'm telnet client
Bye-bye
```
относятся к `go-telnet`.

2) Клиент завершает ввод
```bash
$ go-telnet localhost 4242
...Connected to localhost:4242
I
will be
back!
^D
...EOF
```

```bash
$ nc -l localhost 4242
I
will be
back!
```

Здесь сообщения
```
I
will be
back!
```
и операция Ctrl+D (Unix) относятся к `go-telnet`,

Сообщения
```
...Connected to localhost:4242
...Connection was closed by peer
...EOF
```
являются служебными.
Они **выводятся в STDERR** и не тестируются `test.sh`, их формат на усмотрение автора.

### Критерии оценки
- Пайплайн зелёный - 4 балла
- Добавлены юнит-тесты - до 2 баллов
- Понятность и чистота кода - до 4 баллов

#### Зачёт от 7 баллов

### Подсказки
- `flag.StringVar`, `flag.DurationVar`
- `net.JoinHostPort`, `net.DialTimeout`
- `bufio` / `io`
- `signal.NotifyContext`
- https://stackoverflow.com/questions/51317968/write-on-a-closed-net-conn-but-returned-nil-error

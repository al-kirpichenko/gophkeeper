# Клиент для проекта Gophkeeper.
Клиент работавет в виде консольной утилиты, выполняя команды пользователя

```
>go run main.go
Client version: 0.01
Build: 2023-12-12

welcome to client for a gophkeeper server

Usage:
    client [command]

Available Commands:
    auth        authorization and registration commands
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    keeper      keeper commands
Flags:
    -h, --help            help for client
    -a, --server string   host addr (default "http://localhost:8080")
    Use "client [command] --help" for more information about a command.
```
### Регистрация и авторизация

Пример команды регистрации:

```
go run main.go auth register -u user -p password
```
Пример команды авторизации:

```
go run main.go auth login -u user -p password
```

### Добавление новых данных

Добавление новых данных можно выполнить с помощью команды

```
>go run main.go keeper create 

Usage:
  client keeper [command]

Available Commands:
  create      add secret command
  get         get command

Flags:
  -b, --bin string        binary
  -k, --comment string    comment
  -c, --cvv string        cvv code
  -d, --date string       date
  -h, --help              help for keeper
  -n, --num string        card number
  -p, --password string   password
  -s, --string string     string
  -t, --title string      title
  -u, --username string   username

```
Пример добавления текстовой записи:
```
go run main.go keeper create -t "my text name" -s "custom text" -k comment
```
Пример добавления учетной записи:
```
go run main.go keeper create -t google.com -u username -p password 
```
Пример добавления банковской карты:
```
go run main.go keeper create -t Сбербанк1 -n "1234 5678 9876 5432" -d 12/26 -c 111
```
### Получение данных

Получение данных осуществляется с помощью команды get
```
go run main.go keeper get -t "my secret name" 
```
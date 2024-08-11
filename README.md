# Приватное хранилище информации GophKeeper
Второй выпускной проект на курсе "Продвинутый go разработчик".
## Общие требования
GophKeeper представляет собой клиент-серверную систему, позволяющую пользователю надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.

Сервер должен реализовывать следующую бизнес-логику:
- регистрация, аутентификация и авторизация пользователей;
- хранение приватных данных;
- синхронизация данных между несколькими авторизованными клиентами одного владельца;
- передача приватных данных владельцу по запросу.

Клиент должен реализовывать следующую бизнес-логику:

- аутентификация и авторизация пользователей на удалённом сервере;
- доступ к приватным данным по запросу.

### Типы хранимой информации

- пары логин/пароль;
- произвольные текстовые данные;
- произвольные бинарные данные;
- данные банковских карт.

## Реализация

В нашем проекте мы будем использовать протокол gRPC на стримах на клиенте и сервере для реализации требования 
"синхронизация данных между несколькими авторизованными клиентами одного владельца".
Аутентификация будет простая по паре логин/пароль. Пароль будем хранить в виде хэша. При успешной аутентификации выдаем
JWT токен для решения задачи авторизации.

В системе будет реализован терминальный графический интерфейс с помощью пакета - https://github.com/rivo/tview.
В интерфейсе пользователь сможет: зарегистрироваться, зайти в аккаунт, посмотреть список доступного ему контента или загрузить новый.

### Event Storming системы "GophKeeper"
![img.png](doc/img.png)

### Таблицы бд


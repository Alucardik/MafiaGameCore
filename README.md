# Чекин Игорь, Сервис-ориентированные архитектуры, Домашнее задание 4

## Общая структура проекта

В папке `server` находятся файлы сервера, в папке `client` - клиента. На сервере реализована инфраструктура для поддержания множественных сессий и чат с поддержкой ролей, однако из-за некоторых недоработок одновременно идет только одна сессия.

Для локального запуска сервера необходимо наличие `go` версии `1.17`. Требуется перейти в корневую папку проекта и ввести команду `go run .` (можно передать порт для запуска через `--port`, по умолчанию `8080`). Локальный запуск клиента осуществляется так же из корневой папки проекта командой `go run . --mode=client`.  И сервер, и клиент можно запустить через [докер-образ](https://hub.docker.com/layers/198300674/alucardik/soa-images/mafia-core/images/sha256-1f2070f98e21cb5a5a0aedc8ab3e1a41a911d503931f9f67764629b346191d10?context=repo) (`alucardik/soa-images:mafia-core`), например через команды:

```bash
docker run -i alucardik/soa-images:mafia-core
go run . --mode=client / server
```

## Ход игры

У клиента есть набор команд (описание доступно через команду `help`). Сначала необходимо выполнить команду `connect`, ввести адрес сервера (в случае локального сервера - `localhost`)  c портом (например, `localhost:8080`) и свой ник. В случае успешного подключения вы начнете получать уведомления от сервера, и останется дождаться автоматического начала сессии (от 4 игроков, после чего есть 10 секунд на подключение других участников). После начала сессии новые игроки не могут зайти. В ходе игры вы можете использовать команду `vote`, чтобы проголосовать за убийство одного из игроков или инспекцию игрока (для роли комиссара). Чтобы получить список игроков, используйте `players`. При начале игры вам дается роль, от этого зависит, можете ли голосовать ночью (мафия или комиссар), или нет. Днем комиссар может выполнить команду `expose`, тогда сервер опубликует информацию о мафии, если комиссару удалось ее найти прошлой ночью. День заканчивается, когда все живые игроки выполнят команду `skip` (менять голос до нее можно произвольное число раз, учтен будет последний). Ночью ходят мафия и комиссар через `vote`. Также доступен чат для общения через команду `chat` (призраки не могут его использовать, а ночью сообщения отправляются только среди мафии).
# BitcoinMonitor

BitcoinMonitor - приложение для мониторинга курса биткойна.

## Описание

BitcoinMonitor - это REST API, которое позволяет получать информацию о курсе биткойна в зависимости от валюты.
Приложение использует API [CoinGecko](https://www.coingecko.com/ru/api) для получения информации о курсе биткойна.

## Функциональность

-  Получение курса биткойна в зависимости от валюты;
-  Добавление валюты в список мониторинга;
-  Удаление валюты из списка мониторинга;
-  Получение списка мониторингуемых валют;
-  Получение списка доступных валют;
-  Старт мониторинга;
-  Стоп мониторинга.

## Установка

1. Установите [Docker](https://www.docker.com/get-started);
2. Установите [Docker Compose](https://docs.docker.com/compose/install/);
3. Склонируйте репозиторий;
4. Перейдите в папку с проектом;
5. Выполните команду `docker-compose up -d`;

## Использование

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/40053615-105d0e15-2802-4036-86b4-66935449d9e8?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D40053615-105d0e15-2802-4036-86b4-66935449d9e8%26entityType%3Dcollection%26workspaceId%3Dd6c028b3-486f-4435-a8e0-ca761725bba1)

* либо импортировать в Postman файл `api/Bitcoin Monitor API.postman_collection.json` (нажмите на кнопку "Run in Postman" выше);
* либо открыть swagger документацию по адресу `http://localhost:8080/swagger/index.html`.

### Получение курса биткойна

`GET /bitcoin/price?name=<name>&timestamp=<timestamp>`

-  `name` - валюта, для которой нужно получить курс биткойна;
-  `timestamp` - время, для которого нужно получить курс биткойна.

### Добавление валюты в список мониторинга

`POST /bitcoin/add`

```
{
   "coin": "<name>"
}
```

coin - имя валюты например Bitcoin
Имена валюты которые можно добавить можно узнать с помощью команды `GET /bitcoin/available`

### Удаление валюты из списка мониторинга

`POST /bitcoin/remove`

```
{
   "coin": "<name>"
}
```

coin - имя валюты например Bitcoin
Имена валюты которые моно удалить можно узнать с помощью команды `GET /bitcoin/monitoring`

### Получение списка мониторингуемых валют

`GET /bitcoin/monitoring`

### Получение списка доступных валют

`GET /bitcoin/available`

## Настройка

### ENV переменные

-  `DB_URL` - URL для подключения к базе данных;
-  `APP_PORT` - порт, на котором будет запущено приложение;
-  `APP_HOST` - хост, на котором будет запущено приложение.
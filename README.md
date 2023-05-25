# NATS-LISTENER

## назначение

Микросервис сбора в базе данных поступающих в канал заказов. Сервис слушает канал nats-streaming, валидирует
входящие сообщения, записывает валидные в базу данных postgres. Отдает информацию о заказе по uid из кэша.
При включении сервиса в кэш поднимаются уже существующие в базе данных заказы, что ускоряет их получение.

## инструкция настройки и запуска

Для запуска сервиса требуется создать базу данных postgres. Команда для запуска
контейнера на основе образа postgres 14 в Docker.

```
docker run --name orders -p 5432:5432 -e POSTGRES_PASSWORD=111 -e POSTGRES_DB=nats_test -d postgres:14
```

После запуска контейнера требуется создать таблицу. _!!В дальнейшем структура таблицы будет улучшена!!_ Скрипт для
создания таблицы с заказами:

```sql
CREATE TABLE public.orders
(
    order_uid  varchar NULL,
    order_data varchar NULL
);
```

Для запуска сервера nats-streaming в Docker контейнере требуется ввести в терминале команду:

```
docker run --nats_orders -p 4222:4222 -p 8222:8222 -d nats-streaming
```

Чтобы запустить сервис, нужно набрать в консоли в папке с проектом команду:

```
go run ./cmd/nats-listener.go -p my_folder -f my_config
```

Если флаги не указаны, сервис будет искать файл конфигурации в пути по умолчанию: директория **configs**, файл **
config**.
Для того чтобы указать свои путь и название файла конфигурации можно добавить флаги в команду:

- **--path** или **-p** - путь до директории с файлом конфигурации
- **--config** или **-с** - название файла конфигурации без расширения

Для чтения файла конфигурации используется библиотека [spf13/viper](https://github.com/spf13/viper). Рекомендуется
выбрать
один из поддерживаемых библиотекой форматов конфигурации, например - **yml**, **json** или **toml**.
Пример:

```toml
[nats]
cluster = "test-cluster"
client = "stan-sub"
URL = "127.0.0.1:4222"
subj = "test"

[dataBase]
pgConnLink = "postgres://postgres:111@localhost:5432/nats_test"

[http]
port = "8008"
```

### пояснения к полям:

**nats**

+ **cluster** - имя сервера кластера
+ **client** - уникальное имя клиента
+ **URL** - ссылка для подключения к серверу nats-streaming
+ **subj** - название канала, на который подписывается клиент

**dataBase**

+ **pgConnLink** - ссылка для подключения к базе данных

**http**

+ **port** - порт, по которому будет доступен веб-интерфейс сервиса

## обработчики событий

### GET /get-order

Сервис имеет только один эндпоинт по которому можно получить заказ по его uid.
Пример запроса:

```
http://localhost:8008/get-order?id=xe6ie1dxwtt4zhlmab6
```

Пример структуры ожидаемого ответа:

```json
{
  "order_uid": "xe6ie1dxwtt4zhlmab6",
  "track_number": "KYKDXSLDCVIJCV",
  "entry": "CYAC",
  "delivery": {
    "name": "Lea Kovacek Lakin",
    "phone": "8420532882",
    "zip": "61387",
    "city": "Fort Worth",
    "address": "1125 South Springs town",
    "region": "Baltimore",
    "email": "cathrinemckenzie@kertzmann.name"
  },
  "payment": {
    "transaction": "13n9p7di7954ifc2uw5",
    "request_id": "5112007037",
    "currency": "rub",
    "provider": "provider",
    "amount": 39962,
    "payment_dt": 128389,
    "bank": "Alfonzo Bosco",
    "delivery_cost": 721665,
    "goods_total": 160246,
    "custom_fee": 23771
  },
  "items": [
    {
      "chrt_id": 306430,
      "track_number": "HTDSSRGQJYXANL",
      "price": 778324,
      "rid": "nsb0o0bhm5zlnylf488m5",
      "name": "Abel Upton",
      "sale": 41,
      "size": "52",
      "total_price": 207298,
      "nm_id": 1013636,
      "brand": "Ethel Jakubowski",
      "status": 0
    },
    {
      "chrt_id": 342787,
      "track_number": "HZAVAFMYWOUPFW",
      "price": 356309,
      "rid": "th2f2uprzl8asccz2z9ew",
      "name": "Allen Schamberger",
      "sale": 40,
      "size": "94",
      "total_price": 199379,
      "nm_id": 855203,
      "brand": "Abbigail Hirthe",
      "status": 111
    }
  ],
  "locale": "en",
  "internal_signature": "test",
  "customer_id": "text",
  "delivery_service": "Winston Blick",
  "shardkey": "7",
  "sm_id": 23,
  "date_created": "2023.01.02 22:15:50",
  "oof_shard": "8"
}
```

В веб-интерфейсе заказ выглядит следующим образом: 
![img.png](img.png)

## генерация заказов

Для теста сервиса создан микросервис, который генерирует заказы автоматически и отсылает их в канал nats-streaming. 
Сервис доступен по ссылке [neutral-jungle-creep/nats-publisher](https://github.com/neutral-jungle-creep/nats-publisher)

# Приложение для определения погоды
## Базовая настройка
```shell
docker compose -f provisioning/docker-compose.local.yaml
```
```shell
./cli-wether-app migrations
```

## Запуск
```shell
./cli-wether-app getToday -c Bishkek
```

## Помощь
```shell
./cli-wether-app getToday -h
```
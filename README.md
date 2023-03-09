# Приложение для определения погоды
## Установка зависимостей
```shell
docker compose -f provisioning/docker-compose.local.yaml
```

## Запуск
```shell
./cli-wether-app getToday -c <CitiName> // Например Bishkek
```

## Помощь
```shell
./cli-wether-app getToday -h
```
# Seven Seas

[![golangci-lint](https://github.com/efumagal/sevenseas/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/efumagal/sevenseas/actions/workflows/golangci-lint.yml)
[![test](https://github.com/efumagal/sevenseas/actions/workflows/test.yml/badge.svg)](https://github.com/efumagal/sevenseas/actions/workflows/test.yml)

## Run

```shell
docker compose up --build -d
```

```shell
docker compose down
```

When 

[Redis Web UI](http://localhost:8001/redis-stack/browser)

## Develop

For development it is possible to run the docker compose that is just starting Redis

```shell
docker compose -f docker-compose-development.yml up --build -d
```

To run:

```shell
go run main.go
```

```shell
docker compose down
```

## TO DOs

- Add constructor with validation for `Model`
- `Model.Coordinates` could be a specific struct with lat, lon and validation

## Links

- GORM Guides [https://gorm.io](https://gorm.io)
- Gen Guides [https://gorm.io/gen/index.html](https://gorm.io/gen/index.html)

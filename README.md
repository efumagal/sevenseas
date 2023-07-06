# Seven Seas

[![golangci-lint](https://github.com/efumagal/sevenseas/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/efumagal/sevenseas/actions/workflows/golangci-lint.yml)
[![test](https://github.com/efumagal/sevenseas/actions/workflows/test.yml/badge.svg)](https://github.com/efumagal/sevenseas/actions/workflows/test.yml)

## PlantUML
![uml](https://github.com/efumagal/sevenseas/assets/77152760/30c67497-f312-4b30-8631-f062fa4469c6)

## Run

```shell
docker compose up --build -d
```

```shell
docker compose down
```

When completed access the Redis UI:

[Redis Web UI](http://localhost:8001/redis-stack/browser)  

<img width="618" alt="Screenshot 2023-06-29 at 22 57 08" src="https://github.com/efumagal/sevenseas/assets/77152760/54fb234c-f5d1-48b4-8e91-f396b421a0aa">

(Accept Terms the first time, could probably be avoided with some config passed to the Docker container)

<img width="800" alt="Screenshot 2023-06-29 at 20 17 41" src="https://github.com/efumagal/sevenseas/assets/77152760/545c01dd-4c9b-4c0c-bd6e-15f39a4c108c">


## Develop

For development it is possible to run the docker compose that is just starting Redis

```shell
docker compose -f docker-compose-development.yml up --build -d
```

To run:

```shell
go run main.go

2023/06/30 09:56:27 Starting
2023/06/30 09:56:27 Redis host localhost:6379
2023/06/30 09:56:27 Port file ../data/ports.json
2023/06/30 09:56:27 Added 1632 Took 799.934125ms
```

```shell
docker compose down
```

## Tests

```shell
go test ./...
```

At the moment there are only few tests for demonstration.  
Coverage can be checked with:  

```shell
go test ./... -cover
```

In a real repo there may be a check on minimum coverage.

Unit tests:  
- [Domain](internal/core/domain/model_test.go) 
- [Utils](utils/utils_test.go)
- [Memory Repo](internal/adapters/repository/memory_test.go)  
 
Functional test (sort of):  
- [Injector](internal/adapters/injector/memory_test.go)  

Apart from the tests in the repo there will be end to end tests involving real external resources (DBs, third party services...).  

## TO DOs

- Add constructor with validation for `PortData`
- `PortData.Coordinates` could be a specific struct with lat, lon and validation
- Mock Redis and Postgres to be able to create tests
- Add more tests and maybe a check on the coverage
- Now the service that is injesting to the DB exits immediately, with more time an API (gRPC, REST) could have been added
- Better naming and spend more time on the code structure to fit with Hexagonal patterns

## Notes  
* Generated a file with 100k random ports, that took ~30s to add to the DB.  
In case of large files (although I think the data would probably be on another DB) 
it might be worth investigating if it is possible to split the file and upload to the DB in parallel.  
* Redis was chosen for simplicity but depending on the access patterns (maybe a query by Country) another DB might be more appropriate.
* When tested with Redis the `SavePort` function was also updating if item already present
## Links

- Redis mock [https://github.com/elliotchance/redismock](https://github.com/elliotchance/redismock)

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

When completed access the Redis UI:

[Redis Web UI](http://localhost:8001/redis-stack/browser)  

<img width="418" alt="Screenshot 2023-06-29 at 22 57 08" src="https://github.com/efumagal/sevenseas/assets/77152760/ab58a60d-5940-4936-b5cc-aac844e7439a">

(Accept Terms the first time, could probably be avoided with some config passed to the Docker container)

<img width="800" alt="Redis Web UI" src="https://github.com/efumagal/sevenseas/assets/77152760/962bcbb0-2f46-4017-ada8-c8f6807baf54">

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

## Tests

```shell
go test ./...
```

## TO DOs

- Add constructor with validation for `PortData`
- `PortData.Coordinates` could be a specific struct with lat, lon and validation
- Mock Redis and Postgres to be able to create tests
- Add more tests and maybe a check on the coverage
- Now the app that is injesting to the DB just start with the container, with more time an API (gRPC, REST) could have been added
- Better naming and spend more time on the code structure to fit Hexagonal patterns

## Notes  
* Generated a file with 100k random ports, that took ~30s to add to the DB.  
In case of large files (although I think the data would probably be on another DB) 
it might be worth investigating if it is possible to split the file and upload to the DB in parallel.  
* Redis was chosen for simplicity but depending on the access patterns (maybe a query by Country) another DB might be more appropriate.
* When tested with Redis the `SavePort` function was also updating if item already present
## Links

- Redis mock [https://github.com/elliotchance/redismock](https://github.com/elliotchance/redismock)
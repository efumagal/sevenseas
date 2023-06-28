docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest

## Build Docker image locally

`docker build . -t ports-management:0.0.1`  
`docker run ports-management:0.0.1`

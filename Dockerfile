FROM golang:1.20.5-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ports-management ./cmd/main.go

FROM alpine:latest AS runner
WORKDIR /app
COPY data ./data
COPY --from=builder /app/ports-management .
EXPOSE 8080
ENTRYPOINT ["./ports-management"]
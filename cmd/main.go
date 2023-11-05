package main

import (
	"log"
	"os"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/efumagal/sevenseas/internal/adapters/injector"
	"github.com/efumagal/sevenseas/internal/adapters/repository"
	"github.com/efumagal/sevenseas/internal/core/services"
	"github.com/efumagal/sevenseas/utils"
)

var (
	redisHost = utils.GetEnv("REDIS_ENDPOINT", "localhost:6379")
	portFile  = utils.GetEnv("PORT_FILE", "../data/ports.json")
	svc       *services.PortService
)

func handler(signal os.Signal) {
	if signal == syscall.SIGTERM {
		log.Println("Got kill signal. ")
		log.Println("Program will terminate now.")
		os.Exit(0)
	} else if signal == syscall.SIGINT {
		log.Println("Got CTRL+C signal")
		log.Println("Closing.")
		os.Exit(0)
	} else {
		log.Println("Ignoring signal: ", signal)
	}
}

func main() {
	log.Println("Starting")
	log.Println("Redis host", redisHost)
	log.Println("Port file", portFile)
	start := time.Now()

	repo := "redis"

	switch repo {
	case "postgres":
		store := repository.NewPortPostgresRepository()
		svc = services.NewPortService(store)
	default:
		store := repository.NewPortRedisRepository(redisHost)
		svc = services.NewPortService(store)
	}

	pfs := injector.NewPortStreamService(svc)

	f, err := os.Open(portFile)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", portFile, err.Error())
	}
	defer f.Close()

	inserted, err := pfs.InjectStream(f)
	if err != nil {
		log.Printf("Error decoding stream %v", err.Error())
	}

	timeElapsed := time.Since(start)
	log.Printf("Added %d Took %s", inserted, timeElapsed)

	id := "USUXZ"
	retrievedPort, err := svc.GetPort(id)

	if err != nil {
		log.Printf("ID:%s City:%s", id, retrievedPort.City)
	} else {
		log.Printf("ID:%s not found", id)
	}
}

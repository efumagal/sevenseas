package main

import (
	"log"
	"os"
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

	pfs := injector.NewPortFileService(svc)

	f, err := os.Open(portFile)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", portFile, err.Error())
	}
	defer f.Close()

	err = pfs.InjectStream(f)

	if err != nil {
		log.Printf("Error decoding stream %v", err.Error())
	}

	log.Println("Finished")
	timeElapsed := time.Since(start)
	log.Printf("Took %s", timeElapsed)
}

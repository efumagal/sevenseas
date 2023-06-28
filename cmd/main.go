package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/efumagal/sevenseas/internal/adapters/repository"
	"github.com/efumagal/sevenseas/internal/core/domain"
	"github.com/efumagal/sevenseas/internal/core/services"
)

var (
	redisHost = "localhost:6379"
	svc       *services.PortService
)

func decodeStream(r io.Reader, svc *services.PortService) error {
	dec := json.NewDecoder(r)
	t, err := dec.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("expected {, got %v", t)
	}
	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return err
		}
		key := t.(string)

		var value domain.Model
		if err := dec.Decode(&value); err != nil {
			return err
		}
		// fmt.Printf("key %q, value %#v\n", key, value)

		port := domain.Port{Model: value, ID: key}

		err = svc.SavePort(port)
		if err != nil {
			log.Println(err)
		}

	}
	return nil
}

func main() {
	log.Println("Starting")
	start := time.Now()

	repo := "redis"

	switch repo {
	case "redis":
		store := repository.NewPortRedisRepository(redisHost)
		svc = services.NewPortService(store)
	default:
		store := repository.NewPortPostgresRepository()
		svc = services.NewPortService(store)
	}

	port := domain.Port{}
	port.ID = "pippo"

	err := svc.SavePort(port)
	if err != nil {
		log.Println(err)
	}

	fileName := "../ports.json"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}
	defer f.Close()

	err = decodeStream(f, svc)

	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	log.Println("Finished")
	timeElapsed := time.Since(start)
	log.Printf("Took %s", timeElapsed)
}

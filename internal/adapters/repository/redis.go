package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/efumagal/sevenseas/internal/core/domain"

	"github.com/redis/go-redis/v9"
)

type PortRedisRepository struct {
	client *redis.Client
}

func NewPortRedisRepository(host string) *PortRedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
	return &PortRedisRepository{
		client: client,
	}
}

func (r *PortRedisRepository) SavePort(port domain.Port) error {
	ctx := context.Background()
	json, err := json.Marshal(port.PortData)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, port.ID, json, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *PortRedisRepository) GetPort(id string) (domain.PortData, error) {
	ctx := context.Background()

	objStr, err := r.client.Get(ctx, id).Result()
	if err != nil {
		return domain.PortData{}, err
	}

	b := []byte(objStr)

	portData := &domain.PortData{}
	err = json.Unmarshal(b, portData)
	if err != nil {
		return domain.PortData{}, err
	}

	return *portData, errors.New("port not found")
}

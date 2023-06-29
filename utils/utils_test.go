package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvDefault(t *testing.T) {
	result := GetEnv("REDIS_ENDPOINT", "localhost:6379")
	assert.Equal(t, "localhost:6379", result)
}

func TestGetEnv(t *testing.T) {
	t.Setenv("REDIS_ENDPOINT", "172.10.12.11:6379")
	result := GetEnv("REDIS_ENDPOINT", "localhost:6379")
	assert.Equal(t, "172.10.12.11:6379", result)
}

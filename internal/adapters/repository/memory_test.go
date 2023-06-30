package repository

import (
	"testing"

	"github.com/efumagal/sevenseas/internal/core/domain"
	"github.com/efumagal/sevenseas/internal/core/services"
	"github.com/stretchr/testify/assert"
)

func TestSingleSave(t *testing.T) {
	store := NewPortMemoryRepository()
	svc := services.NewPortService(store)

	id := "ABC"
	city := "Brunswick"

	p := domain.Port{}
	p.ID = id
	p.City = city

	err := svc.SavePort(p)
	assert.Nil(t, err)

	retrievedPort, err := svc.GetPort(id)
	assert.Nil(t, err)
	assert.Equal(t, city, retrievedPort.City)
}

package repository

import (
	"errors"

	"github.com/efumagal/sevenseas/internal/core/domain"
)

type PortMemoryRepository struct {
	allPorts map[string]domain.PortData
}

func NewPortMemoryRepository() *PortMemoryRepository {
	allPorts := make(map[string]domain.PortData)
	return &PortMemoryRepository{
		allPorts: allPorts,
	}
}

func (r *PortMemoryRepository) SavePort(port domain.Port) error {
	r.allPorts[port.ID] = port.PortData
	return nil
}

func (r *PortMemoryRepository) GetPort(id string) (domain.PortData, error) {
	portData, ok := r.allPorts[id]
	if ok {
		return portData, nil
	}
	// could be a sentinel error common for all repositories
	return domain.PortData{}, errors.New("port not found")
}

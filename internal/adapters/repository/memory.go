package repository

import (
	"github.com/efumagal/sevenseas/internal/core/domain"
)

type PortMemoryRepository struct {
	allPorts map[string]domain.Model
}

func NewPortMemoryRepository() *PortMemoryRepository {
	allPorts := make(map[string]domain.Model)
	return &PortMemoryRepository{
		allPorts: allPorts,
	}
}

func (r *PortMemoryRepository) SavePort(port domain.Port) error {
	r.allPorts[port.ID] = port.Model
	return nil
}

package services

import (
	"github.com/efumagal/sevenseas/internal/core/domain"
	"github.com/efumagal/sevenseas/internal/core/ports"
)

type PortService struct {
	repo ports.PortRepository
}

func NewPortService(repo ports.PortRepository) *PortService {
	return &PortService{
		repo: repo,
	}
}

func (m *PortService) SavePort(port domain.Port) error {
	return m.repo.SavePort(port)
}

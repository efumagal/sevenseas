package ports

import "github.com/efumagal/sevenseas/internal/core/domain"

type PortRepository interface {
	SavePort(message domain.Port) error
	GetPort(id string) (domain.PortData, error)
}

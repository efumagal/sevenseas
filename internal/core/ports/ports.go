package ports

import "github.com/efumagal/sevenseas/internal/core/domain"

type PortService interface {
	SavePort(message domain.Port) error
}

type PortRepository interface {
	SavePort(message domain.Port) error
}

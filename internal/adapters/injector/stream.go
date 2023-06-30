package injector

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/efumagal/sevenseas/internal/core/domain"
	"github.com/efumagal/sevenseas/internal/core/services"
)

type PortWriter interface {
	InjectStream(r io.Reader) error
}

type PortStreamService struct {
	svc *services.PortService
}

func NewPortStreamService(svc *services.PortService) *PortStreamService {
	return &PortStreamService{
		svc: svc,
	}
}

func (pfs *PortStreamService) InjectStream(r io.Reader) (int, error) {
	inserted := 0
	dec := json.NewDecoder(r)
	t, err := dec.Token()
	if err != nil {
		return inserted, err
	}
	if t != json.Delim('{') {
		return inserted, fmt.Errorf("expected {, got %v", t)
	}
	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return inserted, err
		}
		key := t.(string)

		var value domain.PortData
		if err := dec.Decode(&value); err != nil {
			return inserted, err
		}

		port := domain.Port{PortData: value, ID: key}

		err = pfs.svc.SavePort(port)
		// Could use sentinel errors or wrap error to return specific reason
		// for failure and handle appropriately
		if err != nil {
			log.Println(err)
		} else {
			inserted++
		}

	}
	return inserted, nil
}

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

type PortFileService struct {
	svc *services.PortService
}

func NewPortFileService(svc *services.PortService) *PortFileService {
	return &PortFileService{
		svc: svc,
	}
}

func (pfs *PortFileService) InjectStream(r io.Reader) error {
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

		port := domain.Port{Model: value, ID: key}

		err = pfs.svc.SavePort(port)
		if err != nil {
			log.Println(err)
		}

	}
	return nil
}

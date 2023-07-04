package domain

import (
	"errors"
	"strings"
)

const IDLength = 5

var (
	// ErrInvalidPortID is returned when the port id is not valid
	ErrInvalidPortID      = errors.New("port Id is not valid")
	ErrEmptyCity          = errors.New("city should not be empty")
	ErrInvalidCoordinates = errors.New("coordinates should be provide as lon lat floats")
)

// This could probably be a Value Object
type PortData struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type Port struct {
	ID string
	PortData
}

func isValidPortID(id string) bool {
	return strings.ToUpper(id) == id && len(id) == IDLength
}

// depending on the usage pattern it could return a pointer to Port
func NewPort(id string, portData PortData) (Port, error) {
	// just validating few fields for simplicity
	if !isValidPortID(id) {
		return Port{}, ErrInvalidPortID
	}

	if len(portData.City) == 0 {
		return Port{}, ErrEmptyCity
	}

	// simplified check, could be more specific
	// coordinates could also be a separate struct
	if len(portData.Coordinates) != 2 {
		return Port{}, ErrInvalidCoordinates
	}

	return Port{ID: id, PortData: portData}, nil
}

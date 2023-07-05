package domain

import (
	"testing"

	"github.com/stretchr/testify/assert" // not ideal to have exernal dep here
)

func TestNewPort(t *testing.T) {
	type checks func(t *testing.T, port Port)
	// Build our needed testcase data struct
	type testCase struct {
		test             string
		id               string
		portData         PortData
		expectedError    error
		additionalChecks checks
	}
	// Create new test cases
	testCases := []testCase{
		{
			test:          "Empty Id validation",
			id:            "",
			portData:      PortData{},
			expectedError: ErrInvalidPortID,
		},
		{
			test:          "Invalid Id lower case",
			id:            "alksa",
			portData:      PortData{},
			expectedError: ErrInvalidPortID,
		},
		{
			test:          "Empty City",
			id:            "AEAJM",
			portData:      PortData{City: ""},
			expectedError: ErrEmptyCity,
		},
		{
			test:          "Invalid coordinates",
			id:            "AEAJM",
			portData:      PortData{City: "Ajman", Coordinates: []float64{1.1}},
			expectedError: ErrInvalidCoordinates,
		},
		{
			test: "Valid",
			id:   "AEAJM",
			portData: PortData{City: "Ajman", Coordinates: []float64{55.5136433,
				25.4052165}},
			expectedError: nil,
			additionalChecks: func(t *testing.T, port Port) {
				assert.Equal(t, "AEAJM", port.ID)
				assert.Equal(t, "Ajman", port.City)
				assert.Len(t, port.Coordinates, 2)
				assert.InDelta(t, 55.513643, port.Coordinates[0], 0.00001)
			},
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new port
			port, err := NewPort(tc.id, tc.portData)
			// Check if the error matches the expected error
			assert.Equal(t, tc.expectedError, err)
			// Run additional checks on created port
			if tc.additionalChecks != nil {
				tc.additionalChecks(t, port)
			}
		})
	}
}

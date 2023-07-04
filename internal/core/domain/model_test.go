package domain

import (
	"testing"

	"github.com/stretchr/testify/assert" // not ideal to have exernal dep here
)

func TestNewPort(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		test          string
		id            string
		portData      PortData
		expectedError error
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
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new port
			_, err := NewPort(tc.id, tc.portData)
			// Check if the error matches the expected error
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

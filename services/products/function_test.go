package products

import (
	"testing"
)

func TestIDValidation(t *testing.T) {
	testCases := []struct {
		desc        string
		input       int
		expectedOut bool
	}{
		{
			desc:        "Success",
			input:       1,
			expectedOut: true,
		},
		{
			desc:        "Failure",
			input:       0,
			expectedOut: false,
		},
	}
	for _, test := range testCases {
		ts := test
		t.Run(ts.desc, func(t *testing.T) {
			result := checkID(ts.input)
			if result != ts.expectedOut {
				t.Errorf("Expected %v obtained %v", ts.expectedOut, result)
			}
		})
	}
}

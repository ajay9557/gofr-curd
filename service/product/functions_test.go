package product

import (
	"strconv"
	"testing"
)

func TestValidateId(t *testing.T) {
	testCases := []struct {
		caseID      int
		input       int
		expectedOut bool
	}{
		// Success case
		{
			caseID:      1,
			input:       1,
			expectedOut: true,
		},
		// Error cases
		{
			caseID:      2,
			input:       -1,
			expectedOut: false,
		},
	}

	for _, test := range testCases {
		tc := test
		t.Run("testing "+strconv.Itoa(tc.caseID), func(t *testing.T) {
			out := validateID(tc.input)
			if out != tc.expectedOut {
				t.Errorf("TestCase[%v] Expected: \\t%v\\nGot: \\t%v\\n\\", tc.caseID, tc.expectedOut, out)
			}
		})
	}
}

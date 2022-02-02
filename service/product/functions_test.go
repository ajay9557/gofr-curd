package product

import (
	"strconv"
	"testing"
)

func TestValidateId(t *testing.T) {
	testCases := []struct {
		caseId      int
		input       int
		expectedOut bool
	}{
		// Success case
		{
			caseId:      1,
			input:       1,
			expectedOut: true,
		},
		// Error cases
		{
			caseId:      2,
			input:       -1,
			expectedOut: false,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			out := validateId(tc.input)
			if out != tc.expectedOut {
				t.Errorf("TestCase[%v] Expected: \\t%v\\nGot: \\t%v\\n\\", tc.caseId, tc.expectedOut, out)
			}
		})
	}
}

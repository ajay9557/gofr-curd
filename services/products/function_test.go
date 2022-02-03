package products

import (
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	testIdValidation(t, app)
}

func testIdValidation(t *testing.T, app *gofr.Gofr) {
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
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			result := checkId(ts.input)
			if result != ts.expectedOut {
				t.Errorf("Expected %v obtained %v", ts.expectedOut, result)
			}
		})
	}

}

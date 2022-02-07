package product

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIdValidation(t *testing.T) {
	testCases := []struct {
		desc        string
		id          int
		expectedRes bool
	}{
		{
			desc:        "True case",
			id:          1,
			expectedRes: true,
		},
		{
			desc:        "False case",
			id:          0,
			expectedRes: false,
		},
	}

	for _, v := range testCases {
		ts := v
		t.Run(ts.desc, func(t *testing.T) {
			res := idValidation(ts.id)
			if !reflect.DeepEqual(ts.expectedRes, res) {
				fmt.Print("expected ", ts.expectedRes, "obtained", res)
			}
		})
	}
}

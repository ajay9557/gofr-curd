package product

import (
	"errors"
	"product/models"
	"product/stores"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
)

func Test_GetProductById(t *testing.T) {
	app := gofr.New()
	controller := gomock.NewController(t)
	mock := stores.NewMockStore(controller)
	service := New(mock)

	testCases := []struct {
		desc           string
		input          int
		mock           []*gomock.Call
		expectedOutput interface{}
		expectedError  error
	}{
		{
			desc:           "Test Case 1",
			input:          1,
			expectedOutput: models.Product{Id: 1, Name: "lg", Type: "machine"},
			expectedError:  nil,
			mock: []*gomock.Call{
				mock.EXPECT().GetProductById(gomock.Any(), 1).Return(models.Product{Id: 1, Name: "lg", Type: "machine"}, nil),
			},
		},
		{
			desc:           "Test Case 2",
			input:          -1,
			mock:           nil,
			expectedOutput: nil,
			expectedError:  errors.New("INVALID ID"),
		},
	}

	for _, tcs := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		result, err := service.GetProductById(ctx, tcs.input)
		if !reflect.DeepEqual(err, tcs.expectedError) {
			t.Errorf("Expected: %s, Output: %s", tcs.expectedError, err)
		}
		if tcs.expectedError == nil && !reflect.DeepEqual(result, tcs.expectedOutput) {
			t.Errorf("Expected: %v, Output: %v", tcs.expectedOutput, result)
		}
	}
}

package products

import (
	"errors"
	"reflect"
	"testing"

	"github.com/arohanzst/testapp/models"
	"github.com/arohanzst/testapp/stores"

	gerror "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
)

func Test_ReadByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := stores.NewMockProduct(ctrl)

	s := New(mock)

	testCases := []struct {
		desc     string
		input    int
		mockCall []*gomock.Call
		expOut   *models.Product
		expErr   error
	}{
		{
			desc:  "Case:1",
			input: 1,
			expOut: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().ReadByID(gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),
			},
		},
		{
			desc:   "Case:2",
			input:  -10,
			expErr: errors.New("Invalid Id"),
		},
		{
			desc:  "Case:3",
			input: 10,
			expErr: gerror.EntityNotFound{
				Entity: "Product",
				ID:     "10",
			},
			mockCall: []*gomock.Call{
				mock.EXPECT().ReadByID(gomock.Any(), 10).Return(nil,
					gerror.EntityNotFound{
						Entity: "Product",
						ID:     "10",
					}),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)

		resp, err := s.ReadByID(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, resp)
		}

	}
}

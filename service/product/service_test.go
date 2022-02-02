package product

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
	"gofrPractice/models"
	"gofrPractice/store"
	"reflect"
	"testing"
)

func TestServices_GetById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)

	testCases := []struct {
		desc     string
		input    int
		mockCall []*gomock.Call
		expOut   *models.Product
		expErr   error
	}{
		{
			desc:  "valid id",
			input: 1,
			expOut: &models.Product{
				Id:   1,
				Name: "Part1",
				Type: "hardware",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "Part1",
					Type: "hardware",
				}, nil),
			},
		},
		{
			desc:  "negative id",
			input: -1,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
		},
		{
			desc:  "id not in database",
			input: 1002,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1002",
			},
			mockCall: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 1002).Return(nil,
					errors.EntityNotFound{
						Entity: "product",
						ID:     "1002",
					}),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)

		resp, err := s.GetById(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, resp)
		}

	}
}

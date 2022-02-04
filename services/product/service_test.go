package product

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"github.com/himanshu-kumar-zs/gofr-curd/store"
	"reflect"
	"testing"
)

func TestProductService_GetByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	service := New(mockStore)

	testcases := []struct {
		desc     string
		inp      int
		exp      *models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc: "id exists",
			inp:  1,
			exp: &models.Product{
				Id:   1,
				Name: "legion",
				Type: "Laptop",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).
					Return(&models.Product{
						1,
						"legion",
						"Laptop",
					}, nil),
			},
		},
		{
			desc: "id does not exists",
			inp:  1,
			exp:  nil,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1",
			},
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).
					Return(nil, errors.EntityNotFound{Entity: "product", ID: "1"}),
			},
		},
		{
			desc: "invalid id",
			inp:  -1,
			exp:  nil,
			expErr: errors.InvalidParam{
				[]string{"id"},
			},
		},
	}
	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)

		out, err := service.GetByID(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.exp, out)
		}

	}
}

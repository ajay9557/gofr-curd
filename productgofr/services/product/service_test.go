package product

import (
	"testing"
	models "zopsmart/productgofr/models"
	store "zopsmart/productgofr/stores"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
//	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
//	"golang.org/x/tools/go/expect"
	"fmt"
	"context"
	"reflect"
)

func TestGetProdById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s:= New(mock)
	tests := []struct {
		desc string
		input int
		mockcall []*gomock.Call
		expectedRes *models.Product
		expectedErr error
	}{
		{
			desc: "Sucess Case",
			input:1,
			expectedRes: &models.Product{
				Id : 1,
				Name : "shirt",
				Type: "fashion",

			},
			mockcall: []*gomock.Call{
				mock.EXPECT().GetProdByID(gomock.Any(),gomock.Any()).Return(&models.Product{
				Id : 1,
				Name : "shirt",
				Type: "fashion",
				},nil),
			},
			expectedErr:  nil,
		},
		{
			desc: "Failure case",
			input:1,
			expectedRes: &models.Product{
				Id : 352,
				Name : "laptop",
				Type: "electronics",

			},
			mockcall: []*gomock.Call{
				mock.EXPECT().GetProdByID(gomock.Any(),gomock.Any()).Return(&models.Product{
					Id : 352,
					Name : "laptop",
					Type: "electronics",
				},nil),
			},
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "352"},

		},

	}

	for _,tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			res, err := s.GetProdByID(ctx,tc.input)
			if err != nil && tc.expectedErr != err {
				fmt.Print("expected ", tc.expectedErr, "obtained", err)
			}
			if !reflect.DeepEqual(tc.expectedRes, res) {
				fmt.Print("expected ", tc.expectedRes, "obtained", res)
			}
	}


	
}
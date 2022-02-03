package product

import (
	"context"
	"errors"
	"fmt"
	"gofr-curd/models"
	"gofr-curd/store"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
)

func TestGetProductById(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mockHandler := New(mockStore)

	testCases := []struct {
		desc        string
		id          int
		mock        []*gomock.Call
		expectedRes models.Product
		expectedErr error
	}{
		{
			desc: "Success case",
			id:   1,
			mock: []*gomock.Call{
				mockStore.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(models.Product{
					Id:   1,
					Name: "jeans",
					Type: "clothes",
				}, nil),
			},
			expectedRes: models.Product{
				Id:   1,
				Name: "jeans",
				Type: "clothes",
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case - 1",
			id:   2,
			mock: []*gomock.Call{
				mockStore.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(models.Product{},
					errors.New("product not found")),
			},
			expectedRes: models.Product{},
			expectedErr: errors.New("product not found"),
		},
		{
			desc:        "Failure case - 2",
			id:          0,
			expectedRes: models.Product{},
			expectedErr: errors.New("invalid id"),
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {

			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()
			res, err := mockHandler.GetByProductId(ts.id, ctx)

			if !reflect.DeepEqual(ts.expectedRes, res) {
				fmt.Print("expected ", ts.expectedRes, "obtained", res)
			}
			if !reflect.DeepEqual(ts.expectedErr, err) {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

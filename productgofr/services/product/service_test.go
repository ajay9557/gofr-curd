package product

import (
	"reflect"
	"testing"
	models "zopsmart/productgofr/models"
	store "zopsmart/productgofr/stores"

	"context"
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
)

func TestGetProdById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	
	tests := []struct {
		desc        string
		Id          int
		mockcall    *gomock.Call
		expectedRes models.Product
		expectedErr error
	}{
		{
			desc: "Sucess Case",
			Id:   1,
			expectedRes: models.Product{
				Id:   1,
				Name: "shirt",
				Type: "fashion",
			},
			mockcall: 
				mock.EXPECT().GetProdByID(ctx, 1).Return(models.Product{
					Id:   1,
					Name: "shirt",
					Type: "fashion",
				}, nil),
			expectedErr: nil,
		},

		{
			desc: "Failure case",
			Id:   352,
			expectedRes: models.Product{},
			mockcall: mock.EXPECT().GetProdByID(ctx, 352).Return(models.Product{}, errors.EntityNotFound{Entity: "product", ID: "352"}),
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "352"},
		},
		{
			desc: "Failure case",
			Id:   -1,
			expectedRes: models.Product{},
			mockcall: nil,
			expectedErr: errors.InvalidParam{Param: []string{"id"}},
	    },

	}

	for _, tc := range tests {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		res, err := s.GetProdByID(ctx, tc.Id)
		if err != nil && reflect.DeepEqual(tc.expectedErr, err) {
			fmt.Print("expected ", tc.expectedErr, "Got", err)
		}
		if !reflect.DeepEqual(tc.expectedRes, res) {
			fmt.Print("expected ", tc.expectedRes, "Got", res)
		}
	}
}

func TestDeleteProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	
	testCases := []struct {
		desc     string
		id       int
		expErr   error
		mockCall *gomock.Call
	}{
		{
			desc:   "success",
			id:     2,
			expErr: nil,
			mockCall:
				mock.EXPECT().DeleteProduct(ctx, 2).Return(nil),
		},
		{
			desc:   "entity not found",
			id:     5,
			expErr: errors.EntityNotFound{Entity: "product", ID: "5"},
			mockCall: mock.EXPECT().DeleteProduct(ctx, 5).Return(errors.EntityNotFound{Entity: "product", ID: "5"}),
		},
		{
			desc:     "invalid param id",
			id:       -1,
			expErr:   errors.InvalidParam{Param: []string{"id"}},
			mockCall: nil,
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		err := s.DeleteProduct(ctx, tc.id)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}

}

func TestUpdateProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	testCases := []struct {
		desc        string
		Input       models.Product
		expectedErr error
		mock        *gomock.Call
	}{
		{
			desc: "Success Case",
			Input: models.Product{
				Id:   1,
				Name: "Laptop",
				Type: "electronics",
			},
			expectedErr: nil,
			mock: 
				mock.EXPECT().UpdateProduct(ctx, models.Product{
					Id:   1,
					Name: "Laptop",
					Type: "electronics",
				}).Return(nil),
		},

		{
			desc:        "Invalid id",
			Input:       models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"id"}},
			mock:        nil,
		},

		{
			desc: "Invalid id",
			Input: models.Product{
				Id: -1,
			},
			expectedErr: errors.InvalidParam{Param: []string{"id"}},
			mock:        nil,
		},
		{
			desc: "Invalid id",
			Input: models.Product{
				Id: 100,
				Name: "Laptop",
				Type: "electronics",
			},
			expectedErr: errors.EntityNotFound{Entity: "products",ID: "id"},
			mock:        nil,
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		 err := s.UpdateProduct(ctx, tc.Input)

		if !reflect.DeepEqual(tc.expectedErr, err) {
			fmt.Print("expected ", tc.expectedErr, "obtained", err)
		}

	}

}

func TestCreateProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	testCases := []struct {
		desc     string
		input    models.Product
		expErr   error
		mockCall *gomock.Call
	}{
		{
			desc: "success",
			input: models.Product{
				Id:   4,
				Name: "mouse",
				Type: "electronics",
			},
			expErr: nil,
			mockCall:
				mock.EXPECT().CreateProduct(gomock.Any(), models.Product{
					Id:   4,
					Name: "mouse",
					Type: "electronics",
				}).Return(nil),
		},

		{
			desc: "invalid param id",
			input: models.Product{
				Id: -1,
			},
			expErr:   errors.InvalidParam{Param: []string{"Id"}},
			mockCall: nil,
		},

		{
			desc: "error entity already exists",
			input: models.Product{
				Id:   1,
				Name: "shirt",
				Type: "fashion",
			},
			expErr: errors.EntityAlreadyExists{},
			mockCall: mock.EXPECT().CreateProduct(gomock.Any(), models.Product{
							Id:   1,
							Name: "shirt",
							Type: "fashion",
						}).Return(errors.EntityAlreadyExists{}),
		},
}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		err := s.CreateProduct(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

	}

}

func TestGetAllProduct(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := store.NewMockStore(ctrl)
	s := New(mock)
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	testCases := []struct {
		desc        string
		expectedRes []models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc: "Fetching all products",
			expectedRes: []models.Product{
				{
					Id:   1,
					Name: "shirt",
					Type: "fashion",
				},
				{
					Id:   2,
					Name: "mobile",
					Type: "electronics",
				},
			},
			expectedErr: nil,
			mockCall:
				mock.EXPECT().GetAllProduct(ctx).Return([]models.Product{
					{
						Id:   1,
						Name: "shirt",
						Type: "fashion",
					},
					{
						Id:   2,
						Name: "mobile",
						Type: "electronics",
					},
				}, nil)},
		{
			desc:        "Error while fetching products",
			expectedRes: []models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products"},
			mockCall: mock.EXPECT().GetAllProduct(ctx).Return([]models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			res, err := s.store.GetAllProduct(ctx)

			if !reflect.DeepEqual(tc.expectedErr,err) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedErr, err)
			}

			if !reflect.DeepEqual(res, tc.expectedRes) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedRes, res)
			}
		})
	}
}
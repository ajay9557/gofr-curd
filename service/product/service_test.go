package product

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	err1 "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/store"
)

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockProductStore(ctrl)
	mockHandler := New(mockStore)

	testCases := []struct {
		desc     string
		input    int
		mockCall []*gomock.Call
		expOut   models.Product
		expErr   error
	}{
		{
			desc:  "Case 1: Success Case",
			input: 1,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetProductByID(gomock.Any(), 1).Return(models.Product{
					ID:   1,
					Name: "name-1",
					Type: "type-1",
				}, nil),
			},
			expOut: models.Product{
				ID:   1,
				Name: "name-1",
				Type: "type-1",
			},
			expErr: nil,
		},
		{
			desc:  "Case 2: Failure Case1",
			input: 22,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(models.Product{},
					errors.New("product not found for the given ID")),
			},
			expOut: models.Product{},
			expErr: errors.New("product not found for the given ID"),
		},
		{
			desc:   "Case 3: Failure Case2",
			input:  0,
			expOut: models.Product{},
			expErr: errors.New("invalID ID"),
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()
			res, err := mockHandler.GetProductByID(ctx, ts.input)

			if !reflect.DeepEqual(ts.expOut, res) {
				fmt.Printf("expected : %v, Got : %v", ts.expOut, res)
			}

			if !reflect.DeepEqual(ts.expErr, err) {
				fmt.Printf("expected : %v, Got : %v", ts.expErr, err)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProductStore(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct {
		desc     string
		mockCall []*gomock.Call
		expOut   []models.Product
		expErr   error
	}{
		{
			desc: "Case 1: Success Case",
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetAllProducts(gomock.Any()).Return([]models.Product{
					{
						ID:   1,
						Name: "name-1",
						Type: "type-1",
					},
				}, nil)},
			expOut: []models.Product{
				{
					ID:   1,
					Name: "name-1",
					Type: "type-1",
				},
			},
			expErr: nil,
		},
		{
			desc: "Case 2: Failure Case",
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetAllProducts(gomock.Any()).Return(nil, errors.New("error while fetching all products details"))},
			expOut: nil,
			expErr: err1.EntityNotFound{Entity: "product", ID: "2"},
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()

			resp, err := mockHandler.GetAllProducts(ctx)

			if !reflect.DeepEqual(ts.expOut, resp) {
				fmt.Printf("expected : %v, Got : %v", ts.expOut, resp)
			}

			if !reflect.DeepEqual(ts.expErr, err) {
				fmt.Printf("expected : %v, Got : %v", ts.expErr, err)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductStore(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct {
		desc     string
		product  models.Product
		mockCall []*gomock.Call
		ExpErr   error
	}{
		{
			desc: "Case 1: Success case",
			product: models.Product{
				ID:   1,
				Name: "name-1",
				Type: "type-1",
			},
			mockCall: []*gomock.Call{mockStore.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(models.Product{
				ID:   1,
				Name: "name-1",
				Type: "type-1",
			}, nil)},
			ExpErr: nil,
		},
		{
			desc: "Case 2: Failure case2",
			product: models.Product{
				ID:   0,
				Name: "name-1",
				Type: "type-1",
			},
			ExpErr: errors.New("invalid id"),
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()

			_, err := mockHandler.UpdateProduct(ctx, ts.product)

			if !reflect.DeepEqual(err, ts.ExpErr) {
				t.Error("Expected: ", ts.ExpErr, "Got: ", err)
			}
		})
	}
}

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductStore(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct {
		desc    string
		product models.Product
		mock    []*gomock.Call
		ExpErr  error
	}{
		{
			desc: "Case 1: Success Case",
			product: models.Product{
				ID:   1,
				Name: "name-1",
				Type: "type-1",
			},
			mock: []*gomock.Call{mockStore.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(models.Product{
				ID:   1,
				Name: "name-1",
				Type: "type-1",
			}, nil)},
			ExpErr: nil,
		},
		{
			desc: "Case 1: Failure Case1",
			product: models.Product{
				ID:   2,
				Name: "name-1",
				Type: "type-1",
			},
			mock: []*gomock.Call{mockStore.EXPECT().
				CreateProduct(gomock.Any(), gomock.Any()).
				Return(models.Product{}, err1.Error("error while inserting new product"))},
			ExpErr: errors.New("cannot create the product"),
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()

			_, err := mockHandler.CreateProduct(ctx, ts.product)

			if !reflect.DeepEqual(err, ts.ExpErr) {
				t.Error("Expected: ", ts.ExpErr, "Got: ", err)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductStore(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct {
		desc   string
		ID     int
		mock   []*gomock.Call
		ExpErr error
	}{
		{
			desc:   "Case 1: Success case",
			ID:     1,
			mock:   []*gomock.Call{mockStore.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(nil)},
			ExpErr: nil,
		},
		{
			desc:   "Case 2: Failure case1",
			ID:     2,
			mock:   []*gomock.Call{mockStore.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(err1.Error("error while deleting product"))},
			ExpErr: err1.Error("error while deleting product"),
		},
		{
			desc:   "Failure case - 2",
			ID:     0,
			ExpErr: err1.Error("error while deleting product"),
		},
	}

	for _, test := range testCases {
		ts := test
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()

			err := mockHandler.DeleteProduct(ctx, ts.ID)

			if !reflect.DeepEqual(err, ts.ExpErr) {
				t.Error("Expected: ", ts.ExpErr, "Got: ", err)
			}
		})
	}
}

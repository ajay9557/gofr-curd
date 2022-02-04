package product

import (
	"context"
	"fmt"
	"gofr-curd/models"
	"gofr-curd/store"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"

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
					errors.EntityNotFound{Entity: "product", ID: "2"}),
			},
			expectedRes: models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "2"},
		},
		{
			desc:        "Failure case - 2",
			id:          0,
			expectedRes: models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"id"}},
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

func TestGetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mockHandler := New(mockStore)

	testCases := []struct {
		desc        string
		mock        []*gomock.Call
		expectedRes []models.Product
		expectedErr error
	}{
		{
			desc: "Success case",
			mock: []*gomock.Call{
				mockStore.EXPECT().GetAllProducts(gomock.Any()).Return([]models.Product{
					{Id: 1,
						Name: "jeans",
						Type: "clothes",
					},
				}, nil),
			},
			expectedRes: []models.Product{
				{
					Id:   1,
					Name: "jeans",
					Type: "clothes"},
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case - 1",
			mock: []*gomock.Call{
				mockStore.EXPECT().GetAllProducts(gomock.Any()).Return(nil, errors.Error("Error in database")),
			},
			expectedRes: nil,
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "2"},
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {

			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()
			res, err := mockHandler.GetProducts(ctx)

			if !reflect.DeepEqual(ts.expectedRes, res) {
				fmt.Print("expected ", ts.expectedRes, "obtained", res)
			}
			if !reflect.DeepEqual(ts.expectedErr, err) {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestInsertProductDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mockHandler := New(mockStore)

	testCases := []struct {
		desc        string
		product     models.Product
		mock        []*gomock.Call
		expectedErr error
	}{
		{
			desc: "Success case",
			product: models.Product{
				Id:   1,
				Name: "shirt",
				Type: "jeans",
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).Return(nil),
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case",
			product: models.Product{
				Id:   2,
				Name: "shirt",
				Type: "jeans",
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).Return(errors.Error("Error in database"))},
			expectedErr: errors.Error("Error in database"),
		},
		{
			desc: "Failure ID case",
			product: models.Product{
				Id:   0,
				Name: "shirt",
				Type: "jeans",
			},
			expectedErr: errors.InvalidParam{Param: []string{"id"}},
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {

			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()
			err := mockHandler.InsertProductDetails(ts.product, ctx)

			if !reflect.DeepEqual(ts.expectedErr, err) {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestUpdateProductDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mockHandler := New(mockStore)

	testCases := []struct {
		desc        string
		product     models.Product
		mock        []*gomock.Call
		expectedErr error
	}{
		{
			desc: "Success case",
			product: models.Product{
				Id:   1,
				Name: "shirt",
				Type: "jeans",
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(nil),
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case",
			product: models.Product{
				Id:   2,
				Name: "shirt",
				Type: "jeans",
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(errors.Error("Error in database"))},
			expectedErr: errors.Error("Error in database"),
		},
		{
			desc: "Failure ID case",
			product: models.Product{
				Id:   0,
				Name: "shirt",
				Type: "jeans",
			},
			expectedErr: errors.InvalidParam{Param: []string{"id"}},
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {

			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()
			err := mockHandler.UpdateProductDetails(ts.product, ctx)

			if !reflect.DeepEqual(ts.expectedErr, err) {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestDeleteProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mockHandler := New(mockStore)

	testCases := []struct {
		desc        string
		id          int
		mock        []*gomock.Call
		expectedErr error
	}{
		{
			desc: "Success case",
			id:   1,
			mock: []*gomock.Call{
				mockStore.EXPECT().DeleteById(gomock.Any(), gomock.Any()).Return(nil),
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case - 1",
			id:   2,
			mock: []*gomock.Call{
				mockStore.EXPECT().DeleteById(gomock.Any(), gomock.Any()).Return(errors.Error("Error in database"))},
			expectedErr: errors.Error("Error in database"),
		},
		{
			desc:        "Failure case - 2",
			id:          0,
			expectedErr: errors.InvalidParam{Param: []string{"id"}},
		},
	}
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {

			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()
			err := mockHandler.DeleteProductById(ts.id, ctx)
			if !reflect.DeepEqual(ts.expectedErr, err) {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

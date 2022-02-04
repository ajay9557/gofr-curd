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

func TestService_GetById(t *testing.T) {
	
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
				mockStore.EXPECT().GetProductById(gomock.Any(), 1).Return(models.Product{
					Id:   1,
					Name: "name-1",
					Type: "type-1",
				}, nil),
			},
				expOut: models.Product{
					Id:   1,
					Name: "name-1",
					Type: "type-1",
				},
				expErr: nil,
		},
		{
			desc:  "Case 2: Failure Case1",
			input: 22,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(models.Product{},
				errors.New("product not found for the given id")),
			},
			expOut: models.Product{},
			expErr: errors.New("product not found for the given id"),
		},
		{
			desc:  "Case 3: Failure Case2",
			input: 0,
			expOut: models.Product{},
			expErr: errors.New("invalid id"),
			},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, gofr.New())
			ctx.Context = context.Background()
			res, err := mockHandler.GetProductById(ctx, test.input)

			if !reflect.DeepEqual(test.expOut, res){
				fmt.Printf("expected : %v, Got : %v",test.expOut, res)
			}

			if !reflect.DeepEqual(test.expErr, err){
				fmt.Printf("expected : %v, Got : %v", test.expErr, err)
			}

		})

	}
}

func TestGetAllUsers(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProductStore(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct {
		desc string
		mockCall []*gomock.Call
		expOut []models.Product
		expErr error
	}{
		{
			desc: "Case 1: Success Case",
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetAllProducts(gomock.Any()).Return([]models.Product{
					{
						Id: 1,
						Name: "name-1",
						Type: "type-1",
					},
				},nil)},
			expOut: []models.Product{
				{
					Id: 1,
					Name: "name-1",
					Type: "type-1",
				},
			},
			expErr: nil,
		},
		{
			desc: "Case 2: Failure Case",
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetAllProducts(gomock.Any()).Return(errors.New("error while fetching all products details"))},
				expOut: nil,
				expErr: err1.EntityNotFound{Entity: "product", ID: "2"},
		},
	}

	for _, test := range testCases{
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil,nil,gofr.New())
			ctx.Context = context.Background()

			resp, err := mockHandler.GetAllProducts(ctx)

			if !reflect.DeepEqual(test.expOut, resp){
				fmt.Printf("expected : %v, Got : %v",test.expOut, resp)
			}

			if !reflect.DeepEqual(test.expErr, err){
				fmt.Printf("expected : %v, Got : %v", test.expErr, err)
			}
		})
	}
}

func TestService_UpdateProduct(t *testing.T){
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductStore(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct{
		desc string
		product models.Product
		mockCall []*gomock.Call
		ExpErr error
	}{
		{
			desc: "Case 1: Success case",
			product: models.Product{
				Id: 1,
				Name: "name-1",
				Type: "type-1",
			},
			mockCall: []*gomock.Call{mockStore.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(nil)},
			ExpErr: nil,
		},
		{
			desc: "Case 2: Failure Case1",
			product: models.Product{
				Id: 99,
				Name: "name-1",
				Type: "type-1",
			},
			mockCall: []*gomock.Call{mockStore.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(err1.Error("Error in database"))},
			ExpErr: err1.Error("Error in database"),
		},
		{
			desc: "Case 2: Failure case2",
			product: models.Product{
				Id:   0,
				Name: "name-1",
				Type: "type-1",
			},
			ExpErr: err1.InvalidParam{Param: []string{"id"}},
		},
	}

	for _, test := range testCases{
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil,nil,gofr.New())
			ctx.Context = context.Background()

			_, err := mockHandler.UpdateProduct(ctx, test.product)

			if !reflect.DeepEqual(err, test.ExpErr) {
				t.Error("Expected: ", test.ExpErr, "Got: ", err)
			}
		})
	}
}

func TestService_CreateProduct(t *testing.T){
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductStore(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct{
		desc     string
		product     models.Product
		mock     []*gomock.Call
		ExpErr error
	}{
		{
			desc: "Case 1: Success Case",
			product: models.Product{
				Id: 1,
				Name: "name-1",
				Type: "type-1",
			},
			mock:     []*gomock.Call{mockStore.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(nil)},
			ExpErr: nil,
		},
		{
			desc: "Case 1: Failure Case1",
			product: models.Product{
				Id: 2,
				Name: "name-1",
				Type: "type-1",
			},
			mock:     []*gomock.Call{mockStore.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(err1.Error("error while inserting new product"))},
			ExpErr: err1.Error("error while inserting new product"),
		},
		{
			desc: "Case 3: Failure Case2",
			product: models.Product{
				Id: 0,
				Name: "name-1",
				Type: "type-1",
			},
			ExpErr: err1.InvalidParam{Param: []string{"id"}},
		},

	}

	for _, test := range testCases{
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil,nil,gofr.New())
			ctx.Context = context.Background()

			_, err := mockHandler.CreateProduct(ctx, test.product)

			if !reflect.DeepEqual(err, test.ExpErr) {
				t.Error("Expected: ", test.ExpErr, "Got: ", err)
			}
		})
	}

}


func TestDeleteProduct(t *testing.T){
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductStore(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct {
		desc     string
		Id       int
		mock     []*gomock.Call
		ExpErr error
	}{
		{
			desc:     "Case 1: Success case",
			Id:       1,
			mock:     []*gomock.Call{mockStore.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(nil)},
			ExpErr: nil,
		},
		{
			desc:     "Case 2: Failure case1",
			Id:       2,
			mock:     []*gomock.Call{mockStore.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(err1.Error("error while deleting product"))},
			ExpErr: err1.Error("error while deleting product"),
		},
		{
			desc:        "Failure case - 2",
			Id:          0,
			ExpErr: err1.InvalidParam{Param: []string{"id"}},
		},
	}

	for _, test := range testCases{
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil,nil,gofr.New())
			ctx.Context = context.Background()

			err := mockHandler.DeleteProduct(ctx, test.Id)

			if !reflect.DeepEqual(err, test.ExpErr) {
				t.Error("Expected: ", test.ExpErr, "Got: ", err)
			}
		})
	}
}





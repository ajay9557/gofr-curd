package product

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	// "developer.zopsmart.com/go/gofr/pkg/errors"
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

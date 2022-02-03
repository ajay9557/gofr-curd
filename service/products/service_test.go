package products

import (
	"context"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/ridhdhish-desai-zs/product-gofr/store"
)

func TestGetById(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductStore := store.NewMockProduct(ctrl)
	productService := New(mockProductStore)

	product := models.Product{
		Id:       1,
		Name:     "mouse",
		Category: "electronics",
	}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc          string
		id            string
		expected      *models.Product
		expectedError error
		mockCall      *gomock.Call
	}{
		{
			desc:          "Case1",
			id:            "1",
			expected:      &product,
			expectedError: nil,
			mockCall:      mockProductStore.EXPECT().GetById(ctx, 1).Return(&product, nil),
		},
		{
			desc:          "Case2",
			id:            "abc",
			expected:      nil,
			expectedError: errors.EntityNotFound{Entity: "products", ID: "abc"},
			mockCall:      nil,
		},
		{
			desc:          "Case3",
			id:            "-1",
			expected:      nil,
			expectedError: errors.EntityNotFound{Entity: "products", ID: "-1"},
			mockCall:      nil,
		},
		{
			desc:          "Case4",
			id:            "1",
			expected:      nil,
			expectedError: errors.EntityNotFound{Entity: "products"},
			mockCall:      mockProductStore.EXPECT().GetById(ctx, 1).Return(nil, errors.EntityNotFound{Entity: "products"}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			p, err := productService.GetById(ctx, tc.id)

			if tc.expectedError != err {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(tc.expected, p) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, p)
			}
		})
	}
}

func TestGet(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductStore := store.NewMockProduct(ctrl)
	productService := New(mockProductStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	products := []*models.Product{
		{
			Id:       1,
			Name:     "mouse",
			Category: "electronics",
		},
	}

	tests := []struct {
		desc          string
		expected      []*models.Product
		expectedError error
		mockCall      *gomock.Call
	}{
		{
			desc: "Fetching all products",
			expected: []*models.Product{
				{
					Id:       1,
					Name:     "mouse",
					Category: "electronics",
				},
			},
			expectedError: nil,
			mockCall:      mockProductStore.EXPECT().Get(ctx).Return(products, nil),
		},
		{
			desc:          "Error while fetching products",
			expected:      nil,
			expectedError: errors.EntityNotFound{Entity: "products"},
			mockCall:      mockProductStore.EXPECT().Get(ctx).Return(nil, errors.EntityNotFound{Entity: "products"}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			products, err := productService.Get(ctx)

			if tc.expectedError != err {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(products, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, products)
			}
		})
	}
}

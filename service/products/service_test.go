package products

import (
	"context"
	"errors"
	"reflect"
	"testing"

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
		desc     string
		id       string
		expected *models.Product
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			id:       "1",
			expected: &product,
			mockCall: mockProductStore.EXPECT().GetById(ctx, 1).Return(&product, nil),
		},
		{
			desc:     "Case2",
			id:       "abc",
			expected: nil,
			mockCall: nil,
		},
		{
			desc:     "Case3",
			id:       "-1",
			expected: nil,
			mockCall: nil,
		},
		{
			desc:     "Case4",
			id:       "1",
			expected: nil,
			mockCall: mockProductStore.EXPECT().GetById(ctx, 1).Return(nil, errors.New("Invalid product id")),
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			p, _ := productService.GetById(ctx, tc.id)

			if !reflect.DeepEqual(tc.expected, p) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, p)
			}
		})
	}
}

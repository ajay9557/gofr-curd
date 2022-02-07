package products

import (
	"context"
	"reflect"
	"strconv"
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

func TestCreate(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProduct(ctrl)
	productService := New(mockStore)

	// Create test context here
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	// Test Inputs
	input1 := models.Product{
		Id:       1,
		Name:     "mouse",
		Category: "gaming",
	}

	tests := []struct {
		desc          string
		input         models.Product
		expected      *models.Product
		expectedError error
		mockCall      []*gomock.Call
	}{
		{
			desc:          "Success Case",
			input:         input1,
			expected:      &input1,
			expectedError: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().Create(ctx, input1).Return(nil),
				mockStore.EXPECT().GetById(ctx, 1).Return(&input1, nil),
			},
		},
		{
			desc:          "Empty request body",
			input:         models.Product{},
			expected:      nil,
			expectedError: errors.Error("Need Product data to create new product"),
			mockCall:      nil,
		},
		{
			desc:          "Error while creating",
			input:         input1,
			expected:      nil,
			expectedError: errors.Error("Connection lost"),
			mockCall: []*gomock.Call{
				mockStore.EXPECT().Create(ctx, input1).Return(errors.Error("Connection lost")),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			pr, err := productService.Create(ctx, tc.input)

			if err != tc.expectedError {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(pr, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, pr)
			}
		})
	}
}

func TestUpdateById(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProduct(ctrl)
	productService := New(mockStore)

	// Context creation
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	// Test inputs
	input1 := models.Product{
		Id:       1,
		Name:     "monitor",
		Category: "gaming",
	}

	tests := []struct {
		desc          string
		product       models.Product
		id            string
		expected      *models.Product
		expectedError error
		mockCall      []*gomock.Call
	}{
		{
			desc:          "Success Case",
			product:       input1,
			id:            "1",
			expected:      &input1,
			expectedError: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetById(ctx, 1).Return(&input1, nil),
				mockStore.EXPECT().UpdateById(ctx, 1, input1).Return(nil),
				mockStore.EXPECT().GetById(ctx, 1).Return(&input1, nil),
			},
		},
		{
			desc:          "Id must be numeric",
			product:       models.Product{},
			id:            "abc",
			expected:      nil,
			expectedError: errors.InvalidParam{Param: []string{"id"}},
			mockCall:      nil,
		},
		{
			desc:          "Id must be greater than 0",
			product:       models.Product{},
			id:            "-1",
			expected:      nil,
			expectedError: errors.InvalidParam{Param: []string{"id"}},
			mockCall:      nil,
		},
		{
			desc:          "Invalid id",
			product:       input1,
			id:            "100",
			expected:      nil,
			expectedError: errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(100)},
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetById(ctx, 100).Return(nil, errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(100)}),
			},
		},
		{
			desc:          "DB Connection lost",
			product:       input1,
			id:            "1",
			expected:      nil,
			expectedError: errors.Error("Connection lost"),
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetById(ctx, 1).Return(&input1, nil),
				mockStore.EXPECT().UpdateById(ctx, 1, input1).Return(errors.Error("Connection lost")),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			pr, err := productService.UpdateById(ctx, tc.id, input1)

			if !reflect.DeepEqual(tc.expectedError, err) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}

			if !reflect.DeepEqual(pr, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, pr)
			}
		})
	}
}

func TestDeleteById(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProduct(ctrl)
	productService := New(mockStore)

	// Context creation
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	input1 := models.Product{
		Id:       1,
		Name:     "monitor",
		Category: "gaming",
	}

	tests := []struct {
		desc          string
		id            string
		expectedError error
		mockCall      []*gomock.Call
	}{
		{
			desc:          "Success Case",
			id:            "1",
			expectedError: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetById(ctx, 1).Return(&input1, nil),
				mockStore.EXPECT().DeleteById(ctx, 1).Return(nil),
			},
		},
		{
			desc:          "Id must be numeric",
			id:            "abc",
			expectedError: errors.InvalidParam{Param: []string{"id"}},
			mockCall:      nil,
		},
		{
			desc:          "Id must be greater than 0",
			id:            "-1",
			expectedError: errors.InvalidParam{Param: []string{"id"}},
			mockCall:      nil,
		},
		{
			desc:          "Invalid id",
			id:            "100",
			expectedError: errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(100)},
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetById(ctx, 100).Return(nil, errors.EntityNotFound{Entity: "products", ID: strconv.Itoa(100)}),
			},
		},
		{
			desc:          "DB Connection lost",
			id:            "1",
			expectedError: errors.Error("Connection lost"),
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetById(ctx, 1).Return(&input1, nil),
				mockStore.EXPECT().DeleteById(ctx, 1).Return(errors.Error("Connection lost")),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := productService.DeleteById(ctx, tc.id)

			if !reflect.DeepEqual(tc.expectedError, err) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}
		})
	}
}

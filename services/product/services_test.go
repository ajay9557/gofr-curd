package product

import (
	"context"
	"reflect"

	"gofr-curd/models"
	"gofr-curd/stores"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := gofr.New()

	mockUserStore := stores.NewMockIstore(ctrl)
	testUserService := New(mockUserStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc        string
		id          string
		expected    *models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc:        "Case1",
			id:          "1",
			expected:    &models.Product{Id: 1, Name: "daikinn", Type: "AC"},
			expectedErr: nil,
			mockCall:    mockUserStore.EXPECT().GetProductByID(ctx, 1).Return(&models.Product{Id: 1, Name: "daikinn", Type: "AC"}, nil),
		},
		{
			desc:        "Case2",
			id:          "100",
			expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall: mockUserStore.EXPECT().
				GetProductByID(ctx, 100).
				Return(&models.Product{}, errors.EntityNotFound{Entity: "products", ID: "100"}),
		},
		{
			desc:        "Case3",
			id:          "anusri",
			expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			mockCall:    nil,
		},

		{
			desc:        "Case4",
			id:          "-100",
			expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			mockCall:    nil,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			p, err := testUserService.GetProductByID(ctx, test.id)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)

			if err == nil && !reflect.DeepEqual(test.expected, p) {
				t.Errorf("Expected: %v, Got: %v", test.expected, p)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := gofr.New()

	mockUserStore := stores.NewMockIstore(ctrl)
	testUserService := New(mockUserStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc string
		// id          string
		expected    []*models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc: "Case1",

			expected: []*models.Product{{Id: 1, Name: "daikin", Type: "AC"},
				{Id: 2, Name: "milton", Type: "Water Bottle"}},
			expectedErr: nil,
			mockCall: mockUserStore.EXPECT().GetAllProducts(ctx).Return([]*models.Product{{Id: 1, Name: "daikin", Type: "AC"},
				{Id: 2, Name: "milton", Type: "Water Bottle"}}, nil),
		},
		{
			desc: "Case2",

			expected:    []*models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products"},
			mockCall: mockUserStore.EXPECT().
				GetAllProducts(ctx).
				Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			p, err := testUserService.GetAllProducts(ctx)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)

			if err == nil && !reflect.DeepEqual(test.expected, p) {
				t.Errorf("Expected: %v, Got: %v", test.expected, p)
			}
		})
	}
}

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := gofr.New()

	mockUserStore := stores.NewMockIstore(ctrl)
	testUserService := New(mockUserStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc string
		// id          string
		input       models.Product
		expected    *models.Product
		expectedErr error
		mockCall    []*gomock.Call
	}{
		{
			desc:        "Case1",
			input:       models.Product{Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			expectedErr: nil,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().
					CreateProduct(ctx, models.Product{Name: "milton", Type: "Water Bottle"}).
					Return( /*&models.Product{}*/ 2, nil),
				mockUserStore.EXPECT().
					GetProductByID(ctx, 2).
					Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),
			},
		},
		{
			desc:        "Case2",
			input:       models.Product{Name: "", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Given Empty data"),
			mockCall:    nil,
		},
		{
			desc:        "Case3",
			input:       models.Product{Id: 2, Name: "", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Name"),
			mockCall:    nil,
		},
		{
			desc:        "Case4",
			input:       models.Product{Id: 2, Name: "milton", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Type"),
			mockCall:    nil,
		},

		{
			desc:        "Case5",
			input:       models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.Error("Couldn't execute query"),
			mockCall: []*gomock.Call{mockUserStore.EXPECT().
				CreateProduct(ctx, models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}).
				Return( /*&models.Product{}*/ 1, errors.Error("Couldn't execute query"))},
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			p, err := testUserService.CreateProduct(ctx, test.input)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)

			if err == nil && !reflect.DeepEqual(test.expected, p) {
				t.Errorf("Expected: %v, Got: %v", test.expected, p)
			}
		})
	}
}
func TestDeleteById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUserStore := stores.NewMockIstore(ctrl)
	testUserService := New(mockUserStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc string
		id   string
		// expected    *models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc: "Case1",
			id:   "1",
			// expected:    &models.Product{Id: 1, Name: "daikinn", Type: "AC"},
			expectedErr: nil,
			mockCall:    mockUserStore.EXPECT().DeleteByID(ctx, 1).Return(nil),
		},
		{
			desc: "Case2",
			id:   "100",
			// expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall:    mockUserStore.EXPECT().DeleteByID(ctx, 100).Return(errors.EntityNotFound{Entity: "products", ID: "100"}),
		},
		{
			desc: "Case3",
			id:   "anusri",
			// expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			mockCall:    nil,
		},

		{
			desc: "Case4",
			id:   "-100",
			// expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			mockCall:    nil,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			err := testUserService.DeleteByID(ctx, test.id)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)
		})
	}
}
func TestUpdateById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUserStore := stores.NewMockIstore(ctrl)
	testUserService := New(mockUserStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc        string
		id          string
		input       models.Product
		expected    *models.Product
		expectedErr error
		mockCall    []*gomock.Call
	}{
		{
			desc:        "Case1",
			id:          "2",
			input:       models.Product{Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			expectedErr: nil,
			mockCall: []*gomock.Call{mockUserStore.EXPECT().UpdateByID(ctx, 2, models.Product{Name: "milton", Type: "Water Bottle"}).Return(2, nil),
				mockUserStore.EXPECT().GetProductByID(ctx, 2).Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),
			},
		},
		{
			desc:     "Case2",
			id:       "100",
			input:    models.Product{Name: "milton", Type: "Water Bottle"},
			expected: &models.Product{},
			expectedErr:/* errors.EntityNotFound{Entity: "products", ID: "100"}*/ errors.Error("Couldn't execute query"),
			mockCall: []*gomock.Call{mockUserStore.EXPECT().
				UpdateByID(ctx, 100, models.Product{Name: "milton", Type: "Water Bottle"}).
				Return(0, errors.Error("Couldn't execute query"))},
		},

		{
			desc:        "Case3",
			id:          "anusri",
			input:       models.Product{Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			mockCall:    nil,
		},

		{
			desc:        "Case4",
			id:          "-100",
			input:       models.Product{Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			mockCall:    nil,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			p, err := testUserService.UpdateByID(ctx, test.id, test.input)

			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)

			if err == nil && !reflect.DeepEqual(test.expected, p) {
				t.Errorf("Expected: %v, Got: %v", test.expected, p)
			}
		})
	}
}

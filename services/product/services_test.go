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
)

func TestGetProductById(t *testing.T) {
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
		expected    *models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc:        "Case1",
			id:          "1",
			expected:    &models.Product{Id: 1, Name: "daikinn", Type: "AC"},
			expectedErr: nil,
			mockCall:    mockUserStore.EXPECT().GetProductById(ctx, 1).Return(&models.Product{Id: 1, Name: "daikinn", Type: "AC"}, nil),
		},
		{
			desc:        "Case2",
			id:          "100",
			expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall:    mockUserStore.EXPECT().GetProductById(ctx, 100).Return(&models.Product{}, errors.EntityNotFound{Entity: "products", ID: "100"}),
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

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			p, err := testUserService.GetProductById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, p) {
				t.Errorf("Expected: %v, Got: %v", test.expected, p)
			}

		})

	}

}

func TestGetAllProducts(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

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

			expected: []*models.Product{&models.Product{Id: 1, Name: "daikin", Type: "AC"},
				&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}},
			expectedErr: nil,
			mockCall: mockUserStore.EXPECT().GetAllProducts(ctx).Return([]*models.Product{&models.Product{Id: 1, Name: "daikin", Type: "AC"},
				&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}}, nil),
		},
		{
			desc: "Case2",

			expected:    []*models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products"},
			mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			p, err := testUserService.GetAllProducts(ctx)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, p) {
				t.Errorf("Expected: %v, Got: %v", test.expected, p)
			}

		})

	}

}

func TestCreateProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

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
			mockCall: []*gomock.Call{mockUserStore.EXPECT().CreateProduct(ctx, models.Product{Name: "milton", Type: "Water Bottle"}).Return( /*&models.Product{}*/ 2, nil),
				mockUserStore.EXPECT().GetProductById(ctx, 2).Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),
			},
			// mockCall: mockUserStore.EXPECT().CreateProduct(ctx, models.Product{Name: "milton", Type: "Water Bottle"}).Return( /*&models.Product{}*/ 2, nil),
		},
		{
			desc:        "Case2",
			input:       models.Product{Name: "", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Given Empty data"),
			mockCall:    nil,
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
		{
			desc:        "Case3",
			input:       models.Product{Id: 2, Name: "", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Name"),
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
			mockCall: nil,
		},
		{
			desc:        "Case4",
			input:       models.Product{Id: 2, Name: "milton", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Type"),
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
			mockCall: nil,
		},

		{
			desc:        "Case5",
			input:       models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.Error("Couldn't execute query"),
			mockCall:    []*gomock.Call{mockUserStore.EXPECT().CreateProduct(ctx, models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}).Return( /*&models.Product{}*/ 1, errors.Error("Couldn't execute query"))},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			p, err := testUserService.CreateProduct(ctx, test.input)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
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
			mockCall:    mockUserStore.EXPECT().DeleteById(ctx, 1).Return(nil),
		},
		{
			desc: "Case2",
			id:   "100",
			// expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall:    mockUserStore.EXPECT().DeleteById(ctx, 100).Return(errors.EntityNotFound{Entity: "products", ID: "100"}),
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

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			err := testUserService.DeleteById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			// if err == nil && !reflect.DeepEqual(test.expected, p) {
			// 	t.Errorf("Expected: %v, Got: %v", test.expected, p)
			// }

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
			mockCall: []*gomock.Call{mockUserStore.EXPECT().UpdateById(ctx, 2, models.Product{Name: "milton", Type: "Water Bottle"}).Return(2, nil),
				mockUserStore.EXPECT().GetProductById(ctx, 2).Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),
			},
		},
		{
			desc:     "Case2",
			id:       "100",
			input:    models.Product{Name: "milton", Type: "Water Bottle"},
			expected: &models.Product{},
			expectedErr:/* errors.EntityNotFound{Entity: "products", ID: "100"}*/ errors.Error("Couldn't execute query"),
			mockCall: []*gomock.Call{mockUserStore.EXPECT().UpdateById(ctx, 100, models.Product{Name: "milton", Type: "Water Bottle"}).Return(0, errors.Error("Couldn't execute query"))},
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
		{
			desc:        "Case5",
			id:          "4",
			input:       models.Product{Name: "", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Given Empty data"),
			mockCall:    nil,
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
		{
			desc:        "Case6",
			id:          "4",
			input:       models.Product{Name: "", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Name"),
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
			mockCall: nil,
		},
		{
			desc:        "Case7",
			id:          "4",
			input:       models.Product{Name: "milton", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Type"),
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
			mockCall: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			p, err := testUserService.UpdateById(ctx, test.id, test.input)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, p) {
				t.Errorf("Expected: %v, Got: %v", test.expected, p)
			}

		})

	}

}

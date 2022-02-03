package product

import (
	"context"
	"reflect"

	"productGofr/models"
	"productGofr/stores"
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
			mockCall:    mockUserStore.EXPECT().UserById(ctx, 1).Return(&models.Product{Id: 1, Name: "daikinn", Type: "AC"}, nil),
		},
		{
			desc:        "Case2",
			id:          "100",
			expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall:    mockUserStore.EXPECT().UserById(ctx, 100).Return(&models.Product{}, errors.EntityNotFound{Entity: "products", ID: "100"}),
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

func TestGetAllUsers(t *testing.T) {
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
			mockCall: mockUserStore.EXPECT().GetAllUsers(ctx).Return([]*models.Product{&models.Product{Id: 1, Name: "daikin", Type: "AC"},
				&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}}, nil),
		},
		{
			desc: "Case2",

			expected:    []*models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products"},
			mockCall:    mockUserStore.EXPECT().GetAllUsers(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			p, err := testUserService.GetAllUsers(ctx)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			if err == nil && !reflect.DeepEqual(test.expected, p) {
				t.Errorf("Expected: %v, Got: %v", test.expected, p)
			}

		})

	}

}

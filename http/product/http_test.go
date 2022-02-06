package product

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"reflect"

	"gofr-curd/models"
	"gofr-curd/services"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
)

func TestGetProductById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := services.NewMockIservice(ctrl)
	testhndlr := Handler{mockUserService}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc        string
		id          string
		expected    *models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{

		{desc: "Case1",
			id:          "1",
			expected:    &models.Product{Id: 1, Name: "daikinn", Type: "AC"},
			expectedErr: nil,
			// mockCall:    mockUserStore.EXPECT().UserById(ctx, 1).Return(&models.Product{Id: 1, Name: "daikinn", Type: "AC"}, nil),
			mockCall: mockUserService.EXPECT().GetProductById(gomock.Any(), "1").Return(&models.Product{Id: 1, Name: "daikinn", Type: "AC"}, nil),
		},
		{
			desc:        "Case2",
			id:          "100",
			expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			// mockCall:    mockUserStore.EXPECT().UserById(ctx, 100).Return(&models.Product{}, errors.EntityNotFound{Entity: "products", ID: "100"}),
			mockCall: mockUserService.EXPECT().GetProductById(gomock.Any(), "100").Return(&models.Product{}, errors.EntityNotFound{Entity: "products", ID: "100"}),
		},
		{
			desc:        "Case3",
			id:          "anusri",
			expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			mockCall:    mockUserService.EXPECT().GetProductById(gomock.Any(), "anusri").Return(&models.Product{}, errors.MissingParam{Param: []string{"anusri"}}),
		},

		{
			desc:        "Case4",
			id:          "-100",
			expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().GetProductById(gomock.Any(), "-100").Return(&models.Product{}, errors.InvalidParam{Param: []string{"-100"}}),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			r := httptest.NewRequest( /*http.MethodGet*/ "GET", "/products/{id}", nil)
			w := httptest.NewRecorder()

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			// req = mux.SetURLVars(req, map[string]string{
			// 	"id": test.id,
			// })
			ctx.SetPathParams(map[string]string{
				"id": test.id,
			})

			// p, err := testhndlr.GetByIdHandler(ctx)
			_, err := testhndlr.GetProductByIdHandler(ctx)
			// p, err := testUserService.GetProductById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			// if err == nil && !reflect.DeepEqual(test.expected, p) {
			// 	t.Errorf("Expected: %v, Got: %v", test.expected, p)
			// }

		})
	}
}

func TestCreateProductHandler(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := services.NewMockIservice(ctrl)
	testhndlr := Handler{mockUserService}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc string
		// id          string
		expected    *models.Product
		input       models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc:        "Case1",
			input:       models.Product{Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			expectedErr: nil,
			mockCall:    mockUserService.EXPECT().CreateProduct(gomock.Any(), models.Product{Name: "milton", Type: "Water Bottle"}).Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),
			// mockUserStore.EXPECT().GetProductById(ctx, 2).Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),
			// mockCall: mockUserStore.EXPECT().CreateProduct(ctx, models.Product{Name: "milton", Type: "Water Bottle"}).Return( /*&models.Product{}*/ 2, nil),
		},
		{
			desc:        "Case2",
			input:       models.Product{Name: "", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Given Empty data"),
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().CreateProduct(gomock.Any(), models.Product{Name: "", Type: ""}).Return(&models.Product{}, errors.Error("Given Empty data")),
		},
		{
			desc:        "Case3",
			input:       models.Product{Id: 2, Name: "", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Name"),
			mockCall:    mockUserService.EXPECT().CreateProduct(gomock.Any(), models.Product{Id: 2, Name: "", Type: "Water Bottle"}).Return(&models.Product{}, errors.Error("Please provide Data for Name")),
			// mockCall: nil,
		},
		{
			desc:        "Case4",
			input:       models.Product{Id: 2, Name: "milton", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Type"),
			mockCall:    mockUserService.EXPECT().CreateProduct(gomock.Any(), models.Product{Id: 2, Name: "milton", Type: ""}).Return(&models.Product{}, errors.Error("Please provide Data for Type")),
			// mockCall: nil,
		},

		{
			desc:        "Case5",
			input:       models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.Error("Couldn't execute query"),
			mockCall:    mockUserService.EXPECT().CreateProduct(gomock.Any(), models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}).Return(&models.Product{}, errors.Error("Couldn't execute query")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			file, _ := json.Marshal(test.input)

			r := httptest.NewRequest( /*http.MethodGet*/ "CREATE", "/products", bytes.NewReader(file))
			w := httptest.NewRecorder()

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			// p, err := testhndlr.GetProductByIdHandler(ctx)
			_, err := testhndlr.CreateProductHandler(ctx)
			// p, err := testUserService.GetProductById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			// if err == nil && !reflect.DeepEqual(test.expected, p) {
			// 	t.Errorf("Expected: %v, Got: %v", test.expected, p)
			// }

		})
	}
}

func TestGetAllProductsHandler(t *testing.T) {

	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := services.NewMockIservice(ctrl)
	testhndlr := Handler{mockUserService}

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
			mockCall: mockUserService.EXPECT().GetAllProducts(gomock.Any()).Return([]*models.Product{&models.Product{Id: 1, Name: "daikin", Type: "AC"},
				&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}}, nil),
		},
		{
			desc: "Case2",

			expected:    []*models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products"},
			mockCall:    mockUserService.EXPECT().GetAllProducts(gomock.Any()).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			r := httptest.NewRequest( /*http.MethodGet*/ "GET", "/products", nil)
			w := httptest.NewRecorder()

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			// p, err := testhndlr.GetProductByIdHandler(ctx)
			_, err := testhndlr.GetAllProductsHandler(ctx)
			// p, err := testUserService.GetProductById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			// if err == nil && !reflect.DeepEqual(test.expected, p) {
			// 	t.Errorf("Expected: %v, Got: %v", test.expected, p)
			// }

		})
	}

}

func TestDeleteByIdHandler(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := services.NewMockIservice(ctrl)
	testhndlr := Handler{mockUserService}

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
			// mockCall:    mockUserStore.EXPECT().GetProductById(ctx, 1).Return(&models.Product{Id: 1, Name: "daikinn", Type: "AC"}, nil),
			mockCall: mockUserService.EXPECT().DeleteById(gomock.Any(), "1").Return(nil),
		},
		{
			desc: "Case2",
			id:   "100",
			// expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			// mockCall:    mockUserStore.EXPECT().GetProductById(ctx, 100).Return(&models.Product{}, errors.EntityNotFound{Entity: "products", ID: "100"}),
			mockCall: mockUserService.EXPECT().DeleteById(gomock.Any(), "100").Return(errors.EntityNotFound{Entity: "products", ID: "100"}),
		},
		{
			desc: "Case3",
			id:   "anusri",
			// expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			mockCall:    mockUserService.EXPECT().DeleteById(gomock.Any(), "anusri").Return(errors.MissingParam{Param: []string{"anusri"}}),
		},

		{
			desc: "Case4",
			id:   "-100",
			// expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().DeleteById(gomock.Any(), "-100").Return(errors.InvalidParam{Param: []string{"-100"}}),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			r := httptest.NewRequest( /*http.MethodGet*/ "DELETE", "/products/{id}", nil)
			w := httptest.NewRecorder()

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			// req = mux.SetURLVars(req, map[string]string{
			// 	"id": test.id,
			// })
			ctx.SetPathParams(map[string]string{
				"id": test.id,
			})

			// p, err := testhndlr.GetProductByIdHandler(ctx)
			_, err := testhndlr.DeleteByIdHandler(ctx)
			// p, err := testUserService.GetProductById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			// if err == nil && !reflect.DeepEqual(test.expected, p) {
			// 	t.Errorf("Expected: %v, Got: %v", test.expected, p)
			// }

		})
	}

}

func TestUpdateByIdHandler(t *testing.T) {

	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := services.NewMockIservice(ctrl)
	testhndlr := Handler{mockUserService}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc        string
		id          string
		input       models.Product
		expected    *models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc:        "Case1",
			id:          "2",
			input:       models.Product{Name: "miltonn", Type: "Water Bottlee"},
			expected:    &models.Product{Id: 2, Name: "miltonn", Type: "Water Bottlee"},
			expectedErr: nil,
			mockCall:    mockUserService.EXPECT().UpdateById(gomock.Any(), "2", models.Product{Name: "miltonn", Type: "Water Bottlee"}).Return(&models.Product{Id: 2, Name: "miltonn", Type: "Water Bottlee"}, nil),
			// mockUserService.EXPECT().GetProductById(ctx, 2).Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),

		},
		{
			desc:     "Case2",
			id:       "100",
			input:    models.Product{Name: "milton", Type: "Water Bottle"},
			expected: &models.Product{},
			expectedErr:/* errors.EntityNotFound{Entity: "products", ID: "100"}*/ errors.Error("Couldn't execute query"),
			mockCall: mockUserService.EXPECT().UpdateById(gomock.Any(), "100", models.Product{Name: "milton", Type: "Water Bottle"}).Return(&models.Product{}, errors.Error("Couldn't execute query")),
		},

		{
			desc:        "Case3",
			id:          "anusri",
			input:       models.Product{Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().UpdateById(gomock.Any(), "anusri", models.Product{Name: "milton", Type: "Water Bottle"}).Return(&models.Product{}, errors.MissingParam{Param: []string{"anusri"}}),
		},

		{
			desc:        "Case4",
			id:          "-100",
			input:       models.Product{Name: "milton", Type: "Water Bottle"},
			expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().UpdateById(gomock.Any(), "-100", models.Product{Name: "milton", Type: "Water Bottle"}).Return(&models.Product{}, errors.InvalidParam{Param: []string{"-100"}}),
		},
		{
			desc:        "Case5",
			id:          "4",
			input:       models.Product{Name: "", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Given Empty data"),
			mockCall:    mockUserService.EXPECT().UpdateById(gomock.Any(), "4", models.Product{Name: "", Type: ""}).Return(&models.Product{}, errors.Error("Given Empty data")),
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
		{
			desc:        "Case6",
			id:          "4",
			input:       models.Product{Name: "", Type: "Water Bottlee"},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Name"),
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
			mockCall: mockUserService.EXPECT().UpdateById(gomock.Any(), "4", models.Product{Name: "", Type: "Water Bottlee"}).Return(&models.Product{}, errors.Error("Please provide Data for Name")),
		},
		{
			desc:        "Case7",
			id:          "4",
			input:       models.Product{Name: "miltonn", Type: ""},
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Type"),
			// mockCall:    mockUserStore.EXPECT().GetAllProducts(ctx).Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
			mockCall: mockUserService.EXPECT().UpdateById(gomock.Any(), "4", models.Product{Name: "miltonn", Type: ""}).Return(&models.Product{}, errors.Error("Please provide Data for Type")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			file, _ := json.Marshal(test.input)
			r := httptest.NewRequest( /*http.MethodGet*/ "UPDATE", "/products/{id}", bytes.NewReader(file))
			w := httptest.NewRecorder()

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			ctx.SetPathParams(map[string]string{
				"id": test.id,
			})

			// p, err := testhndlr.GetProductByIdHandler(ctx)
			_, err := testhndlr.UpdateByIdHandler(ctx)
			// p, err := testUserService.GetProductById(ctx, test.id)
			if !reflect.DeepEqual(err, test.expectedErr) {
				t.Error("expected: ", test.expectedErr, "obtained: ", err)
			}
			// if err == nil && !reflect.DeepEqual(test.expected, p) {
			// 	t.Errorf("Expected: %v, Got: %v", test.expected, p)
			// }

		})
	}

}

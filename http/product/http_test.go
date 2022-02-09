package product

import (
	"bytes"
	"context"
	"gofr-curd/models"
	"gofr-curd/services"
	"net/http/httptest"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func MockHTTP(app *gofr.Gofr, method string) *gofr.Context {
	r := httptest.NewRequest( /*http.MethodGet*/ method, "/products/{id}", nil)
	w := httptest.NewRecorder()

	req := request.NewHTTPRequest(r)
	res := responder.NewContextualResponder(w, r)

	ctx := gofr.NewContext(res, req, app)

	return ctx
}

func MockHTTP2(app *gofr.Gofr, method string, body []byte) *gofr.Context {
	r := httptest.NewRequest( /*http.MethodGet*/ method, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()
	req := request.NewHTTPRequest(r)
	res := responder.NewContextualResponder(w, r)

	ctx := gofr.NewContext(res, req, app)

	return ctx
}

func TestGetProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := gofr.New()

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
			mockCall: mockUserService.EXPECT().
				GetProductByID(gomock.Any(), "1").
				Return(&models.Product{Id: 1, Name: "daikinn", Type: "AC"}, nil),
		},
		{
			desc:        "Case2",
			id:          "100",
			expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall: mockUserService.EXPECT().
				GetProductByID(gomock.Any(), "100").
				Return(&models.Product{}, errors.EntityNotFound{Entity: "products", ID: "100"}),
		},
		{
			desc:        "Case3",
			id:          "anusri",
			expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			mockCall: mockUserService.EXPECT().
				GetProductByID(gomock.Any(), "anusri").
				Return(&models.Product{}, errors.MissingParam{Param: []string{"anusri"}}),
		},

		{
			desc:        "Case4",
			id:          "-100",
			expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().
				GetProductByID(gomock.Any(), "-100").
				Return(&models.Product{}, errors.InvalidParam{Param: []string{"-100"}}),
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := MockHTTP(app, "GET")

			ctx.SetPathParams(map[string]string{
				"id": test.id,
			})

			_, err := testhndlr.GetProductByIDHandler(ctx)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)
		})
	}
}

func TestCreateProductHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := gofr.New()

	mockUserService := services.NewMockIservice(ctrl)
	testhndlr := Handler{mockUserService}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc string
		// id          string
		expected *models.Product
		// input       models.Product
		input       []byte
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc: "Case1",
			// input:       models.Product{Name: "milton", Type: "Water Bottle"},
			input:       []byte(`{"name":"milton","type":"Water Bottle"}`),
			expected:    &models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			expectedErr: nil,
			mockCall: mockUserService.EXPECT().
				CreateProduct(gomock.Any(), models.Product{Name: "milton", Type: "Water Bottle"}).
				Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),
		},
		{
			desc: "Case2",
			// input:       models.Product{Name: "", Type: ""},
			input:       []byte(`{"name":"","type":""}`),
			expected:    &models.Product{},
			expectedErr: errors.Error("Given Empty data"),
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().
				CreateProduct(gomock.Any(), models.Product{Name: "", Type: ""}).
				Return(&models.Product{}, errors.Error("Given Empty data")),
		},
		{
			desc: "Case3",
			// input:       models.Product{Id: 2, Name: "", Type: "Water Bottle"},
			input:       []byte(`{"name":"","type":"Water Bottle"}`),
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Name"),
			mockCall: mockUserService.EXPECT().
				CreateProduct(gomock.Any(), models.Product{Name: "", Type: "Water Bottle"}).
				Return(&models.Product{}, errors.Error("Please provide Data for Name")),
			// mockCall: nil,
		},
		{
			desc: "Case4",
			// input:       models.Product{Id: 2, Name: "milton", Type: ""},
			input:       []byte(`{"name":"milton","type":""}`),
			expected:    &models.Product{},
			expectedErr: errors.Error("Please provide Data for Type"),
			mockCall: mockUserService.EXPECT().
				CreateProduct(gomock.Any(), models.Product{Name: "milton", Type: ""}).
				Return(&models.Product{}, errors.Error("Please provide Data for Type")),
			// mockCall: nil,
		},

		{
			desc: "Case5",
			// input:       models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			input:       []byte(`{"name":"milton","type":"Water Bottle"}`),
			expected:    &models.Product{},
			expectedErr: errors.Error("Couldn't execute query"),
			mockCall: mockUserService.EXPECT().
				CreateProduct(gomock.Any(), models.Product{Name: "milton", Type: "Water Bottle"}).
				Return(&models.Product{}, errors.Error("Couldn't execute query")),
		},
		{
			desc: "Case6",
			// input:       models.Product{Id: 2, Name: "milton", Type: "Water Bottle"},
			input:       []byte(`{"Some unchangeable data"}`),
			expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"body"}},
			mockCall:    nil,
			// mockCall:    mockUserService.EXPECT().
			// CreateProduct(gomock.Any(), models.Product{Name: "milton", Type: "Water Bottle"}).
			// Return(&models.Product{}, errors.Error("Couldn't execute query")),
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := MockHTTP2(app, "CREATE", test.input)

			_, err := testhndlr.CreateProductHandler(ctx)
			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)
		})
	}
}
func TestGetAllProductsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := gofr.New()

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

			expected: []*models.Product{{Id: 1, Name: "daikin", Type: "AC"},
				{Id: 2, Name: "milton", Type: "Water Bottle"}},
			expectedErr: nil,
			mockCall: mockUserService.EXPECT().GetAllProducts(gomock.Any()).Return([]*models.Product{{Id: 1, Name: "daikin", Type: "AC"},
				{Id: 2, Name: "milton", Type: "Water Bottle"}}, nil),
		},
		{
			desc: "Case2",

			expected:    []*models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products"},
			mockCall: mockUserService.EXPECT().
				GetAllProducts(gomock.Any()).
				Return( /*&models.Product{}*/ []*models.Product{}, errors.EntityNotFound{Entity: "products"}),
		},
	}
	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			r := httptest.NewRequest( /*http.MethodGet*/ "GET", "/products", nil)
			w := httptest.NewRecorder()

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			_, err := testhndlr.GetAllProductsHandler(ctx)

			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)
		})
	}
}

func TestDeleteByIdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := gofr.New()

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
			mockCall: mockUserService.EXPECT().DeleteByID(gomock.Any(), "1").Return(nil),
		},
		{
			desc: "Case2",
			id:   "100",
			// expected:    &models.Product{},
			expectedErr: errors.EntityNotFound{Entity: "products", ID: "100"},
			mockCall: mockUserService.EXPECT().
				DeleteByID(gomock.Any(), "100").
				Return(errors.EntityNotFound{Entity: "products", ID: "100"}),
		},
		{
			desc: "Case3",
			id:   "anusri",
			// expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			mockCall: mockUserService.EXPECT().
				DeleteByID(gomock.Any(), "anusri").
				Return(errors.MissingParam{Param: []string{"anusri"}}),
		},

		{
			desc: "Case4",
			id:   "-100",
			// expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().DeleteByID(gomock.Any(), "-100").Return(errors.InvalidParam{Param: []string{"-100"}}),
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			ctx := MockHTTP(app, "DELETE")
			ctx.SetPathParams(map[string]string{
				"id": test.id,
			})
			_, err := testhndlr.DeleteByIDHandler(ctx)

			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)
		})
	}
}
func TestUpdateByIdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := gofr.New()

	mockUserService := services.NewMockIservice(ctrl)
	testhndlr := Handler{mockUserService}

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	tests := []struct {
		desc string
		id   string
		// input       models.Product
		input       []byte
		expected    *models.Product
		expectedErr error
		mockCall    *gomock.Call
	}{
		{
			desc: "Case1",
			id:   "2",
			// input:       models.Product{Name: "miltonn", Type: "Water Bottlee"},
			input:       []byte(`{"name":"miltonn","type":"Water Bottlee"}`),
			expected:    &models.Product{Id: 2, Name: "miltonn", Type: "Water Bottlee"},
			expectedErr: nil,
			mockCall: mockUserService.EXPECT().
				UpdateByID(gomock.Any(), "2", models.Product{Name: "miltonn", Type: "Water Bottlee"}).
				Return(&models.Product{Id: 2, Name: "miltonn", Type: "Water Bottlee"}, nil),
			// mockUserService.EXPECT().GetProductById(ctx, 2).Return(&models.Product{Id: 2, Name: "milton", Type: "Water Bottle"}, nil),

		},
		{
			desc: "Case2",
			id:   "100",
			// input:    models.Product{Name: "milton", Type: "Water Bottle"},
			input:    []byte(`{"name":"milton","type":"Water Bottle"}`),
			expected: &models.Product{},
			expectedErr:/* errors.EntityNotFound{Entity: "products", ID: "100"}*/ errors.Error("Couldn't execute query"),
			mockCall: mockUserService.EXPECT().
				UpdateByID(gomock.Any(), "100", models.Product{Name: "milton", Type: "Water Bottle"}).
				Return(&models.Product{}, errors.Error("Couldn't execute query")),
		},

		{
			desc: "Case3",
			id:   "anusri",
			// input:       models.Product{Name: "milton", Type: "Water Bottle"},
			input:       []byte(`{"name":"milton","type":"Water Bottle"}`),
			expected:    &models.Product{},
			expectedErr: errors.MissingParam{Param: []string{"anusri"}},
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().
				UpdateByID(gomock.Any(), "anusri", models.Product{Name: "milton", Type: "Water Bottle"}).
				Return(&models.Product{}, errors.MissingParam{Param: []string{"anusri"}}),
		},

		{
			desc: "Case4",
			id:   "-100",
			// input:       models.Product{Name: "milton", Type: "Water Bottle"},
			input:       []byte(`{"name":"milton","type":"Water Bottle"}`),
			expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"-100"}},
			// mockCall:    nil,
			mockCall: mockUserService.EXPECT().
				UpdateByID(gomock.Any(), "-100", models.Product{Name: "milton", Type: "Water Bottle"}).
				Return(&models.Product{}, errors.InvalidParam{Param: []string{"-100"}}),
		},
		{
			desc:        "Case5",
			id:          "2",
			input:       []byte(`{"Some unchangeable data"}`),
			expected:    &models.Product{},
			expectedErr: errors.InvalidParam{Param: []string{"body"}},
			mockCall:    nil,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.desc, func(t *testing.T) {
			r := httptest.NewRequest( /*http.MethodGet*/ "UPDATE", "/products/{id}", bytes.NewReader(test.input))
			w := httptest.NewRecorder()

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			ctx.SetPathParams(map[string]string{
				"id": test.id,
			})

			_, err := testhndlr.UpdateByIDHandler(ctx)

			assert.Equal(t, err, test.expectedErr, "%s, failed.\n", test.desc)
		})
	}
}

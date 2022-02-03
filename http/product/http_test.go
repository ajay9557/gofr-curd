package product

import (
	"context"
	"net/http/httptest"
	"reflect"

	"productGofr/models"
	"productGofr/services"
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
			_, err := testhndlr.GetByIdHandler(ctx)
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

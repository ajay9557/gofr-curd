package products

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	gofrError "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
	"github.com/ridhdhish-desai-zs/product-gofr/service"
)

func TestGetByIdHandler(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := service.NewMockProduct(ctrl)
	handler := New(mockProductService)

	product := models.Product{
		Id:       1,
		Name:     "mouse",
		Category: "electronics",
	}

	tests := []struct {
		desc          string
		id            string
		expectedError error
		mockCall      *gomock.Call
	}{
		{
			desc:          "Successfull operation case",
			id:            "1",
			expectedError: nil,
			mockCall:      mockProductService.EXPECT().GetById(gomock.Any(), "1").Return(&product, nil),
		},
		{
			desc:          "id must be a number",
			id:            "abc",
			expectedError: gofrError.EntityNotFound{Entity: "products", ID: "abc"},
			mockCall:      mockProductService.EXPECT().GetById(gomock.Any(), "abc").Return(nil, gofrError.EntityNotFound{Entity: "products", ID: "abc"}),
		},
		{
			desc:          "id must be greater than 0",
			id:            "-1",
			expectedError: gofrError.EntityNotFound{Entity: "products", ID: "-1"},
			mockCall:      mockProductService.EXPECT().GetById(gomock.Any(), "-1").Return(nil, gofrError.EntityNotFound{Entity: "products", ID: "-1"}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/products/{id}", nil)
			res := httptest.NewRecorder()

			r := request.NewHTTPRequest(req)
			w := responder.NewContextualResponder(res, req)

			ctx := gofr.NewContext(w, r, app)
			ctx.Context = context.Background()

			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})

			_, err := handler.GetByIdHandler(ctx)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetHandler(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductService := service.NewMockProduct(ctrl)
	handler := New(mockProductService)

	products := []*models.Product{
		{
			Id:       1,
			Name:     "mouse",
			Category: "electronics",
		},
	}

	tests := []struct {
		desc          string
		expectedError error
		mockCall      *gomock.Call
	}{
		{
			desc:          "Successfull operation case",
			expectedError: nil,
			mockCall:      mockProductService.EXPECT().Get(gomock.Any()).Return(products, nil),
		},
		{
			desc:          "Something went wrong",
			expectedError: gofrError.EntityNotFound{Entity: "products"},
			mockCall:      mockProductService.EXPECT().Get(gomock.Any()).Return(nil, gofrError.EntityNotFound{Entity: "products"}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			_, err := handler.GetHandler(ctx)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}
		})
	}
}

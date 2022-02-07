package products

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
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
		ID:       1,
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
			desc:          "Successful operation case",
			id:            "1",
			expectedError: nil,
			mockCall:      mockProductService.EXPECT().GetByID(gomock.Any(), "1").Return(&product, nil),
		},
		{
			desc:          "id must be a number",
			id:            "abc",
			expectedError: gofrError.EntityNotFound{Entity: "products", ID: "abc"},
			mockCall: mockProductService.EXPECT().GetByID(gomock.Any(), "abc").Return(
				nil,
				gofrError.EntityNotFound{Entity: "products", ID: "abc"},
			),
		},
		{
			desc:          "id must be greater than 0",
			id:            "-1",
			expectedError: gofrError.EntityNotFound{Entity: "products", ID: "-1"},
			mockCall: mockProductService.EXPECT().GetByID(gomock.Any(), "-1").Return(
				nil,
				gofrError.EntityNotFound{Entity: "products", ID: "-1"},
			),
		},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/products/{id}", nil)
			res := httptest.NewRecorder()

			r := request.NewHTTPRequest(req)
			w := responder.NewContextualResponder(res, req)

			ctx := gofr.NewContext(w, r, app)
			ctx.Context = context.Background()

			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})

			_, err := handler.GetByIDHandler(ctx)

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
			ID:       1,
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
			desc:          "Successful operation case",
			expectedError: nil,
			mockCall:      mockProductService.EXPECT().Get(gomock.Any()).Return(products, nil),
		},
		{
			desc:          "Something went wrong",
			expectedError: gofrError.EntityNotFound{Entity: "products"},
			mockCall: mockProductService.EXPECT().Get(gomock.Any()).Return(
				nil,
				gofrError.EntityNotFound{Entity: "products"},
			),
		},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()
			_, err := handler.GetHandler(ctx)

			fmt.Println(!errors.Is(err, tc.expectedError))

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestCreateHandler(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockProduct(ctrl)
	productHandler := New(mockService)

	p := models.Product{
		ID:       1,
		Name:     "mouse",
		Category: "electronics",
	}

	tests := []struct {
		desc          string
		expectedError error
		body          models.Product
		mockCall      *gomock.Call
	}{
		{
			desc:          "Success case",
			expectedError: nil,
			body:          p,
			mockCall:      mockService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&p, nil),
		},
		{
			desc:          "Empty body",
			expectedError: gofrError.MissingParam{Param: []string{"name", "category"}},
			body:          models.Product{},
			mockCall:      nil,
		},
		{
			desc:          "Error while creating",
			expectedError: errors.New("SOMETHING WENT WRONG"),
			body: models.Product{
				Name:     "mouse",
				Category: "electronics",
			},
			mockCall: mockService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("SOMETHING WENT WRONG")),
		},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			pr, _ := json.Marshal(tc.body)
			req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(pr))
			res := httptest.NewRecorder()

			r := request.NewHTTPRequest(req)
			w := responder.NewContextualResponder(res, req)

			ctx := gofr.NewContext(w, r, app)

			_, err := productHandler.CreateProductHandler(ctx)

			fmt.Println(!errors.Is(err, tc.expectedError))

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestUpdateHandler(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockProduct(ctrl)
	productHandler := New(mockService)

	p := models.Product{
		ID:       1,
		Name:     "mouse",
		Category: "electronics",
	}

	tests := []struct {
		desc          string
		expectedError error
		id            string
		body          models.Product
		mockCall      *gomock.Call
	}{
		{
			desc:          "Success case",
			expectedError: nil,
			id:            "1",
			body: models.Product{
				Category: "electronics",
			},
			mockCall: mockService.EXPECT().UpdateByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(&p, nil),
		},
		{
			desc:          "Empty body",
			id:            "1",
			expectedError: gofrError.MissingParam{Param: []string{"name", "category"}},
			body:          models.Product{},
			mockCall:      nil,
		},
		{
			desc:          "Id must be number",
			expectedError: gofrError.InvalidParam{Param: []string{"id"}},
			id:            "abc",
			body: models.Product{
				Category: "electronics",
			},
			mockCall: mockService.EXPECT().UpdateByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(
				nil,
				gofrError.InvalidParam{Param: []string{"id"}},
			),
		},
		{
			desc:          "Id must be greater than 0",
			expectedError: gofrError.InvalidParam{Param: []string{"id"}},
			id:            "-1",
			body: models.Product{
				Category: "electronics",
			},
			mockCall: mockService.EXPECT().UpdateByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(
				nil,
				gofrError.InvalidParam{Param: []string{"id"}},
			),
		},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			pr, _ := json.Marshal(tc.body)
			req := httptest.NewRequest(http.MethodPut, "/products/{id}", bytes.NewBuffer(pr))
			res := httptest.NewRecorder()

			r := request.NewHTTPRequest(req)
			w := responder.NewContextualResponder(res, req)

			ctx := gofr.NewContext(w, r, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})

			_, err := productHandler.UpdateProductHandler(ctx)

			fmt.Println(!errors.Is(err, tc.expectedError))

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestDeleteHandler(t *testing.T) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockProduct(ctrl)
	productHandler := New(mockService)

	tests := []struct {
		desc          string
		expectedError error
		id            string
		mockCall      *gomock.Call
	}{
		{
			desc:          "Success case",
			expectedError: nil,
			id:            "1",
			mockCall:      mockService.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(nil),
		},
		{
			desc:          "Id must be number",
			expectedError: gofrError.InvalidParam{Param: []string{"id"}},
			id:            "abc",
			mockCall:      mockService.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(gofrError.InvalidParam{Param: []string{"id"}}),
		},
		{
			desc:          "Id must be greater than 0",
			expectedError: gofrError.InvalidParam{Param: []string{"id"}},
			id:            "-1",
			mockCall:      mockService.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(gofrError.InvalidParam{Param: []string{"id"}}),
		},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/products/{id}", nil)
			res := httptest.NewRecorder()

			r := request.NewHTTPRequest(req)
			w := responder.NewContextualResponder(res, req)

			ctx := gofr.NewContext(w, r, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})

			_, err := productHandler.DeleteProductHandler(ctx)

			fmt.Println(!errors.Is(err, tc.expectedError))

			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("Expected: %v, Got: %v", tc.expectedError, err)
			}
		})
	}
}

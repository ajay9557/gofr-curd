package product

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/service"
)

func TestGetProductById(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockProductService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc     string
		id       string
		mockCall []*gomock.Call
		expOut   interface{}
		expErr   error
	}{
		{
			desc: "Case 1: Success Case",
			id:   "1",
			mockCall: []*gomock.Call{
				mockService.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(models.Product{
					Id:   1,
					Name: "name-1",
					Type: "type-1",
				}, nil),
			},
			expOut: models.Product{
				Id:   1,
				Name: "name-1",
				Type: "type-1",
			},
			expErr: nil,
		},
		{
			desc:   "Case 2: Failure Case1 invalid id",
			id:     "2q",
			expOut: nil,
			expErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Case 3: Failure Case2",
			id:   "99",
			mockCall: []*gomock.Call{
				mockService.EXPECT().GetProductById(gomock.Any(), gomock.Any()).
					Return(models.Product{}, errors.EntityNotFound{Entity: "product", ID: "99"}),
			},
			expOut: nil,
			expErr: errors.EntityNotFound{Entity: "product", ID: "99"},
		},
	}

	app := gofr.New()

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			httpResRec := httptest.NewRecorder()
			httpReq := httptest.NewRequest("GET", "/product", nil)

			req := request.NewHTTPRequest(httpReq)
			res := responder.NewContextualResponder(httpResRec, httpReq)

			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": test.id,
			})

			result, err := mockHandler.GetProductById(ctx)

			if !reflect.DeepEqual(test.expOut, result) {
				fmt.Printf("Expected : %v, Got : %v", test.expOut, result)
			}

			if !reflect.DeepEqual(test.expErr, err) {
				fmt.Printf("Expected : %v, Got : %v", test.expErr, err)
			}

		})
	}

}

func TestGetAllProducts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockProductService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc     string
		mockCall []*gomock.Call
		expOut   interface{}
		expErr   error
	}{
		{
			desc: "Case 1: Success Case",
			mockCall: []*gomock.Call{
				mockService.EXPECT().GetAllProducts(gomock.Any()).Return([]models.Product{
					{
						Id:   1,
						Name: "name-1",
						Type: "type-1",
					},
				}, nil)},
			expOut: []models.Product{
				{
					Id:   1,
					Name: "name-1",
					Type: "type-1",
				},
			},
			expErr: nil,
		},
		{
			desc:     "Case 2: Failure Case",
			mockCall: []*gomock.Call{mockService.EXPECT().GetAllProducts(gomock.Any()).Return(nil, errors.Error("internal error"))},
			expOut:   nil,
			expErr:   errors.Error("internal error"),
		},
	}
	app := gofr.New()

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/product", nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			resp, err := mockHandler.GetAllProducts(ctx)
			if !reflect.DeepEqual(test.expOut, resp) {
				t.Error("expected ", test.expOut, "obtained", resp)
			}
			if !reflect.DeepEqual(test.expErr, err) {
				t.Error("expected ", test.expErr, "obtained", err)
			}
		})
	}

}

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockService := service.NewMockProductService(ctrl)

	mockHandler := New(mockService)

	testCases := []struct {
		desc    string
		product []byte
		mock    []*gomock.Call
		ExpOut  interface{}
		ExpErr  error
	}{
		{
			desc: "Case 1: Success Case",
			product: []byte(`{
				"Id":   1,
				"Name": "name-1",
				"Type": "type-1"
			}`),
			mock: []*gomock.Call{mockService.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(nil)},
			ExpOut: models.Product{
				Id:   1,
				Name: "name-1",
				Type: "type-1",
			},
			ExpErr: nil,
		},
		{
			desc: "Case 1: Failure Case",
			product: []byte(`{
				"Id":   8,
				"Name": "name-8",
				"Type": "type-8"
			}`),
			mock:   []*gomock.Call{mockService.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).Return(errors.Error("internal error"))},
			ExpOut: nil,
			ExpErr: errors.Error("internal error"),
		},
	}

	app := gofr.New()

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/product", bytes.NewReader(ts.product))

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			resp, err := mockHandler.CreateProduct(ctx)
			if !reflect.DeepEqual(ts.ExpOut, resp) {
				t.Error("expected ", ts.ExpOut, "obtained", resp)
			}
			if !reflect.DeepEqual(ts.ExpErr, err) {
				t.Error("expected ", ts.ExpErr, "obtained", err)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := service.NewMockProductService(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct {
		desc   string
		Id     string
		mock   []*gomock.Call
		ExpErr error
	}{
		{
			desc:   "Case 1: Success case",
			Id:     "1",
			mock:   []*gomock.Call{mockStore.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).Return(nil)},
			ExpErr: nil,
		},
		{
			desc:   "Case 2: Failure case1",
			Id:     "",
			ExpErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:   "Failure case - invalid id",
			Id:     "1a",
			ExpErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Failure case - 3",
			Id:   "3",
			mock: []*gomock.Call{
				mockStore.EXPECT().DeleteProduct(gomock.Any(), gomock.Any()).
					Return(errors.Error("internal error")),
			},
			ExpErr: errors.Error("internal error"),
		},
	}

	app := gofr.New()

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/product", nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": ts.Id,
			})
			_, err := mockHandler.DeleteProduct(ctx)
			if !reflect.DeepEqual(ts.ExpErr, err) {
				t.Error("expected ", ts.ExpErr, "obtained", err)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := service.NewMockProductService(ctrl)

	mockHandler := New(mockStore)

	testCases := []struct {
		desc     string
		product  []byte
		mockCall []*gomock.Call
		ExpOut   interface{}
		ExpErr   error
	}{
		{
			desc: "Case 1: Success case",
			product: []byte(`{
				"Id":   1,
				"Name": "name-1",
				"Type": "type-1"
			}`),
			mockCall: []*gomock.Call{mockStore.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(nil)},
			ExpOut: models.Product{
				Id:   1,
				Name: "name-1",
				Type: "type-1",
			},
			ExpErr: nil,
		},
		{
			desc: "Case 2: Failure Case",
			product: []byte(`{
				"Id":   2,
				"Name": "name-2",
				"Type": "type-2",
			}`),
			ExpOut: nil,
			ExpErr: errors.Error("cannot update product data"),
		},
	}

	app := gofr.New()

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/product", bytes.NewReader(ts.product))

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			resp, err := mockHandler.UpdateProduct(ctx)
			if !reflect.DeepEqual(ts.ExpOut, resp) {
				t.Error("expected ", ts.ExpOut, "obtained", resp)
			}
			if !reflect.DeepEqual(ts.ExpErr, err) {
				t.Error("expected ", ts.ExpErr, "obtained", err)
			}
		})
	}
}


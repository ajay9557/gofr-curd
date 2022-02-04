package product

import (
	"bytes"
	"gofr-curd/models"
	"gofr-curd/service"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
)

func TestGetById(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc        string
		id          string
		mock        []*gomock.Call
		expectedRes interface{}
		expectedErr error
	}{
		{
			desc: "Success case",
			id:   "1",
			mock: []*gomock.Call{
				mockService.EXPECT().GetByProductId(gomock.Any(), gomock.Any()).Return(models.Product{
					Id:   1,
					Name: "jeans",
					Type: "clothes",
				}, nil),
			},
			expectedRes: models.Product{
				Id:   1,
				Name: "jeans",
				Type: "clothes",
			},
			expectedErr: nil,
		},
		{
			desc:        "Failure case - empty id",
			id:          "",
			expectedRes: nil,
			expectedErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:        "Failure case - invalid id",
			id:          "1a",
			expectedRes: nil,
			expectedErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Failure case - 3",
			id:   "3",
			mock: []*gomock.Call{
				mockService.EXPECT().GetByProductId(gomock.Any(), gomock.Any()).
					Return(models.Product{}, errors.EntityNotFound{Entity: "product", ID: "3"}),
			},
			expectedRes: nil,
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "3"},
		},
	}

	app := gofr.New()

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/product", nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": ts.id,
			})
			resp, err := mockHandler.GetById(ctx)
			if !reflect.DeepEqual(ts.expectedRes, resp) {
				t.Error("expected ", ts.expectedRes, "obtained", resp)
			}
			if !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestGetAllProductDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc        string
		mock        []*gomock.Call
		expectedRes interface{}
		expectedErr error
	}{
		{
			desc: "Success case",
			mock: []*gomock.Call{
				mockService.EXPECT().GetProducts(gomock.Any()).Return([]models.Product{
					{Id: 1,
						Name: "jeans",
						Type: "clothes",
					},
				}, nil),
			},
			expectedRes: []models.Product{
				{
					Id:   1,
					Name: "jeans",
					Type: "clothes",
				},
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case",
			mock: []*gomock.Call{
				mockService.EXPECT().GetProducts(gomock.Any()).Return(nil, errors.Error("internal error"))},
			expectedRes: nil,
			expectedErr: errors.Error("internal error"),
		},
	}
	app := gofr.New()

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/product", nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			resp, err := mockHandler.GetAllProductDetails(ctx)
			if !reflect.DeepEqual(ts.expectedRes, resp) {
				t.Error("expected ", ts.expectedRes, "obtained", resp)
			}
			if !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestInsertProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc        string
		product     []byte
		mock        []*gomock.Call
		expectedRes interface{}
		expectedErr error
	}{
		{
			desc: "Success case",
			product: []byte(`{
				"Id":   1,
				"Name": "jeans",
				"Type": "clothes"
			}`),
			mock: []*gomock.Call{
				mockService.EXPECT().InsertProductDetails(gomock.Any(), gomock.Any()).Return(nil),
			},
			expectedRes: models.Product{
				Id:   1,
				Name: "jeans",
				Type: "clothes",
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case - invalid body",
			product: []byte(`{
				"Id":   2,
				"Name": "jeans",
				"Type": "clothes",
			}`),
			expectedRes: nil,
			expectedErr: errors.InvalidParam{Param: []string{"body"}},
		},
		{
			desc: "Failure case 2",
			product: []byte(`{
				"Id":   3,
				"Name": "jeans",
				"Type": "clothes"
			}`),
			mock: []*gomock.Call{
				mockService.EXPECT().InsertProductDetails(gomock.Any(), gomock.Any()).Return(
					errors.Error("internal errror"),
				),
			},
			expectedRes: nil,
			expectedErr: errors.Error("internal errror"),
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
			resp, err := mockHandler.InsertProduct(ctx)
			if !reflect.DeepEqual(ts.expectedRes, resp) {
				t.Error("expected ", ts.expectedRes, "obtained", resp)
			}
			if !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestUpdateProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc        string
		product     []byte
		mock        []*gomock.Call
		expectedRes interface{}
		expectedErr error
	}{
		{
			desc: "Success case",
			product: []byte(`{
				"Id":   1,
				"Name": "jeans",
				"Type": "clothes"
			}`),
			mock: []*gomock.Call{
				mockService.EXPECT().UpdateProductDetails(gomock.Any(), gomock.Any()).Return(nil),
			},
			expectedRes: models.Product{
				Id:   1,
				Name: "jeans",
				Type: "clothes",
			},
			expectedErr: nil,
		},
		{
			desc: "Failure case - invalid body",
			product: []byte(`{
				"Id":   2,
				"Name": "jeans",
				"Type": "clothes",
			}`),
			expectedRes: nil,
			expectedErr: errors.InvalidParam{Param: []string{"body"}},
		},
		{
			desc: "Failure case 2",
			product: []byte(`{
				"Id":   3,
				"Name": "jeans",
				"Type": "clothes"
			}`),
			mock: []*gomock.Call{
				mockService.EXPECT().UpdateProductDetails(gomock.Any(), gomock.Any()).Return(
					errors.Error("internal errror"),
				),
			},
			expectedRes: nil,
			expectedErr: errors.Error("internal errror"),
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
			resp, err := mockHandler.UpdateProductById(ctx)
			if !reflect.DeepEqual(ts.expectedRes, resp) {
				t.Error("expected ", ts.expectedRes, "obtained", resp)
			}
			if !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

func TestDeleteByProductId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc        string
		id          string
		mock        []*gomock.Call
		expectedErr error
	}{
		{
			desc: "Success case",
			id:   "1",
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteProductById(gomock.Any(), gomock.Any()).Return(nil),
			},
			expectedErr: nil,
		},
		{
			desc:        "Failure case - empty id",
			id:          "",
			expectedErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:        "Failure case - invalid id",
			id:          "1a",
			expectedErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Failure case - 3",
			id:   "3",
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteProductById(gomock.Any(), gomock.Any()).
					Return(errors.Error("internal error")),
			},
			expectedErr: errors.Error("internal error"),
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
				"id": ts.id,
			})
			_, err := mockHandler.DeleteByProductId(ctx)
			if !reflect.DeepEqual(ts.expectedErr, err) {
				t.Error("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

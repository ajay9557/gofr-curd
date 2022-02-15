package product

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"

	//	"strconv"
	"testing"
	models "zopsmart/productgofr/models"
	service "zopsmart/productgofr/services"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
)





func TestGetByID(t *testing.T) {

	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := service.NewMockServices(ctrl)
	h := New(mock)
	
	testCases := []struct {
		desc   string
		input  string
		calls  *gomock.Call
		resp   models.Response
		expErr error
	}{
		{
			desc:  "success case",
			input: "1",
			resp: models.Response{
				Data: models.Product{
					Id:   1,
					Name: "shirt",
					Type: "fashion",
				},
				Message:    "data retrieved",
				StatusCode: http.StatusOK,
			},
			calls: 
				mock.EXPECT().GetProdByID(gomock.Any(), gomock.Any()).Return(models.Product{
					Id:   1,
					Name: "shirt",
					Type: "fashion",
				}, nil),
			expErr: nil,
		},
		{
			desc:  "id not in database",
			input: "102",
			resp:  models.Response{},
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "102",
			},
			calls: mock.EXPECT().GetProdByID(gomock.Any(), gomock.Any()).Return(models.Product{}, errors.EntityNotFound{
				Entity: "product",
				ID:     "102",
			}),
		},
		{
			desc:  "invalid id",
			input: "-1",
			resp:  models.Response{},
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
			calls: mock.EXPECT().GetProdByID(gomock.Any(), -1).Return(models.Product{}, errors.InvalidParam{
				Param: []string{"id"}}),
		},
		{
			desc:  "invalid id",
			input: "abc",
			resp:  models.Response{},
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
		},
		{
			desc:  "missing params",
			input: "",
			resp: models.Response{},
			expErr: errors.MissingParam{Param: []string{"id"},
			},
		},
	 }

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/product", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.input,
			})

			resp, err := h.GetProdByIdHandler(ctx)
			if !reflect.DeepEqual(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}

			if tc.expErr == nil && !reflect.DeepEqual(resp, tc.resp) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.resp, resp)
			}
		})
	}
}

func TestGet(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := service.NewMockServices(ctrl)
	h := New(mock)
	testCases := []struct {
		desc     string
		mockCall *gomock.Call
		expResp  []models.Product
		expErr   error
	}{
		{
			desc: "Success case",
			mockCall: 
				mock.EXPECT().GetAllProd(gomock.Any()).
					Return([]models.Product{
						{
							Id:   1,
							Name: "shirt",
							Type: "fashion",
						},
						{
							Id:   2,
							Name: "mobile",
							Type: "electronics",
						},
					}, nil),
			expResp: []models.Product{
				{
					Id:   1,
					Name: "shirt",
					Type: "fashion",
				},
				{
					Id:   2,
					Name: "mobile",
					Type: "electronics",
				},
			},
			expErr: nil,
		},
		
		{
			desc:    "error getting products",
			expResp: nil,
			expErr:  errors.EntityNotFound{Entity: "products", ID: "all"},
			mockCall: 
				mock.EXPECT().GetAllProd(gomock.Any()).Return(nil, errors.EntityNotFound{Entity: "products", ID: "all"}),
		},
	}

	for _, test := range testCases {
		tc := test
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/product", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		resp, err := h.GetAllProductHandler(ctx)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expResp) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expResp, resp)
		}
	}
}

func TestCreate(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := service.NewMockServices(ctrl)
	h := New(mock)
	testCases := []struct {
		desc     string
		input    []byte
		expErr   error
		mockCall *gomock.Call
	}{
		{
			desc:   "success",
			input:  []byte(`{"id":3,"name": "laptop","type": "electronics"}`),
			expErr: nil,
			mockCall:
				mock.EXPECT().CreateProduct(gomock.Any(), models.Product{
					Id:   3,
					Name: "laptop",
					Type: "electronics",
				}).Return(nil),
		},
		{
			desc:     "error binding",
			input:    []byte(`{mock error invalid body}`),
			mockCall: nil,
			expErr:   errors.InvalidParam{Param: []string{"body"}},
		},
		{
			desc:  "error from service",
			input: []byte(`{"id":3,"name": "test","type": "example"}`),
			mockCall: 
				mock.EXPECT().CreateProduct(gomock.Any(), gomock.Any()).
					Return(errors.EntityAlreadyExists{}),
			expErr: errors.EntityAlreadyExists{},
		},
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/product", bytes.NewReader(tc.input))
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)

		_, err := h.CreateProductHandler(ctx)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}

func TestUpdate(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := service.NewMockServices(ctrl)
	h := New(mock)


	testCases := []struct {
		desc     string
		id       string
		input    []byte
		expResp  models.Product
		expErr   error
		mockCall *gomock.Call
	}{
		{
			desc:  "Success case",
			id:    "1",
			input: []byte(`{"name":"updatedname","type":"updatedtype"}`),
			expResp: models.Product{
				Id:   1,
				Name: "updatedname",
				Type: "updatedtype",
			},
			expErr: nil,
			mockCall: 
				mock.EXPECT().UpdateProduct(gomock.Any(), models.Product{
					Id:   1,
					Name: "updatedname",
					Type: "updatedtype",
				}).Return(nil),
		},
		{
			desc:   "missing params",
			id:     "",
			expErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:   "invalid params",
			id:     "asd",
			expErr: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc:   "binding err",
			id:     "1",
			input:  []byte(`mock error binding`),
			expErr: errors.InvalidParam{Param: []string{"body"}},
		},
		{
			desc:   "error in update",
			id:     "1",
			input:  []byte(`{"name":"updatedname","type":"updatedtype"}`),
			expErr: errors.Error("error updating record"),
			mockCall: 
				mock.EXPECT().UpdateProduct(gomock.Any(), models.Product{
					Id:   1,
					Name: "updatedname",
					Type: "updatedtype",
				}).Return(errors.Error("error updating record")),
		},
	}

	for _, tc := range testCases {
	
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/product", bytes.NewReader(tc.input))
	req := request.NewHTTPRequest(r)
	res := responder.NewContextualResponder(w, r)
	ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})

		_, err := h.UpdateProductHandler(ctx)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}

func TestDelete(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := service.NewMockServices(ctrl)
	h := New(mock)
	testCases := []struct {
		desc     string
		id       string
		expErr   error
		expResp  *models.Response
		mockCall *gomock.Call
	}{
		{
			desc: "success",
			id:   "1",
			expResp: &models.Response{
				Message:    "deleted successfully",
				StatusCode: http.StatusOK,
			},
			mockCall: 
				mock.EXPECT().DeleteProduct(gomock.Any(), 1).
					Return(nil),
		},
		{
			desc:   "missing params",
			id:     "",
			expErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:   "invalid params",
			id:     "abc",
			expErr: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc:   "error deleting records",
			id:     "1",
			expErr: errors.Error("error deleting record"),
			mockCall: 
				mock.EXPECT().DeleteProduct(gomock.Any(), 1).Return(errors.Error("error deleting record")),
		},
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/product", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})

		resp, err := h.DeleteProductHandler(ctx)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expResp) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expResp, resp)
		}
	}
}

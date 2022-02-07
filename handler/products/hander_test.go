package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	models "zopsmart/gofr-curd/model"
	services "zopsmart/gofr-curd/service"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	gomock "github.com/golang/mock/gomock"
)

func Test_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockProductservice(ctrl)
	h := New(serv)
	app := gofr.New()

	tcs := []struct {
		desc           string
		ID             string
		err            error
		expectedOutput interface{}
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   "1",
			err:  nil,
			expectedOutput: models.Product{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			mock: []*gomock.Call{serv.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(models.Product{
				ID: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc: "Failure-2",
			ID:   "45",
			err: errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			},
			expectedOutput: nil,
			mock: []*gomock.Call{serv.EXPECT().GetByID(gomock.Any(), gomock.Any()).
				Return(models.Product{}, errors.EntityNotFound{
					Entity: "Product",
					ID:     "45",
				})},
		},
	}

	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.ID,
			})
			resp, err := h.GetByID(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(resp, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, resp)
			}
		})
	}
}

func Test_UpdateByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockProductservice(ctrl)
	h := New(serv)
	app := gofr.New()

	tcs := []struct {
		desc           string
		ID             string
		input          []byte
		err            error
		expectedOutput interface{}
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   "1",
			err:  nil,
			input: []byte(`{
				"Id":0,
				"Name": "Shirts",
				"Type": "US POLO"
			}`),
			expectedOutput: models.Product{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			mock: []*gomock.Call{serv.EXPECT().UpdateByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Product{
				ID: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc: "Failure-2",
			ID:   "45",
			input: []byte(`{
				"Id":0,
				"Name": "Shirts",
				"Type": "US POLO"
			}`),
			err: errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			},
			expectedOutput: nil,
			mock: []*gomock.Call{serv.EXPECT().UpdateByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Product{}, errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			})},
		},
		{
			desc: "Success",
			ID:   "1",
			err:  errors.InvalidParam{Param: []string{"body"}},
			input: []byte(`
				"Id":0,
				"Name": "Shirts",
				"Type": "US POLO"
			`),
			expectedOutput: nil,
			mock:           nil,
		},
	}

	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", bytes.NewReader(tc.input))
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.ID,
			})
			resp, err := h.UpdateByID(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(resp, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, resp)
			}
		})
	}
}

func Test_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockProductservice(ctrl)
	h := New(serv)
	app := gofr.New()

	tcs := []struct {
		desc           string
		err            error
		expectedOutput interface{}
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			err:  nil,
			expectedOutput: []models.Product{{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			}},
			mock: []*gomock.Call{serv.EXPECT().GetProducts(gomock.Any()).Return([]models.Product{{
				ID: 1, Name: "Shirts", Type: "US POLO"}}, nil)},
		}, {
			desc:           "Failure",
			err:            errors.InvalidParam{Param: []string{"body"}},
			expectedOutput: nil,
			mock: []*gomock.Call{serv.EXPECT().GetProducts(gomock.Any()).
				Return([]models.Product{}, errors.InvalidParam{Param: []string{"body"}})},
		},
	}
	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			resp, err := h.GetProducts(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(resp, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, resp)
			}
		})
	}
}

func Test_AddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockProductservice(ctrl)
	h := New(serv)
	app := gofr.New()
	tcs := []struct {
		desc           string
		err            error
		input          []byte
		expectedOutput interface{}
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			err:  nil,
			input: []byte(`{
				"Id":0,
				"Name": "Shirts",
				"Type": "US POLO"
			}`),
			expectedOutput: models.Product{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			mock: []*gomock.Call{serv.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(models.Product{
				ID: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc: "Failure",
			err:  errors.Error("internal error"),
			input: []byte(`{
				"Id":0,
				"Name": "Shirts",
				"Type": "US POLO"
			}`),
			expectedOutput: nil,
			mock: []*gomock.Call{serv.EXPECT().AddProduct(gomock.Any(), gomock.Any()).
				Return(models.Product{}, errors.Error("internal error"))},
		},
		{
			desc: "Failure",
			err:  errors.InvalidParam{Param: []string{"body"}},
			input: []byte(`
				"Id":0,
				"Name": "Shirts",
				"Type": "US POLO"
			`),
			expectedOutput: nil,
			mock:           nil,
		},
	}

	for _, test := range tcs {
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", bytes.NewReader(tc.input))
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			resp, err := h.AddProduct(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(resp, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, resp)
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockProductservice(ctrl)
	h := New(serv)
	app := gofr.New()
	tcs := []struct {
		desc string
		err  error
		ID   string
		mock []*gomock.Call
	}{
		{
			desc: "Success",
			err:  nil,
			ID:   "1",
			mock: []*gomock.Call{serv.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure-2",
			ID:   "45",
			err: errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			},
			mock: []*gomock.Call{serv.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			})},
		},
	}

	for _, test := range tcs {
		tc := test

		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.ID,
			})
			_, err := h.DeleteByID(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

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

func Test_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockProductservice(ctrl)
	h := New(serv)
	app := gofr.New()

	tcs := []struct {
		desc           string
		Id             string
		err            error
		expectedOutput interface{}
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   "1",
			err:  nil,
			expectedOutput: models.Product{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			mock: []*gomock.Call{serv.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(models.Product{
				Id: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc: "Failure-2",
			Id:   "45",
			err: errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			},
			expectedOutput: nil,
			mock: []*gomock.Call{serv.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(models.Product{}, errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			})},
		},
	}

	for _, tc := range tcs {

		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.Id,
			})
			resp, err := h.GetById(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(resp, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, resp)
			}
		})
	}
}

func Test_UpdateById(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockProductservice(ctrl)
	h := New(serv)
	app := gofr.New()

	tcs := []struct {
		desc           string
		Id             string
		input []byte
		err            error
		expectedOutput interface{}
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   "1",
			err:  nil,
			input: []byte(`{
				"Id":0,
				"Name": "Shirts",
				"Type": "US POLO"
			}`),
			expectedOutput: models.Product{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			mock: []*gomock.Call{serv.EXPECT().UpdateById(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Product{
				Id: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc: "Failure-2",
			Id:   "45",
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
			mock: []*gomock.Call{serv.EXPECT().UpdateById(gomock.Any(), gomock.Any(),gomock.Any()).Return(models.Product{}, errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			})},
		},
		{
			desc: "Success",
			Id:   "1",
			err:  errors.InvalidParam{Param: []string{"body"}},
			input: []byte(`
				"Id":0,
				"Name": "Shirts",
				"Type": "US POLO"
			`),
			expectedOutput: nil,
			mock: nil,
		},

	}

	for _, tc := range tcs {

		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", bytes.NewReader(tc.input))
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.Id,
			})
			resp, err := h.UpdateById(ctx)
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
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			}},
			mock: []*gomock.Call{serv.EXPECT().GetProducts(gomock.Any()).Return([]models.Product{{
				Id: 1, Name: "Shirts", Type: "US POLO"}}, nil)},
		}, {
			desc:           "Failure",
			err:            errors.InvalidParam{Param: []string{"body"}},
			expectedOutput: nil,
			mock:           []*gomock.Call{serv.EXPECT().GetProducts(gomock.Any()).Return([]models.Product{}, errors.InvalidParam{Param: []string{"body"}})},
		},
	}
	for _, tc := range tcs {

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
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			mock: []*gomock.Call{serv.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(models.Product{
				Id: 1, Name: "Shirts", Type: "US POLO"}, nil)},
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
			mock:           []*gomock.Call{serv.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(models.Product{}, errors.Error("internal error"))},
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
	for _, tc := range tcs {

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
		id   string
		mock []*gomock.Call
	}{
		{
			desc: "Success",
			err:  nil,
			id:   "1",
			mock: []*gomock.Call{serv.EXPECT().DeleteById(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure-2",
			id:   "45",
			err: errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			},
			mock: []*gomock.Call{serv.EXPECT().DeleteById(gomock.Any(), gomock.Any()).Return(errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			})},
		},
	}
	for _, tc := range tcs {

		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})
			_, err := h.DeleteById(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}

}

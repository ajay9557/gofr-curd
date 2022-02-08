package products

import (
	"bytes"
	"gofr-curd/models"
	"gofr-curd/services"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
)

func Test_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockService(ctrl)
	h := Handler{serv}
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
			mock: []*gomock.Call{serv.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(models.Product{
				ID: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc:           "Failure",
			ID:             "",
			err:            errors.MissingParam{Param: []string{"id"}},
			expectedOutput: nil,
		},
		{
			desc:           "Failure-1",
			ID:             "s",
			err:            errors.InvalidParam{Param: []string{"id"}},
			expectedOutput: nil,
		},
		{
			desc: "Failure-2",
			ID:   "45",
			err: errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			},
			expectedOutput: nil,
			mock: []*gomock.Call{serv.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(models.Product{}, errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			})},
		},
	}

	for _, test := range tcs {
		tc := test
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)

		t.Run(tc.desc, func(t *testing.T) {
			ctx.SetPathParams(map[string]string{
				"id": tc.ID,
			})
			_, err := h.GetByID(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func Test_DeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockService(ctrl)
	h := Handler{serv}
	app := gofr.New()

	tcs := []struct {
		desc string
		ID   string
		err  error
		mock []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   "1",
			err:  nil,
			mock: []*gomock.Call{serv.EXPECT().DeleteByProductID(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure",
			ID:   "absd123",
			err:  errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Failure -1",
			ID:   "",
			err:  errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Failure -2",
			ID:   "123",
			err:  errors.EntityNotFound{Entity: "Product", ID: "123"},
			mock: []*gomock.Call{serv.EXPECT().DeleteByProductID(gomock.Any(), gomock.Any()).Return(
				errors.EntityNotFound{Entity: "Product", ID: "123"})},
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

func Test_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockService(ctrl)
	h := Handler{serv}
	app := gofr.New()
	tcs := []struct {
		desc    string
		ID      string
		err     error
		product []byte
		mock    []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   "1",
			err:  nil,
			product: []byte(`
			{
				"ID":5,
				"Name":"jeans",
				"Type":"clothes"
			}
			`),
			mock: []*gomock.Call{serv.EXPECT().UpdateByProductID(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure",
			ID:   "1",
			err:  errors.InvalidParam{Param: []string{"body"}},
			product: []byte(`
			
				"ID":5,
				"Name":"jeans",
				"Type":"clothes"
			`),
		},
		{
			desc: "Failure-1",
			ID:   "1",
			err:  errors.Error("Internal DB error"),
			product: []byte(`
			{
				"ID":1,
				"Name":"jeans",
				"Type":"clothes"
			}
			`),
			mock: []*gomock.Call{serv.EXPECT().UpdateByProductID(gomock.Any(), gomock.Any()).Return(errors.Error("Internal DB error"))},
		},
	}

	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", bytes.NewReader(tc.product))
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"ID": tc.ID,
			})
			_, err := h.Update(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func Test_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockService(ctrl)
	h := Handler{serv}
	app := gofr.New()

	tcs := []struct {
		desc string
		ID   string
		body []byte
		err  error
		mock []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   "1",
			body: []byte(`
			{
			"ID":1,
			"name":"jeans",
			"type":"clothes"
			}`),
			err: nil,
			mock: []*gomock.Call{serv.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).Return(models.Product{
				ID:   1,
				Name: "jeans",
				Type: "clothes",
			}, nil)},
		},
		{
			desc: "Failure",
			ID:   "0",
			body: []byte(`
			{
			"ID":0,
			"name":"jeans",
			"type":"clothes"
			}`),
			err: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-1",
			ID:   "123",
			body: []byte(`
			"ID":123,
			"name":"jeans",
			"type":"clothes"`),
			err: errors.InvalidParam{Param: []string{"body"}},
		},
		{
			desc: "Failure-2",
			ID:   "1",
			body: []byte(`
			{
			"ID":1,
			"name":"jeans",
			"type":"clothes"
			}`),
			err: errors.Error("Internal DB error"),
			mock: []*gomock.Call{serv.EXPECT().InsertProduct(gomock.Any(), gomock.Any()).Return(
				models.Product{}, errors.Error("Internal DB error"))},
		},
	}
	for _, test := range tcs {
		tc := test
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", bytes.NewReader(tc.body))
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			_, err := h.Insert(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func Test_GetAllProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockService(ctrl)
	h := Handler{serv}
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
		},
		{
			desc:           "Failure",
			err:            errors.Error("Internal DB error"),
			expectedOutput: nil,
			mock:           []*gomock.Call{serv.EXPECT().GetProducts(gomock.Any()).Return([]models.Product{}, errors.Error("Internal DB error"))},
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
			_, err := h.GetAllProducts(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

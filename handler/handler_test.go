package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"github.com/himanshu-kumar-zs/gofr-curd/services"
)

func TestHandler_GetByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockSrv := services.NewMockProduct(ctrl)
	h := New(mockSrv)

	testcases := []struct {
		desc   string
		inp    string
		exp    *models.Product
		expErr error
		mock   []*gomock.Call
	}{
		{
			"id exists",
			"1",
			&models.Product{
				ID:   1,
				Name: "abc",
				Type: "xyz",
			},
			nil,
			[]*gomock.Call{
				mockSrv.EXPECT().GetByID(gomock.Any(), 1).
					Return(&models.Product{
						ID:   1,
						Name: "abc",
						Type: "xyz",
					}, nil),
			},
		},
		{
			desc:   "missing param",
			inp:    "",
			exp:    nil,
			expErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:   "invalid param",
			inp:    "1a",
			exp:    nil,
			expErr: errors.InvalidParam{Param: []string{"id"}},
		},
		// no need to check id exists or not in Handler it will be checked in service layer
		{
			desc:   "id does not exists",
			inp:    "1002",
			exp:    nil,
			expErr: errors.EntityNotFound{Entity: "product", ID: "1002"},
			mock: []*gomock.Call{
				mockSrv.EXPECT().GetByID(gomock.Any(), 1002).
					Return(nil, errors.EntityNotFound{Entity: "product", ID: "1002"}),
			},
		},
	}

	for _, tcs := range testcases {
		// make request and response
		r := httptest.NewRequest("GET", fmt.Sprintf("/product/%s", tcs.inp), nil)
		w := httptest.NewRecorder()

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		ctx.SetPathParams(map[string]string{
			"id": tcs.inp,
		})

		out, err := h.GetByID(ctx)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected err %v, got %v", tcs.desc, tcs.expErr, err)
		}

		if err == nil && !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

var prod = &models.Product{
	ID:   1,
	Name: "EliteBook",
	Type: "Laptop",
}

func TestHandler_Update(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockSrv := services.NewMockProduct(ctrl)
	h := New(mockSrv)

	testcases := []struct {
		desc   string
		inp    string
		body   *models.Product
		exp    *models.Product
		expErr error
		mock   []*gomock.Call
	}{
		{
			"id exists",
			"1",
			prod,
			prod,
			nil,
			[]*gomock.Call{
				mockSrv.EXPECT().Update(gomock.Any(), prod).
					Return(prod, nil),
			},
		},
		{
			desc:   "missing param",
			inp:    "",
			exp:    nil,
			expErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:   "invalid param",
			inp:    "1a",
			exp:    nil,
			expErr: errors.InvalidParam{Param: []string{"id"}},
		},
	}

	for _, tcs := range testcases {
		// make request and response
		pr, _ := json.Marshal(tcs.body)
		r := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/product/%s", tcs.inp), bytes.NewBuffer(pr))
		w := httptest.NewRecorder()

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		ctx.SetPathParams(map[string]string{
			"id": tcs.inp,
		})

		out, err := h.Update(ctx)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected err %v, got %v", tcs.desc, tcs.expErr, err)
		}

		if err == nil && !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestHandler_Create(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockSrv := services.NewMockProduct(ctrl)
	h := New(mockSrv)

	testcases := []struct {
		desc   string
		body   *models.Product
		exp    *models.Product
		expErr error
		mock   []*gomock.Call
	}{
		{
			"success case ",
			prod,
			prod,
			nil,
			[]*gomock.Call{
				mockSrv.EXPECT().Create(gomock.Any(), prod).
					Return(prod, nil),
			},
		},
	}

	for _, tcs := range testcases {
		// make request and response
		pr, _ := json.Marshal(tcs.body)
		r := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(pr))
		w := httptest.NewRecorder()

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		out, err := h.Create(ctx)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected err %v, got %v", tcs.desc, tcs.expErr, err)
		}

		if err == nil && !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestHandler_Delete(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockSrv := services.NewMockProduct(ctrl)
	h := New(mockSrv)

	testcases := []struct {
		desc   string
		inp    string
		expErr error
		mock   []*gomock.Call
	}{
		{
			"id exists",
			"1",
			nil,
			[]*gomock.Call{
				mockSrv.EXPECT().Delete(gomock.Any(), 1).
					Return(nil),
			},
		},
		{
			desc:   "missing param",
			inp:    "",
			expErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:   "invalid param",
			inp:    "1a",
			expErr: errors.InvalidParam{Param: []string{"id"}},
		},
	}

	for _, tcs := range testcases {
		// make request and response
		r := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/product/%s", tcs.inp), nil)
		w := httptest.NewRecorder()

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		ctx.SetPathParams(map[string]string{
			"id": tcs.inp,
		})

		resp, err := h.Delete(ctx)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected err %v, got %v", tcs.desc, tcs.expErr, err)
		}

		if err == nil && !reflect.DeepEqual(resp, "successfully deleted") {
			t.Errorf("%v, expected %v, got %v", tcs.desc, "successfully deleted", resp.(string))
		}
	}
}

func TestHandler_GetAll(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockSrv := services.NewMockProduct(ctrl)
	h := New(mockSrv)

	testcases := []struct {
		desc   string
		exp    []*models.Product
		expErr error
		mock   []*gomock.Call
	}{
		{
			"some product exists",
			[]*models.Product{prod},
			nil,
			[]*gomock.Call{
				mockSrv.EXPECT().GetAll(gomock.Any()).
					Return([]*models.Product{prod}, nil),
			},
		},
		{
			"empty table",
			[]*models.Product{},
			errors.EntityNotFound{
				Entity: "product",
			},
			[]*gomock.Call{
				mockSrv.EXPECT().GetAll(gomock.Any()).
					Return([]*models.Product{}, errors.EntityNotFound{
						Entity: "product",
					}),
			},
		},
	}

	for _, tcs := range testcases {
		// make request and response
		r := httptest.NewRequest(http.MethodGet, "/product", nil)
		w := httptest.NewRecorder()

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		out, err := h.GetAll(ctx)

		if err == nil && !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected err %v, got %v", tcs.desc, tcs.exp, out)
		}

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected err %v, got %v", tcs.desc, tcs.expErr, err)
		}
	}
}

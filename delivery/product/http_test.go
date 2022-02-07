package product

import (
	"bytes"
	"gofr-curd/models"
	"gofr-curd/service"
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

func TestGetById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := service.NewMockServices(ctrl)
	h := New(mock)

	testCases := []struct {
		desc   string
		input  string
		calls  []*gomock.Call
		resp   *models.Response
		expErr error
	}{
		{
			desc:  "success case",
			input: "1",
			resp: &models.Response{
				Data: models.Product{
					Id:   1,
					Name: "test",
					Type: "example",
				},
			},
			calls: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "test",
					Type: "example",
				}, nil),
			},
			expErr: nil,
		},
		{
			desc:  "id not in database",
			input: "102",
			resp:  nil,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "102",
			},
			calls: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 102).Return(nil, errors.EntityNotFound{
					Entity: "product",
					ID:     "102",
				}),
			},
		},
		{
			desc:  "invalid id",
			input: "-1",
			resp:  nil,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
			calls: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), -1).Return(nil, errors.InvalidParam{
					Param: []string{"id"},
				}),
			},
		},
		{
			desc:  "invalid id",
			input: "abc",
			resp:  nil,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
		},
		{
			desc:  "missing params",
			input: "",
			resp:  nil,
			expErr: errors.MissingParam{
				Param: []string{"id"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			ctx.SetPathParams(map[string]string{
				"id": tc.input,
			})

			resp, err := h.GetById(ctx)
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
		mockCall []*gomock.Call
		expResp  []*models.Product
		expErr   error
	}{
		{
			desc: "success",
			mockCall: []*gomock.Call{
				mock.EXPECT().Get(gomock.Any()).
					Return([]*models.Product{
						&models.Product{
							Id:   1,
							Name: "test",
							Type: "example",
						},
						&models.Product{
							Id:   2,
							Name: "this",
							Type: "that",
						},
					}, nil),
			},
			expResp: []*models.Product{
				&models.Product{
					Id:   1,
					Name: "test",
					Type: "example",
				},
				&models.Product{
					Id:   2,
					Name: "this",
					Type: "that",
				},
			},
			expErr: nil,
		},
		{
			desc:    "error getting products",
			expResp: nil,
			expErr:  errors.EntityNotFound{Entity: "products", ID: "all"},
			mockCall: []*gomock.Call{
				mock.EXPECT().Get(gomock.Any()).Return(nil, errors.EntityNotFound{Entity: "products", ID: "all"}),
			},
		},
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)
		resp, err := h.Get(ctx)
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
		mockCall []*gomock.Call
	}{
		{
			desc:   "success",
			input:  []byte(`{"name": "test","type": "example"}`),
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().Create(gomock.Any(), models.Product{
					Name: "test",
					Type: "example",
				}).Return(&models.Product{
					Id:   3,
					Name: "test",
					Type: "example",
				}, nil),
			},
		},
		{
			desc:     "error binding",
			input:    []byte(`{mock error invalid body}`),
			mockCall: nil,
			expErr:   errors.InvalidParam{Param: []string{"body"}},
		},
		{
			desc:     "error invalid id",
			input:    []byte(`{"id":1}`),
			mockCall: nil,
			expErr:   errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc:  "error from service",
			input: []byte(`{"name": "test","type": "example"}`),
			mockCall: []*gomock.Call{
				mock.EXPECT().Create(gomock.Any(), gomock.Any()).
					Return(nil, errors.EntityAlreadyExists{}),
			},
			expErr: errors.EntityAlreadyExists{},
		},
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://dummy", bytes.NewReader(tc.input))

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)
		_, err := h.Create(ctx)
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
		expResp  *models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc:  "success",
			id:    "1",
			input: []byte(`{"name":"hello","type":"world"}`),
			expResp: &models.Product{
				Id:   1,
				Name: "hello",
				Type: "world",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().Update(gomock.Any(), models.Product{
					Id:   1,
					Name: "hello",
					Type: "world",
				}).Return(&models.Product{
					Id:   1,
					Name: "hello",
					Type: "world",
				}, nil),
			},
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
			input:  []byte(`{"name":"hello","type":"world"}`),
			expErr: errors.Error("error updating record"),
			mockCall: []*gomock.Call{
				mock.EXPECT().Update(gomock.Any(), models.Product{
					Id:   1,
					Name: "hello",
					Type: "world",
				}).Return(nil, errors.Error("error updating record")),
			},
		},
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://dummy", bytes.NewReader(tc.input))

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})
		_, err := h.Update(ctx)
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
		mockCall []*gomock.Call
	}{
		{
			desc: "success",
			id:   "1",
			expResp: &models.Response{
				Message:    "deleted successfully",
				StatusCode: http.StatusOK,
			},
			mockCall: []*gomock.Call{
				mock.EXPECT().Delete(gomock.Any(), 1).
					Return(nil),
			},
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
			desc:   "error deleting records",
			id:     "1",
			expErr: errors.Error("error deleting record"),
			mockCall: []*gomock.Call{
				mock.EXPECT().Delete(gomock.Any(), 1).Return(errors.Error("error deleting record")),
			},
		},
	}
	for _, tc := range testCases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)
		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})
		resp, err := h.Delete(ctx)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expResp) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expResp, resp)
		}

	}
}

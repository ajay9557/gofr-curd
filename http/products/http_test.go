package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	gerror "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/arohanzst/testapp/models"
	"github.com/arohanzst/testapp/services"
	"github.com/golang/mock/gomock"
)

func Test_ReadByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockProductService := services.NewMockProduct(ctrl)
	handler := New(mockProductService)

	testCases := []struct {
		desc   string
		input  string
		calls  []*gomock.Call
		resp   *models.Product
		expErr error
	}{
		{
			desc:  "Case:1",
			input: "1",
			resp: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			},
			calls: []*gomock.Call{
				mockProductService.EXPECT().ReadByID(gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),
			},
			expErr: nil,
		},
		{
			desc:  "Case:2",
			input: "10",
			resp:  nil,
			expErr: gerror.EntityNotFound{
				Entity: "Product",
				ID:     "10",
			},
			calls: []*gomock.Call{
				mockProductService.EXPECT().ReadByID(gomock.Any(), 10).Return(nil, gerror.EntityNotFound{
					Entity: "Product",
					ID:     "10",
				}),
			},
		},
		{
			desc:  "Case:3",
			input: "-1",
			resp:  nil,
			expErr: gerror.InvalidParam{
				Param: []string{"id"},
			},
			calls: []*gomock.Call{
				mockProductService.EXPECT().ReadByID(gomock.Any(), gomock.Any()).Return(nil, gerror.InvalidParam{
					Param: []string{"id"},
				}),
			},
		},
		{
			desc:  "Case:4",
			input: "dededed",
			resp:  nil,
			expErr: gerror.InvalidParam{
				Param: []string{"id"},
			},
		},
		{
			desc:  "Case:5",
			input: "",
			resp:  nil,
			expErr: gerror.MissingParam{
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

			//id, err := strconv.Atoi(tc.input)

			resp, err := handler.ReadByIdHandler(ctx)
			if !reflect.DeepEqual(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}
			if tc.expErr == nil && !reflect.DeepEqual(resp, tc.resp) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.resp, resp)
			}
		})
	}
}

func Test_Read(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockProductService := services.NewMockProduct(ctrl)
	handler := New(mockProductService)

	testCases := []struct {
		desc   string
		calls  []*gomock.Call
		resp   []models.Product
		expErr error
	}{
		{
			desc: "Case:1-Success",
			resp: []models.Product{{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			}},
			calls: []*gomock.Call{
				mockProductService.EXPECT().Read(gomock.Any()).Return([]models.Product{{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}}, nil),
			},
			expErr: nil,
		},
		{
			desc:   "Case:2-Failure, Internal Server Error",
			resp:   nil,
			expErr: errors.New("Internal Server Error"),
			calls: []*gomock.Call{
				mockProductService.EXPECT().Read(gomock.Any()).Return(nil, errors.New("Internal Server Error")),
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

			//id, err := strconv.Atoi(tc.input)

			resp, err := handler.ReadHandler(ctx)
			if !reflect.DeepEqual(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}
			if tc.expErr == nil && !reflect.DeepEqual(resp, tc.resp) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.resp, resp)
			}
		})
	}
}

func Test_Create(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockProductService := services.NewMockProduct(ctrl)
	handler := New(mockProductService)

	testCases := []struct {
		desc   string
		calls  []*gomock.Call
		resp   *models.Product
		body   models.Product
		expErr error
	}{
		{
			desc: "Case:1-Success",
			resp: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			},
			calls: []*gomock.Call{
				mockProductService.EXPECT().Create(gomock.Any(), &models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),
			},
			body: models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			},
			expErr: nil,
		},
		{
			desc:   "Case:2-Failure, Internal Server Error",
			resp:   nil,
			expErr: errors.New("Internal Server Error"),
			calls: []*gomock.Call{
				mockProductService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("Internal Server Error")),
			},
		},
		{
			desc:   "Case:3-Failure, Invalid Body",
			resp:   nil,
			expErr: gerror.MissingParam{Param: []string{"Name", "Type"}},
			body:   models.Product{},
			calls: []*gomock.Call{
				mockProductService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, gerror.MissingParam{Param: []string{"Name", "Type"}}),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			b, _ := json.Marshal(tc.body)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "http://dummy", bytes.NewReader(b))

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			//id, err := strconv.Atoi(tc.input)

			resp, err := handler.CreateHandler(ctx)
			if !reflect.DeepEqual(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}
			if tc.expErr == nil && !reflect.DeepEqual(resp, tc.resp) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.resp, resp)
			}
		})
	}
}

func Test_Update(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockProductService := services.NewMockProduct(ctrl)
	handler := New(mockProductService)

	testCases := []struct {
		desc   string
		calls  []*gomock.Call
		id     string
		body   models.Product
		resp   *models.Product
		expErr error
	}{
		{
			desc: "Case:1-Success",
			resp: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Daily-Use",
			},
			calls: []*gomock.Call{
				mockProductService.EXPECT().Update(gomock.Any(), gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Daily-Use",
				}, nil),
			},
			id: "1",
			body: models.Product{

				Type: "Daily-Use",
			},
			expErr: nil,
		},
		{
			desc:   "Case:2-Failure, Internal Server Error",
			resp:   nil,
			expErr: errors.New("Internal Server Error"),
			id:     "1",
			body: models.Product{

				Type: "Daily-Use",
			},
			calls: []*gomock.Call{
				mockProductService.EXPECT().Update(gomock.Any(), gomock.Any(), 1).Return(nil, errors.New("Internal Server Error")),
			},
		},
		{
			desc:   "Case:3-Failure, Invalid Body",
			resp:   nil,
			expErr: gerror.MissingParam{Param: []string{"Name", "Type"}},
			id:     "1",
			body:   models.Product{},
			calls: []*gomock.Call{
				mockProductService.EXPECT().Update(gomock.Any(), gomock.Any(), 1).Return(nil, gerror.MissingParam{Param: []string{"Name", "Type"}}),
			},
		},
		{
			desc: "Case:4-Failure, Invalid Id",
			resp: nil,
			expErr: gerror.MissingParam{
				Param: []string{"id"},
			},
			id: "",
			body: models.Product{

				Type: "Daily-Use",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			b, _ := json.Marshal(tc.body)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "http://dummy", bytes.NewReader(b))

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})

			//id, err := strconv.Atoi(tc.input)

			resp, err := handler.UpdateHandler(ctx)
			if !reflect.DeepEqual(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}
			if tc.expErr == nil && !reflect.DeepEqual(resp, tc.resp) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.resp, resp)
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockProductService := services.NewMockProduct(ctrl)
	handler := New(mockProductService)

	testCases := []struct {
		desc    string
		expResp interface{}
		calls   []*gomock.Call
		id      string
		expErr  error
	}{
		{
			desc: "Case:1-Success",
			calls: []*gomock.Call{
				mockProductService.EXPECT().Delete(gomock.Any(), 1).Return(nil),
			},
			id:      "1",
			expErr:  nil,
			expResp: "Deleted successfully",
		},
		{
			desc:   "Case:2-Failure, Internal Server Error",
			expErr: errors.New("Internal Server Error"),
			id:     "1",
			calls: []*gomock.Call{
				mockProductService.EXPECT().Delete(gomock.Any(), 1).Return(errors.New("Internal Server Error")),
			},
			expResp: nil,
		},
		{
			desc: "Case:3-Failure, Invalid Id",
			expErr: gerror.MissingParam{
				Param: []string{"id"},
			},
			id:      "",
			expResp: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "http://dummy", nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})

			//id, err := strconv.Atoi(tc.input)

			resp, err := handler.DeleteHandler(ctx)
			if !reflect.DeepEqual(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}

			if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expResp) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expResp, resp)
			}

		})
	}
}

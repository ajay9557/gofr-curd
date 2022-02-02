package product

import (
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
		resp   *models.Product
		expErr error
	}{
		{
			desc:  "success case",
			input: "1",
			resp: &models.Product{
				Id:   1,
				Name: "test",
				Type: "example",
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

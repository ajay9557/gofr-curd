package user

import (
	ers "errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
	"github.com/shaurya-zopsmart/Gopr-devlopment/Services"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

func TestGetById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := Services.NewMockServiceint(ctrl)
	h := New(mock)

	testCases := []struct {
		desc   string
		input  string
		calls  []*gomock.Call
		resp   *model.Product
		expErr error
	}{
		{
			desc:  "success case",
			input: "1",
			resp: &model.Product{
				Id:   1,
				Name: "test",
				Type: "example",
			},
			calls: []*gomock.Call{
				mock.EXPECT().GetId(gomock.Any(), 1).Return(&model.Product{
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
				mock.EXPECT().GetId(gomock.Any(), 102).Return(nil, errors.EntityNotFound{
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
				mock.EXPECT().GetId(gomock.Any(), -1).Return(nil, errors.InvalidParam{
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
			r := httptest.NewRequest(http.MethodGet, "http://hahalol", nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			ctx.SetPathParams(map[string]string{
				"id": tc.input,
			})

			resp, err := h.GetByID(ctx)
			if !ers.Is(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}
			if tc.expErr == nil && !reflect.DeepEqual(resp, tc.resp) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.resp, resp)
			}
		})
	}
}

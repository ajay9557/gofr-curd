package handler

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"github.com/himanshu-kumar-zs/gofr-curd/services"
	"net/http/httptest"
	"reflect"
	"testing"
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
				1,
				"abc",
				"xyz",
			},
			nil,
			[]*gomock.Call{
				mockSrv.EXPECT().GetByID(gomock.Any(), 1).
					Return(&models.Product{
						1,
						"abc",
						"xyz",
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
		// no need to check id exists or not in handler it will be checked in service layer
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

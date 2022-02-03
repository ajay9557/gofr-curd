package products

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
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
	mock := services.NewMockProduct(ctrl)
	h := New(mock)

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
				mock.EXPECT().ReadByID(gomock.Any(), 1).Return(&models.Product{
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
			expErr: errors.EntityNotFound{
				Entity: "Product",
				ID:     "10",
			},
			calls: []*gomock.Call{
				mock.EXPECT().ReadByID(gomock.Any(), 10).Return(nil, errors.EntityNotFound{
					Entity: "Product",
					ID:     "10",
				}),
			},
		},
		{
			desc:  "Case:3",
			input: "-1",
			resp:  nil,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
			calls: []*gomock.Call{
				mock.EXPECT().ReadByID(gomock.Any(), gomock.Any()).Return(nil, errors.InvalidParam{
					Param: []string{"id"},
				}),
			},
		},
		{
			desc:  "Case:4",
			input: "dededed",
			resp:  nil,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
		},
		{
			desc:  "Case:5",
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

			//id, err := strconv.Atoi(tc.input)

			resp, err := h.ReadByIdHandler(ctx)
			if !reflect.DeepEqual(err, tc.expErr) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
			}
			if tc.expErr == nil && !reflect.DeepEqual(resp, tc.resp) {
				t.Errorf("%s : expected %v, but got %v", tc.desc, tc.resp, resp)
			}
		})
	}
}

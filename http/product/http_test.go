package product

import (
	"fmt"
	"gofr-curd/models"
	"gofr-curd/service"
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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc        string
		id          string
		mock        []*gomock.Call
		expectedRes interface{}
		expectedErr error
	}{
		{
			desc: "Success case",
			id:   "1",
			mock: []*gomock.Call{
				mockService.EXPECT().GetByProductId(gomock.Any(), gomock.Any()).Return(models.Product{
					Id:   1,
					Name: "jeans",
					Type: "clothes",
				}, nil),
			},
			expectedRes: models.Product{
				Id:   1,
				Name: "jeans",
				Type: "clothes",
			},
			expectedErr: nil,
		},
		{
			desc:        "Failure case - empty id",
			id:          "",
			expectedRes: nil,
			expectedErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:        "Failure case - invalid id",
			id:          "1a",
			expectedRes: nil,
			expectedErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Failure case - 3",
			id:   "3",
			mock: []*gomock.Call{
				mockService.EXPECT().GetByProductId(gomock.Any(), gomock.Any()).
					Return(models.Product{}, errors.EntityNotFound{Entity: "product", ID: "3"}),
			},
			expectedRes: nil,
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "3"},
		},
	}

	app := gofr.New()

	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/product", nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": ts.id,
			})
			resp, err := mockHandler.GetById(ctx)
			if !reflect.DeepEqual(ts.expectedRes, resp) {
				fmt.Print("expected ", ts.expectedRes, "obtained", resp)
			}
			if !reflect.DeepEqual(ts.expectedErr, err) {
				fmt.Print("expected ", ts.expectedErr, "obtained", err)
			}
		})
	}
}

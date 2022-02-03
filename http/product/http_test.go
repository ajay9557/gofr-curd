package product

import (
	// "errors"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
	"github.com/tejas/gofr-crud/models"
	"github.com/tejas/gofr-crud/service"
)

func TestGetProductById(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockProductService(ctrl)
	mockHandler := New(mockService)

	testCases := []struct {
		desc     string
		id       string
		mockCall []*gomock.Call
		expOut   interface{}
		expErr   error
	}{
		{
			desc: "Case 1: Success Case",
			id:   "1",
			mockCall: []*gomock.Call{
				mockService.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(models.Product{
					Id:   1,
					Name: "name-1",
					Type: "type-1",
				}, nil),
			},
			expOut: models.Product{
				Id:   1,
				Name: "name-1",
				Type: "type-1",
			},
			expErr: nil,
		},
		{
			desc:   "Case 2: Failure Case1 invalid id",
			id:     "2q",
			expOut: nil,
			expErr: errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Case 3: Failure Case2",
			id:   "99",
			mockCall: []*gomock.Call{
				mockService.EXPECT().GetProductById(gomock.Any(), gomock.Any()).
					Return(models.Product{}, errors.EntityNotFound{Entity: "product", ID: "99"}),
			},
			expOut: nil,
			expErr: errors.EntityNotFound{Entity: "product", ID: "99"},
		},
	}

	app := gofr.New()

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			httpResRec := httptest.NewRecorder()
			httpReq := httptest.NewRequest("GET", "/product", nil)

			req := request.NewHTTPRequest(httpReq)
			res := responder.NewContextualResponder(httpResRec, httpReq)

			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": test.id,
			})

			result, err := mockHandler.GetProductById(ctx)

			if !reflect.DeepEqual(test.expOut, result) {
				fmt.Printf("Expected : %v, Got : %v", test.expOut, result)
			}

			if !reflect.DeepEqual(test.expErr, err) {
				fmt.Printf("Expected : %v, Got : %v", test.expErr, err)
			}

		})
	}

}

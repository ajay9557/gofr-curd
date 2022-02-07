package product

import (
//	"net/http"
	"net/http/httptest"
	"testing"
	models "zopsmart/productgofr/models"
	service "zopsmart/productgofr/services"
//	store "zopsmart/productgofr/stores"

	// "context"
	// "fmt"
	"reflect"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"github.com/golang/mock/gomock"
)


func TestGetUserById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := service.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		desc string
		id        string
		mock      *gomock.Call
		expectedErr  error
		expectedRes interface{}
	}{
		{
			desc: "Success case",
			id: "1",
			expectedErr: nil,
			mock: 
				mockService.EXPECT().GetProdByID(gomock.Any(),1).Return(&models.Product{
					Id:    1,
					Name:  "shirt",
					Type: "fashion",
				}, nil),
			expectedRes: &models.Product{
				Id:    1,
				Name:  "shirt",
				Type: "fashion",
			}, 
		},
		{
			desc: "Failure case",
			id: "id",
			expectedErr: errors.MissingParam{Param: []string{"id"}},
			expectedRes: nil,
		},

		{
			desc: "Failure case",
			id: "",
			expectedErr: errors.MissingParam{Param: []string{"id"}},
			expectedRes: nil,
		},
		{
			desc: "Failure case",
			id: "234",
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "234"},
			expectedRes: nil,
			mock: mockService.EXPECT().GetProdByID(gomock.Any(),1).Return(&models.Product{},
				errors.EntityNotFound{Entity: "product",ID: "234"} ),
		},
	}
	

	for _, tc := range testCases {
		t.Run("sucess test case",func(t *testing.T) {

			w := httptest.NewRecorder()
			url := "/product/" + tc.id
			r := httptest.NewRequest("GET", url, nil)

			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)

			ctx := gofr.NewContext(res, req, app)

			ctx.SetPathParams(map[string]string{
				"id": tc.id,
			})

			resp, err := mock.GetProdByIdHandler(ctx)
			if !reflect.DeepEqual(err,tc.expectedErr) {
				t.Errorf("Expected %v, but got %v", tc.expectedErr, err)
			}

			if !reflect.DeepEqual(resp,tc.expectedRes) {
				t.Errorf("Expected %v, but got %v", tc.expectedRes, resp)
			}
			
		})
	}
	
}
package product

import (
	"bytes"
	goError "errors"
	"net/http"
	"net/http/httptest"
	"product/models"
	"product/services"
	"reflect"
	"strconv"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
)

func Test_GetByIdHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := services.NewMockService(controller)
	mockHandler := New(mockService)

	testCases := []struct {
		desc             string
		input            string
		mock             []*gomock.Call
		expectedError    error
		expectedResponse interface{}
	}{
		{
			desc:  "Test Case 1",
			input: "1",
			mock: []*gomock.Call{
				mockService.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(models.Product{Id: 1, Name: "lg", Type: "machine"}, nil),
			},
			expectedError:    nil,
			expectedResponse: models.Response{Data: models.Product{Id: 1, Name: "lg", Type: "machine"}, Message: "Product Found", StatusCode: 200},
		},
		{
			desc:  "Test Case 2",
			input: "-1",
			mock: []*gomock.Call{
				mockService.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(models.Product{}, goError.New("INVALID ID")),
			},
			expectedError:    errors.EntityNotFound{Entity: "product", ID: "-1"},
			expectedResponse: nil,
		},
	}

	application := gofr.New()

	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			recoder := httptest.NewRecorder()
			testRequest := httptest.NewRequest("GET", "/product", nil)

			req := request.NewHTTPRequest(testRequest)
			res := responder.NewContextualResponder(recoder, testRequest)

			ctx := gofr.NewContext(res, req, application)

			ctx.SetPathParams(map[string]string{
				"id": tcs.input,
			})

			response, err := mockHandler.GetByIdHandler(ctx)
			if !reflect.DeepEqual(tcs.expectedResponse, response) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedResponse, response)
			}
			if !reflect.DeepEqual(tcs.expectedError, err) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedError, err)
			}
		})
	}
}

func Test_GetAllProductHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := services.NewMockService(controller)
	mockHandler := New(mockService)

	testCases := []struct {
		desc             string
		mock             []*gomock.Call
		expectedError    error
		expectedResponse interface{}
	}{
		{
			desc: "Test Case 1",
			mock: []*gomock.Call{
				mockService.EXPECT().GetAllProduct(gomock.Any()).Return([]models.Product{{Id: 1, Name: "lg", Type: "machine"}}, nil),
			},
			expectedError:    nil,
			expectedResponse: models.Response{Data: []models.Product{{Id: 1, Name: "lg", Type: "machine"}}, Message: "Products Found", StatusCode: 200},
		},
	}

	application := gofr.New()

	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			recoder := httptest.NewRecorder()
			testRequest := httptest.NewRequest("GET", "/product", nil)

			req := request.NewHTTPRequest(testRequest)
			res := responder.NewContextualResponder(recoder, testRequest)

			ctx := gofr.NewContext(res, req, application)

			response, err := mockHandler.GetAllProductHandler(ctx)
			if !reflect.DeepEqual(tcs.expectedResponse, response) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedResponse, response)
			}
			if !reflect.DeepEqual(tcs.expectedError, err) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedError, err)
			}
		})
	}
}

func Test_AddProductHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := services.NewMockService(controller)
	mockHandler := New(mockService)

	testCases := []struct {
		desc             string
		input            []byte
		mock             []*gomock.Call
		expectedError    error
		expectedResponse interface{}
	}{
		{
			desc:  "Test Case 1",
			input: []byte(`{"id": 1, "name": "lg", "type": "machine"}`),
			mock: []*gomock.Call{
				mockService.EXPECT().AddProduct(gomock.Any(), models.Product{Id: 1, Name: "lg", Type: "machine"}).Return(nil),
			},
			expectedError:    nil,
			expectedResponse: models.Response{Data: "Product Added", Message: "Saved", StatusCode: 200},
		},
		{
			desc:  "Test Case 2",
			input: []byte(`{"id": -1, "name": "lg", "type": "machine"}`),
			mock: []*gomock.Call{
				mockService.EXPECT().AddProduct(gomock.Any(), models.Product{Id: -1, Name: "lg", Type: "machine"}).Return(goError.New("FAILED TO ADD PRODUCT")),
			},
			expectedError:    errors.EntityNotFound{Entity: "FAILED TO ADD PRODUCT", ID: ""},
			expectedResponse: nil,
		},
	}

	application := gofr.New()

	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			recoder := httptest.NewRecorder()
			testRequest := httptest.NewRequest("POST", "/product/add", bytes.NewReader(tcs.input))

			req := request.NewHTTPRequest(testRequest)
			res := responder.NewContextualResponder(recoder, testRequest)

			ctx := gofr.NewContext(res, req, application)

			response, err := mockHandler.AddProductHandler(ctx)
			if !reflect.DeepEqual(tcs.expectedResponse, response) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedResponse, response)
			}
			if !reflect.DeepEqual(tcs.expectedError, err) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedError, err)
			}
		})
	}
}

func Test_UpdateProductHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := services.NewMockService(controller)
	mockHandler := New(mockService)

	testCases := []struct {
		desc             string
		input            []byte
		mock             []*gomock.Call
		expectedError    error
		expectedResponse interface{}
	}{
		{
			desc:  "Test Case 1",
			input: []byte(`{"id": 1, "name": "lg", "type": "machine"}`),
			mock: []*gomock.Call{
				mockService.EXPECT().UpdateProduct(gomock.Any(), models.Product{Id: 1, Name: "lg", Type: "machine"}).Return(nil),
			},
			expectedError:    nil,
			expectedResponse: models.Response{Data: "Product Updated", Message: "Successfull", StatusCode: http.StatusOK},
		},
		{
			desc:  "Test Case 2",
			input: []byte(`{"id": -1, "name": "lg", "type": "machine"}`),
			mock: []*gomock.Call{
				mockService.EXPECT().UpdateProduct(gomock.Any(), models.Product{Id: -1, Name: "lg", Type: "machine"}).Return(goError.New("FAILED TO ADD PRODUCT")),
			},
			expectedError:    errors.EntityNotFound{Entity: "FAILED TO UPDATE PRODUCT", ID: ""},
			expectedResponse: nil,
		},
	}

	application := gofr.New()

	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			recoder := httptest.NewRecorder()
			testRequest := httptest.NewRequest("PUT", "/product/update", bytes.NewReader(tcs.input))

			req := request.NewHTTPRequest(testRequest)
			res := responder.NewContextualResponder(recoder, testRequest)

			ctx := gofr.NewContext(res, req, application)

			response, err := mockHandler.UpdateProductHandler(ctx)
			if !reflect.DeepEqual(tcs.expectedResponse, response) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedResponse, response)
			}
			if !reflect.DeepEqual(tcs.expectedError, err) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedError, err)
			}
		})
	}
}

func Test_DeleteProductHandler(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockService := services.NewMockService(controller)
	mockHandler := New(mockService)

	testCases := []struct {
		desc             string
		input            int
		mock             []*gomock.Call
		expectedError    error
		expectedResponse interface{}
	}{
		{
			desc:  "Test Case 1",
			input: 1,
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteProduct(gomock.Any(), 1).Return(nil),
			},
			expectedError:    nil,
			expectedResponse: models.Response{Data: "Product Deleted", Message: "Successfull", StatusCode: http.StatusOK},
		},
		{
			desc:             "Test Case 2",
			input:            -1,
			mock:             nil,
			expectedError:    errors.EntityNotFound{Entity: "INVALID INPUTS", ID: ""},
			expectedResponse: nil,
		},
	}

	application := gofr.New()

	for _, tcs := range testCases {
		t.Run(tcs.desc, func(t *testing.T) {
			recoder := httptest.NewRecorder()
			testRequest := httptest.NewRequest("DELETE", "/product/delete", nil)

			req := request.NewHTTPRequest(testRequest)
			res := responder.NewContextualResponder(recoder, testRequest)

			ctx := gofr.NewContext(res, req, application)

			ctx.SetPathParams(map[string]string{
				"id": strconv.Itoa(tcs.input),
			})

			response, err := mockHandler.DeleteProductHandler(ctx)
			if !reflect.DeepEqual(tcs.expectedResponse, response) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedResponse, response)
			}
			if !reflect.DeepEqual(tcs.expectedError, err) {
				t.Errorf("Expected: %s, Got: %s", tcs.expectedError, err)
			}
		})
	}
}

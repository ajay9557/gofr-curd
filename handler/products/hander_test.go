package products

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	models "zopsmart/gofr-curd/model"
	services "zopsmart/gofr-curd/service"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"
)

func Test_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	serv := services.NewMockProductservice(ctrl)
	h := New(serv)
	app := gofr.New()

	tcs := []struct {
		desc           string
		Id             string
		err            error
		expectedOutput interface{}
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   "1",
			err:  nil,
			expectedOutput: models.Product{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			mock: []*gomock.Call{serv.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(models.Product{
				Id: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc: "Failure-2",
			Id:   "45",
			err: errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			},
			expectedOutput: nil,
			mock: []*gomock.Call{serv.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(models.Product{}, errors.EntityNotFound{
				Entity: "Product",
				ID:     "45",
			})},
		},
	}

	for _, tc := range tcs {

		t.Run(tc.desc, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
			req := request.NewHTTPRequest(r)
			res := responder.NewContextualResponder(w, r)
			ctx := gofr.NewContext(res, req, app)
			ctx.SetPathParams(map[string]string{
				"id": tc.Id,
			})
			resp, err := h.GetById(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(resp, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, resp)
			}
		})
	}
}

package products

import (
	"context"
	"gofr-curd/models"
	"gofr-curd/stores"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	gomock "github.com/golang/mock/gomock"
)

func TestGetByUserId(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc           string
		Id             int
		expectedOutput models.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   1,
			expectedOutput: models.Product{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetId(gomock.Any(), gomock.Any()).Return(models.Product{
				Id: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc:           "Failure",
			Id:             0,
			expectedOutput: models.Product{},
			err:            errors.EntityNotFound{Entity: "Product", ID: "0"},
		},
		{
			desc:           "Failure-2",
			Id:             412345,
			expectedOutput: models.Product{},
			err:            errors.EntityNotFound{Entity: "Product", ID: "412345"},
			mock:           []*gomock.Call{mockStore.EXPECT().GetId(gomock.Any(), gomock.Any()).Return(models.Product{}, errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.GetByUserId(ctx, tc.Id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}

}

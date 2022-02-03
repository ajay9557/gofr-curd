package products

import (
	"context"
	"reflect"
	"testing"
	"zopsmart/gofr-curd/model"
	"zopsmart/gofr-curd/store"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	gomock "github.com/golang/mock/gomock"
)

func TestGetByUserId(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProductstorer(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc           string
		Id             string
		expectedOutput model.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   "1",
			expectedOutput: model.Product{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(model.Product{
				Id: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc:           "Failure",
			Id:             "",
			expectedOutput: model.Product{},
			err:            errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:           "Failure-2",
			Id:             "412345",
			expectedOutput: model.Product{},
			err:            errors.EntityNotFound{Entity: "product", ID: "412345"},
			mock:           []*gomock.Call{mockStore.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(model.Product{}, errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
		{
			desc:           "Failure-3",
			Id:             "s",
			expectedOutput: model.Product{},
			err:            errors.InvalidParam{Param: []string{"id"}},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.GetByID(ctx, tc.Id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}

}

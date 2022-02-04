package products

import (
	"context"
	"fmt"
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

func Test_UpdateById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProductstorer(ctrl)
	mock := New(mockStore)
	tcs := []struct {
		desc           string
		Id             string
		input          model.Product
		expectedOutput model.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   "1",
			input: model.Product{
				Id:   0,
				Name: "Shirts",
				Type: "US POLO",
			},
			expectedOutput: model.Product{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().UpdateById(gomock.Any(), gomock.Any()).Return(model.Product{
				Id: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc:           "Failure",
			Id:             "",
			expectedOutput: model.Product{},
			err:            errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Success",
			Id:   "1",
			input: model.Product{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-2",
			Id:   "412345",
			input: model.Product{
				Id:   0,
				Name: "Shirts",
				Type: "US POLO",
			},
			expectedOutput: model.Product{},
			err:            errors.EntityNotFound{Entity: "product", ID: "412345"},
			mock:           []*gomock.Call{mockStore.EXPECT().UpdateById(gomock.Any(), gomock.Any()).Return(model.Product{}, errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
		
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.UpdateById(ctx, tc.input, tc.Id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}

}

func Test_DeleteById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProductstorer(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc string
		Id   string
		err  error
		mock []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   "1",
			err:  nil,
			mock: []*gomock.Call{mockStore.EXPECT().DeleteById(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure",
			Id:   "",
			err:  errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-2",
			Id:   "412345",
			err:  errors.EntityNotFound{Entity: "product", ID: "412345"},
			mock: []*gomock.Call{mockStore.EXPECT().DeleteById(gomock.Any(), gomock.Any()).Return(errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
		{
			desc: "Failure-3",
			Id:   "s",
			err:  errors.InvalidParam{Param: []string{"id"}},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			err := mock.DeleteById(ctx, tc.Id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}

		})
	}

}

func Test_GetProducts(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProductstorer(ctrl)
	mock := New(mockStore)
	tcs := []struct {
		desc           string
		expectedOutput []model.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			expectedOutput: []model.Product{{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			}},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetProducts(gomock.Any()).Return([]model.Product{{
				Id: 1, Name: "Shirts", Type: "US POLO"}}, nil)},
		},
		{
			desc:           "Failure",
			expectedOutput: nil,
			err:            fmt.Errorf("unable to retrieve"),
			mock:           []*gomock.Call{mockStore.EXPECT().GetProducts(gomock.Any()).Return(nil, fmt.Errorf("unable to retrieve"))},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.GetProducts(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}

}

func Test_AddProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockProductstorer(ctrl)
	mock := New(mockStore)
	tcs := []struct {
		desc           string
		input          model.Product
		expectedOutput model.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			input: model.Product{
				Id:   0,
				Name: "Shirts",
				Type: "US POLO",
			},
			expectedOutput: model.Product{
				Id:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err:  nil,
			mock: []*gomock.Call{mockStore.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(1, nil)},
		},
		{
			desc:           "Failure",
			input:          model.Product{Id: 2, Name: "Shirts", Type: "US POLO"},
			expectedOutput: model.Product{},
			err:            errors.InvalidParam{Param: []string{"id"}},
			mock:           nil,
		},
		{
			desc: "Failure-2",
			input: model.Product{
				Id:   0,
				Name: "Shirts",
				Type: "US POLO",
			},
			expectedOutput: model.Product{},
			err:            errors.DB{},
			mock:           []*gomock.Call{mockStore.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(0, errors.DB{})},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.AddProduct(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}

}

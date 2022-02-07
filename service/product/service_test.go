package product

import (
	"context"
	"gofr-curd/model"
	store "gofr-curd/stores"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
)

func TestGetByUserId(t *testing.T) {

	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc           string
		Id             int
		expectedOutput model.ProductDetails
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   1,
			expectedOutput: model.ProductDetails{
				Id:    1,
				Name:  "vivo",
				Types: "smartphone",
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(model.ProductDetails{
				Id: 1, Name: "vivo", Types: "smartphone"}, nil)},
		},
		{
			desc:           "Failure",
			Id:             0,
			expectedOutput: model.ProductDetails{},
			err:            errors.EntityNotFound{Entity: "product", ID: "0"},
		},
		{
			desc:           "Failure-2",
			Id:             412345,
			expectedOutput: model.ProductDetails{},
			err:            errors.EntityNotFound{Entity: "product", ID: "412345"},
			mock:           []*gomock.Call{mockStore.EXPECT().GetProductById(gomock.Any(), gomock.Any()).Return(model.ProductDetails{}, errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.GetByProductId(tc.Id, ctx)
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

	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc           string
		Id             int
		expectedOutput model.ProductDetails
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   1,
			err:  nil,
			mock: []*gomock.Call{mockStore.EXPECT().DeleteProductId(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure",
			Id:   0,
			err:  errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-1",
			Id:   412345,
			err:  errors.Error("Internal Database Error"),
			mock: []*gomock.Call{mockStore.EXPECT().DeleteProductId(gomock.Any(), gomock.Any()).Return(errors.Error("Internal Database Error"))},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			err := mock.DeleteByProductId(ctx, tc.Id)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
		})
	}
}

func Test_UpdateProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc  string
		Id    int
		input model.ProductDetails
		err   error
		mock  []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   1,
			input: model.ProductDetails{
				Id:    1,
				Name:  "vivo",
				Types: "smarthone",
			},
			err:  nil,
			mock: []*gomock.Call{mockStore.EXPECT().UpdateProductById(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure",
			Id:   0,
			input: model.ProductDetails{
				Id:    0,
				Name:  "vivo",
				Types: "smartphone",
			},
			err: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-1",
			Id:   412345,
			input: model.ProductDetails{
				Id:    1,
				Name:  "samsung",
				Types: "smartphone",
			},
			err:  errors.Error("Internal DB error"),
			mock: []*gomock.Call{mockStore.EXPECT().UpdateProductById(gomock.Any(), gomock.Any()).Return(errors.Error("Internal DB error"))},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			err := mock.UpdateByProductId(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
		})
	}
}
func Test_InsertProduct(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc  string
		Id    int
		input model.ProductDetails
		err   error
		mock  []*gomock.Call
	}{
		{
			desc: "Failure",
			Id:   0,
			input: model.ProductDetails{
				Id:    0,
				Name:  "vivo",
				Types: "smartphone",
			},
			err: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Success",
			Id:   1,
			input: model.ProductDetails{
				Id:    1,
				Name:  "vivo",
				Types: "smartphone",
			},
			err: errors.Error("Internal DB error"),
			mock: []*gomock.Call{mockStore.EXPECT().CreateProducts(gomock.Any(), gomock.Any()).Return(model.ProductDetails{
				Id:    1,
				Name:  "vivo",
				Types: "smartphone",
			}, errors.Error("Internal DB error"))},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.InsertProduct(ctx, tc.input)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.input) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.input, res)
			}
		})
	}
}

func Test_GetProducts(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc  string
		Id    int
		input []model.ProductDetails
		err   error
		mock  []*gomock.Call
	}{
		{
			desc: "Success",
			Id:   1,
			input: []model.ProductDetails{{
				Id:    1,
				Name:  "vivo",
				Types: "smartphone",
			},
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetAll(gomock.Any()).Return([]model.ProductDetails{
				{
					Id:    1,
					Name:  "vivo",
					Types: "smartphone",
				},
			}, nil)},
		},
		{
			desc:  "Failure",
			Id:    1,
			input: nil,
			err:   errors.Error("Internal DB error"),
			mock:  []*gomock.Call{mockStore.EXPECT().GetAll(gomock.Any()).Return(nil, errors.Error("Internal DB error"))},
		},
	}
	for _, tc := range tcs {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.GetProducts(ctx)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.input) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.input, res)
			}
		})
	}
}

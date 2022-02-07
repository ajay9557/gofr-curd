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

func TestGetByUserID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductstorer(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc           string
		ID             string
		expectedOutput model.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   "1",
			expectedOutput: model.Product{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).Return(model.Product{
				ID: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc:           "Failure",
			ID:             "",
			expectedOutput: model.Product{},
			err:            errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc:           "Failure-2",
			ID:             "412345",
			expectedOutput: model.Product{},
			err:            errors.EntityNotFound{Entity: "product", ID: "412345"},
			mock: []*gomock.Call{mockStore.EXPECT().GetProductByID(gomock.Any(), gomock.Any()).
				Return(model.Product{}, errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
		{
			desc:           "Failure-3",
			ID:             "s",
			expectedOutput: model.Product{},
			err:            errors.InvalidParam{Param: []string{"id"}},
		},
	}
	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)

		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.GetByID(ctx, tc.ID)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}

func Test_UpdateByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductstorer(ctrl)
	mock := New(mockStore)
	tcs := []struct {
		desc           string
		ID             string
		input          model.Product
		expectedOutput model.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   "1",
			input: model.Product{
				ID:   0,
				Name: "Shirts",
				Type: "US POLO",
			},
			expectedOutput: model.Product{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().UpdateByID(gomock.Any(), gomock.Any()).
				Return(model.Product{
					ID: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc:           "Failure",
			ID:             "",
			expectedOutput: model.Product{},
			err:            errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Success",
			ID:   "1",
			input: model.Product{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-2",
			ID:   "412345",
			input: model.Product{
				ID:   0,
				Name: "Shirts",
				Type: "US POLO",
			},
			expectedOutput: model.Product{},
			err:            errors.EntityNotFound{Entity: "product", ID: "412345"},
			mock: []*gomock.Call{mockStore.EXPECT().UpdateByID(gomock.Any(), gomock.Any()).
				Return(model.Product{}, errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
	}

	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)

		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.UpdateByID(ctx, tc.input, tc.ID)
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
			if !reflect.DeepEqual(res, tc.expectedOutput) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expectedOutput, res)
			}
		})
	}
}

func Test_DeleteByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := store.NewMockProductstorer(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc string
		ID   string
		err  error
		mock []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   "1",
			err:  nil,
			mock: []*gomock.Call{mockStore.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure",
			ID:   "",
			err:  errors.MissingParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-2",
			ID:   "412345",
			err:  errors.EntityNotFound{Entity: "product", ID: "412345"},
			mock: []*gomock.Call{mockStore.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).
				Return(errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
		{
			desc: "Failure-3",
			ID:   "s",
			err:  errors.InvalidParam{Param: []string{"id"}},
		},
	}
	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)

		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			err := mock.DeleteByID(ctx, tc.ID)
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
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			}},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetProducts(gomock.Any()).Return([]model.Product{{
				ID: 1, Name: "Shirts", Type: "US POLO"}}, nil)},
		},
		{
			desc:           "Failure",
			expectedOutput: nil,
			err:            fmt.Errorf("unable to retrieve"),
			mock: []*gomock.Call{mockStore.EXPECT().GetProducts(gomock.Any()).
				Return(nil, fmt.Errorf("unable to retrieve"))},
		},
	}

	for _, test := range tcs {
		tc := test
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
				ID:   0,
				Name: "Shirts",
				Type: "US POLO",
			},
			expectedOutput: model.Product{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err:  nil,
			mock: []*gomock.Call{mockStore.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(1, nil)},
		},
		{
			desc:           "Failure",
			input:          model.Product{ID: 2, Name: "Shirts", Type: "US POLO"},
			expectedOutput: model.Product{},
			err:            errors.InvalidParam{Param: []string{"id"}},
			mock:           nil,
		},
		{
			desc: "Failure-2",
			input: model.Product{
				ID:   0,
				Name: "Shirts",
				Type: "US POLO",
			},
			expectedOutput: model.Product{},
			err:            errors.DB{},
			mock:           []*gomock.Call{mockStore.EXPECT().AddProduct(gomock.Any(), gomock.Any()).Return(0, errors.DB{})},
		},
	}

	for _, test := range tcs {
		tc := test
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

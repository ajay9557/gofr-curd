package products

import (
	"context"
	"reflect"
	"testing"

	"github.com/Training/gofr-curd/models"
	"github.com/Training/gofr-curd/stores"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	gomock "github.com/golang/mock/gomock"
)

func TestGetByUserID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc           string
		ID             int
		expectedOutput models.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   1,
			expectedOutput: models.Product{
				ID:   1,
				Name: "Shirts",
				Type: "US POLO",
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetID(gomock.Any(), gomock.Any()).Return(models.Product{
				ID: 1, Name: "Shirts", Type: "US POLO"}, nil)},
		},
		{
			desc:           "Failure",
			ID:             0,
			expectedOutput: models.Product{},
			err:            errors.EntityNotFound{Entity: "product", ID: "0"},
		},
		{
			desc:           "Failure-2",
			ID:             412345,
			expectedOutput: models.Product{},
			err:            errors.EntityNotFound{Entity: "product", ID: "412345"},
			mock: []*gomock.Call{mockStore.EXPECT().GetID(gomock.Any(), gomock.Any()).Return(
				models.Product{}, errors.EntityNotFound{Entity: "product", ID: "412345"})},
		},
	}
	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.GetByUserID(ctx, tc.ID)

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

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc           string
		ID             int
		expectedOutput models.Product
		err            error
		mock           []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   1,
			err:  nil,
			mock: []*gomock.Call{mockStore.EXPECT().DeleteID(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure",
			ID:   0,
			err:  errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-1",
			ID:   412345,
			err:  errors.Error("Internal DB error"),
			mock: []*gomock.Call{mockStore.EXPECT().DeleteID(gomock.Any(), gomock.Any()).Return(errors.Error("Internal DB error"))},
		},
	}
	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			err := mock.DeleteByProductID(ctx, tc.ID)
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

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc  string
		ID    int
		input models.Product
		err   error
		mock  []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   1,
			input: models.Product{
				ID:   1,
				Name: "Shirts",
				Type: "Lenin",
			},
			err:  nil,
			mock: []*gomock.Call{mockStore.EXPECT().UpdateID(gomock.Any(), gomock.Any()).Return(nil)},
		},
		{
			desc: "Failure",
			ID:   0,
			input: models.Product{
				ID:   0,
				Name: "Shirts",
				Type: "Lenin",
			},
			err: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Failure-1",
			ID:   412345,
			input: models.Product{
				ID:   1,
				Name: "Shirts",
				Type: "Lenin",
			},
			err:  errors.Error("Internal DB error"),
			mock: []*gomock.Call{mockStore.EXPECT().UpdateID(gomock.Any(), gomock.Any()).Return(errors.Error("Internal DB error"))},
		},
	}
	for _, test := range tcs {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			err := mock.UpdateByProductID(ctx, tc.input)
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

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc  string
		ID    int
		input models.Product
		err   error
		mock  []*gomock.Call
	}{
		{
			desc: "Failure",
			ID:   0,
			input: models.Product{
				ID:   0,
				Name: "Shirts",
				Type: "Lenin",
			},
			err: errors.InvalidParam{Param: []string{"id"}},
		},
		{
			desc: "Success",
			ID:   1,
			input: models.Product{
				ID:   1,
				Name: "Shirts",
				Type: "Lenin",
			},
			err: errors.Error("Internal DB error"),
			mock: []*gomock.Call{mockStore.EXPECT().CreateProducts(gomock.Any(), gomock.Any()).Return(models.Product{
				ID:   1,
				Name: "Shirts",
				Type: "Lenin",
			}, errors.Error("Internal DB error"))},
		},
	}
	for _, test := range tcs {
		tc := test
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

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc  string
		ID    int
		input []models.Product
		err   error
		mock  []*gomock.Call
	}{
		{
			desc: "Success",
			ID:   1,
			input: []models.Product{{
				ID:   1,
				Name: "Shirts",
				Type: "Lenin",
			},
			},
			err: nil,
			mock: []*gomock.Call{mockStore.EXPECT().GetAll(gomock.Any()).Return([]models.Product{
				{
					ID:   1,
					Name: "Shirts",
					Type: "Lenin",
				},
			}, nil)},
		},
		{
			desc:  "Failure",
			ID:    1,
			input: nil,
			err:   errors.Error("Internal DB error"),
			mock:  []*gomock.Call{mockStore.EXPECT().GetAll(gomock.Any()).Return(nil, errors.Error("Internal DB error"))},
		},
	}
	for _, test := range tcs {
		tc := test
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

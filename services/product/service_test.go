package product

import (
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
	"github.com/himanshu-kumar-zs/gofr-curd/models"
	"github.com/himanshu-kumar-zs/gofr-curd/store"
)

var prod = &models.Product{
	ID:   1,
	Name: "EliteBook",
	Type: "Laptop",
}

func TestProductService_GetByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	service := New(mockStore)

	testcases := []struct {
		desc     string
		inp      int
		exp      *models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc: "id exists",
			inp:  1,
			exp: &models.Product{
				ID:   1,
				Name: "legion",
				Type: "Laptop",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).
					Return(&models.Product{
						ID:   1,
						Name: "legion",
						Type: "Laptop",
					}, nil),
			},
		},
		{
			desc: "id does not exists",
			inp:  1,
			exp:  nil,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1",
			},
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).
					Return(nil, errors.EntityNotFound{Entity: "product", ID: "1"}),
			},
		},
		{
			desc: "invalid id",
			inp:  -1,
			exp:  nil,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
		},
	}
	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)

		out, err := service.GetByID(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.expErr, err)
		}

		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestProductService_Update(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	service := New(mockStore)

	testcases := []struct {
		desc     string
		inp      *models.Product
		exp      *models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc:   "id exists",
			inp:    prod,
			exp:    prod,
			expErr: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).Return(prod, nil),
				mockStore.EXPECT().Update(gomock.Any(), prod).
					Return(nil),
				mockStore.EXPECT().GetByID(gomock.Any(), 1).Return(prod, nil),
			},
		},
		{
			desc: "id does not exists",
			inp:  prod,
			exp:  nil,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1",
			},
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).
					Return(nil, errors.EntityNotFound{Entity: "product", ID: "1"}),
			},
		},
		{
			desc: "invalid id",
			inp: &models.Product{
				ID:   -1,
				Name: "legion",
				Type: "Laptop",
			},
			exp: nil,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
		},
		{
			desc:   "could not update",
			inp:    prod,
			exp:    nil,
			expErr: errors.Error("could not update"),
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).Return(prod, nil),
				mockStore.EXPECT().Update(gomock.Any(), prod).
					Return(errors.Error("could not update")),
			},
		},
	}
	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)

		out, err := service.Update(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.expErr, err)
		}

		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestProductService_Delete(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	service := New(mockStore)

	testcases := []struct {
		desc     string
		inp      int
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc:   "id exists",
			inp:    1,
			expErr: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).Return(prod, nil),
				mockStore.EXPECT().Delete(gomock.Any(), 1).
					Return(nil),
			},
		},
		{
			desc: "id does not exists",
			inp:  1,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1",
			},
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetByID(gomock.Any(), 1).
					Return(nil, errors.EntityNotFound{Entity: "product", ID: "1"}),
			},
		},
		{
			desc: "invalid id",
			inp:  -1,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
		},
	}
	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)

		err := service.Delete(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
	}
}

func TestProductService_Create(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	service := New(mockStore)

	testcases := []struct {
		desc     string
		inp      *models.Product
		exp      *models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc:   "success case",
			inp:    prod,
			exp:    prod,
			expErr: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().Create(gomock.Any(), prod).
					Return(prod, nil),
			},
		},
		{
			desc: "missing name case",
			inp:  &models.Product{Type: "abc"},
			exp:  nil,
			expErr: errors.MissingParam{
				Param: []string{"name"},
			},
		},
		{
			desc: "missing type",
			inp: &models.Product{
				Name: "legion"},
			exp: nil,
			expErr: errors.MissingParam{
				Param: []string{"type"},
			},
		},
	}
	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)

		out, err := service.Create(ctx, tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.expErr, err)
		}

		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestGetAll(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockStore(ctrl)
	service := New(mockStore)

	testcases := []struct {
		desc     string
		exp      []*models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc:   "success case",
			exp:    []*models.Product{prod},
			expErr: nil,
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetAll(gomock.Any()).
					Return([]*models.Product{prod}, nil),
			},
		},
		{
			desc: "no row case",
			exp:  nil,
			expErr: errors.EntityNotFound{
				Entity: "product",
			},
			mockCall: []*gomock.Call{
				mockStore.EXPECT().GetAll(gomock.Any()).
					Return([]*models.Product{}, errors.EntityNotFound{Entity: "product"}),
			},
		},
	}
	for _, tcs := range testcases {
		ctx := gofr.NewContext(nil, nil, app)

		out, err := service.GetAll(ctx)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.expErr, err)
		}

		if err == nil && !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v , expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

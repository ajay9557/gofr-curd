package product

import (
	"gofr-curd/models"
	"gofr-curd/store"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
)

func TestServices_GetById(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)

	testCases := []struct {
		desc     string
		input    int
		mockCall []*gomock.Call
		expOut   *models.Product
		expErr   error
	}{
		{
			desc:  "valid id",
			input: 1,
			expOut: &models.Product{
				Id:   1,
				Name: "Part1",
				Type: "hardware",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "Part1",
					Type: "hardware",
				}, nil),
			},
		},
		{
			desc:  "negative id",
			input: -1,
			expErr: errors.InvalidParam{
				Param: []string{"id"},
			},
		},
		{
			desc:  "id not in database",
			input: 1002,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1002",
			},
			mockCall: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 1002).Return(nil,
					errors.EntityNotFound{
						Entity: "product",
						ID:     "1002",
					}),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)

		resp, err := s.GetById(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, resp)
		}

	}
}

func TestServices_Get(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)

	testCases := []struct {
		desc     string
		expOut   []*models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc: "success case",
			expOut: []*models.Product{
				&models.Product{
					Id:   1,
					Name: "test",
					Type: "example",
				},
				&models.Product{
					Id:   2,
					Name: "this",
					Type: "that",
				},
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().Get(gomock.Any()).
					Return([]*models.Product{
						&models.Product{
							Id:   1,
							Name: "test",
							Type: "example",
						},
						&models.Product{
							Id:   2,
							Name: "this",
							Type: "that",
						},
					}, nil),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)

		resp, err := s.Get(ctx)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, resp)
		}
	}
}

func TestServices_Create(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)

	testCases := []struct {
		desc     string
		input    models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc: "success",
			input: models.Product{
				Id:   3,
				Name: "mouse",
				Type: "electronics",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().Create(gomock.Any(), models.Product{
					Id:   3,
					Name: "mouse",
					Type: "electronics",
				}).Return(nil),
				mock.EXPECT().GetById(gomock.Any(), 3).Return(&models.Product{
					Id:   3,
					Name: "mouse",
					Type: "electronics",
				}, nil),
			},
		},
		{
			desc: "invalid param id",
			input: models.Product{
				Id: -1,
			},
			expErr:   errors.InvalidParam{Param: []string{"id"}},
			mockCall: nil,
		},
		{
			desc: "error entity already exists",
			input: models.Product{
				Id:   1,
				Name: "mouse",
				Type: "electronics",
			},
			expErr: errors.EntityAlreadyExists{},
			mockCall: []*gomock.Call{
				mock.EXPECT().Create(gomock.Any(), models.Product{
					Id:   1,
					Name: "mouse",
					Type: "electronics",
				}).Return(errors.EntityAlreadyExists{}),
			},
		},
		{
			desc: "error creating product",
			input: models.Product{
				Id:   3,
				Name: "mouse",
				Type: "electronics",
			},
			expErr: errors.EntityNotFound{Entity: "product", ID: "3"},
			mockCall: []*gomock.Call{
				mock.EXPECT().Create(gomock.Any(), models.Product{
					Id:   3,
					Name: "mouse",
					Type: "electronics",
				}).Return(nil),
				mock.EXPECT().GetById(gomock.Any(), 3).Return(nil, errors.EntityNotFound{Entity: "product", ID: "3"}),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)

		resp, err := s.Create(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, &tc.input) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, &tc.input, resp)
		}
	}
}

func TestServices_Update(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)

	testCases := []struct {
		desc     string
		input    models.Product
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc: "success",
			input: models.Product{
				Id:   1,
				Name: "mouse",
				Type: "electronics",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().Update(gomock.Any(), models.Product{
					Id:   1,
					Name: "mouse",
					Type: "electronics",
				}).Return(nil),
				mock.EXPECT().GetById(gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "mouse",
					Type: "electronics",
				}, nil),
			},
		},
		{
			desc: "invalid param id",
			input: models.Product{
				Id: -1,
			},
			expErr:   errors.InvalidParam{Param: []string{"id"}},
			mockCall: nil,
		},
		{
			desc: "error updating record",
			input: models.Product{
				Id:   1,
				Name: "test",
				Type: "example",
			},
			expErr: errors.Error("error updating record"),
			mockCall: []*gomock.Call{
				mock.EXPECT().Update(gomock.Any(), models.Product{
					Id:   1,
					Name: "test",
					Type: "example",
				}).Return(errors.Error("error updating record")),
			},
		},
		{
			desc: "entity not found",
			input: models.Product{
				Id:   100,
				Name: "this",
				Type: "that",
			},
			expErr: errors.EntityNotFound{Entity: "product", ID: "100"},
			mockCall: []*gomock.Call{
				mock.EXPECT().Update(gomock.Any(), models.Product{
					Id:   100,
					Name: "this",
					Type: "that",
				}).Return(nil),
				mock.EXPECT().GetById(gomock.Any(), 100).Return(nil, errors.EntityNotFound{Entity: "product", ID: "100"}),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)

		resp, err := s.Update(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, &tc.input) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, &tc.input, resp)
		}
	}
}

func TestServices_Delete(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := store.NewMockStore(ctrl)
	s := New(mock)

	testCases := []struct {
		desc     string
		id       int
		expErr   error
		mockCall []*gomock.Call
	}{
		{
			desc:   "success",
			id:     1,
			expErr: nil,
			mockCall: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "test",
					Type: "example",
				}, nil),
				mock.EXPECT().Delete(gomock.Any(), 1).Return(nil),
			},
		},
		{
			desc:   "entity not found",
			id:     3,
			expErr: errors.EntityNotFound{Entity: "product", ID: "3"},
			mockCall: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 3).Return(nil, errors.EntityNotFound{Entity: "product", ID: "3"}),
			},
		},
		{
			desc:     "invalid param id",
			id:       -1,
			expErr:   errors.InvalidParam{Param: []string{"id"}},
			mockCall: nil,
		},
		{
			desc:   "error deleting product",
			id:     1,
			expErr: errors.Error("error deleting record"),
			mockCall: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 1).Return(&models.Product{
					Id:   1,
					Name: "test",
					Type: "example",
				}, nil),
				mock.EXPECT().Delete(gomock.Any(), 1).Return(errors.Error("error deleting record")),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)

		err := s.Delete(ctx, tc.id)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}

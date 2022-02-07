package product

import (
	"context"
	"gofr-curd/models"
	"gofr-curd/store"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()

	testGet(t, app)
	testGetByID(t, app)
	testCreate(t, app)
	testUpdate(t, app)
	testDelete(t, app)
}

func testGet(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc   string
		expErr error
		expOut []*models.Product
	}{
		{
			desc:   "success case",
			expErr: nil,
			expOut: []*models.Product{
				{
					ID:   1,
					Name: "test",
					Type: "example",
				},
				{
					ID:   2,
					Name: "this",
					Type: "that",
				},
			},
		},
	}
	for _, test := range testCases {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		s := New()
		out, err := s.Get(ctx)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

func testGetByID(t *testing.T, app *gofr.Gofr) {
	testCases := []struct {
		desc   string
		input  int
		expErr error
		expOut *models.Product
	}{
		{
			desc:   "success case",
			input:  1,
			expErr: nil,
			expOut: &models.Product{
				ID:   1,
				Name: "test",
				Type: "example",
			},
		},
		{
			desc:  "entity not in database",
			input: 1022,
			expErr: errors.EntityNotFound{
				Entity: "product",
				ID:     "1022",
			},
		},
	}

	for _, test := range testCases {
		tc := test
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		s := New()

		out, err := s.GetByID(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

		if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
		}
	}
}

type testCase struct {
	desc   string
	input  models.Product
	expErr error
}

func setTest(app *gofr.Gofr) (store.Store, *gofr.Context) {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()
	s := New()

	return s, ctx
}

func testCreate(t *testing.T, app *gofr.Gofr) {
	tesCases := []testCase{
		{
			desc: "success case",
			input: models.Product{
				ID:   3,
				Name: "this",
				Type: "that",
			},
			expErr: nil,
		},
	}

	for _, tc := range tesCases {
		s, ctx := setTest(app)
		err := s.Create(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}

func testUpdate(t *testing.T, app *gofr.Gofr) {
	tesCases := []testCase{
		{
			desc: "success case",
			input: models.Product{
				ID:   3,
				Name: "hello",
				Type: "world",
			},
			expErr: nil,
		},
	}

	for _, tc := range tesCases {
		s, ctx := setTest(app)
		err := s.Update(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}

func testDelete(t *testing.T, app *gofr.Gofr) {
	tesCases := []struct {
		desc   string
		input  int
		expErr error
	}{
		{
			desc:   "success case",
			input:  3,
			expErr: nil,
		},
	}

	for _, tc := range tesCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()
		s := New()
		err := s.Delete(ctx, tc.input)

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
	}
}

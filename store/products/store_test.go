package products

import (
	"context"

	"reflect"
	"testing"
	"zopsmart/gofr-curd/model"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	//	"github.com/DATA-DOG/go-sqlmock"
)

func Test_GetById(t *testing.T) {
	tests := []struct {
		desc   string
		id     int
		err    error
		output model.Product
	}{
		{
			desc: "Get existent id",
			id:   1,
			err:  nil,
			output: model.Product{
				Id:   1,
				Name: "Reebok",
				Type: "Bats",
			},
		},
		{
			desc: "Get no-existent id",
			id:   45,
			err: errors.EntityNotFound{
				Entity: "product",
				ID:     "45",
			},
			output: model.Product{},
		},
	}
	for _, tc := range tests {
		app := gofr.New()

		//app.ORM = database
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()
		res, err := store.GetProductById(ctx, tc.id)
		if !reflect.DeepEqual(res, tc.output) {
			t.Errorf("expected %v got %v", tc.output, res)
		}
		if tc.err != err {
			t.Errorf("expected %s got %s", tc.err, err)
		}
	}
}

func Test_GetProducts(t *testing.T) {
	tests := []struct {
		desc   string
		err    error
		output []model.Product
	}{
		{
			desc: "Success",
			err:  nil,
			output: []model.Product{
				{
					Id:   1,
					Name: "Reebok",
					Type: "Bats",
				}, {
					Id:   2,
					Name: "Mehfil",
					Type: "Biryani",
				},
			},
		},
	}
	for _, tc := range tests {
		app := gofr.New()

		//app.ORM = database
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()
		res, err := store.GetProducts(ctx)
		if !reflect.DeepEqual(res, tc.output) {
			t.Errorf("expected %v got %v", tc.output, res)
		}
		if tc.err != err {
			t.Errorf("expected %s got %s", tc.err, err)
		}
	}
}

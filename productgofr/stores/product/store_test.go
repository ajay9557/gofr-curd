package product

import (
	"context"
	models "zopsmart/productgofr/models"
//	stores "zopsmart/productgofr/stores"
 //   gofrErr "errors"
	"reflect"
	"testing"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)



func TestGetAll(t *testing.T) {
	app := gofr.New()
testCases := []struct {
	desc   string
	expErr error
	expOut []models.Product
}{
	{
		desc:   "success case",
		expErr: nil,
		expOut: []models.Product{
			{
				Id:   1,
				Name: "shirt",
				Type: "fashion",
			},
			{
				Id:   2,
				Name: "mobile",
				Type: "electronics",
			},
		},
	},
	}

for _, tc := range testCases {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	s := New()
	res, err := s.GetAllProduct(ctx)

	if !reflect.DeepEqual(err, tc.expErr) {
		t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
	}

	if tc.expErr == nil && !reflect.DeepEqual(res, tc.expOut) {
		t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, res)
	}
}
}

func TestGetByID(t *testing.T) {
	app := gofr.New()
testCases := []struct {
	desc   string
	input  int
	expErr error
	expOut models.Product
}{
	{
		desc:   "Success Case",
		input:  1,
		expErr: nil,
		expOut: models.Product{
			Id:   1,
			Name: "shirt",
			Type: "fashion",
		},
	},
	{
		desc:  "Failure case 1",
		input: 10,
		expErr: errors.EntityNotFound{
			Entity: "product",
			ID:     "10",
		},
	},
	{
		desc:  "Failure case 2",
		input: -1,
		expErr: errors.EntityNotFound{
			Entity: "product",
			ID:     "-1",
		},
	},
}

for _, tc := range testCases {
	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	s := New()

	out, err := s.GetProdByID(ctx, tc.input)
	if !reflect.DeepEqual(err, tc.expErr) {
		t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
	}

	if tc.expErr == nil && !reflect.DeepEqual(out, tc.expOut) {
		t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, out)
	}
}
}



func TestCreate(t *testing.T) {
	app := gofr.New()
tesCases := []struct {
	desc   string
	input  models.Product
	expErr error
	expRed models.Product
	}{
	{
		desc: "success case",
		input: models.Product{
			Id:   3,
			Name: "harry potter",
			Type: "books",
		},
		expErr: nil,
	},
	// {
	// 	desc: "Failure case",
	// 	input: models.Product{
	// 		Name: "",
	// 		Type: "",
	// 	},
	// 	expErr: errors.EntityNotFound{
	// 		Entity: "product",
	// 	},
	// },
}

for _, tc := range tesCases {
	
	ctx := gofr.NewContext(nil, nil, app)
    ctx.Context = context.Background()
    s := New()
	err := s.CreateProduct(ctx, tc.input)

	if !reflect.DeepEqual(err, tc.expErr) {
		t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
	}
}
}

func TestUpdate(t *testing.T) {
	app := gofr.New()
tesCases := []struct{
	 desc   string
	 input  models.Product
	 expErr error
	 } {
	{
		desc: "success case",
		input: models.Product{
			Id:   3,
			Name: "",
			Type: "daily needs",
		},
		expErr: nil,
	},
}

for _, tc := range tesCases {
	ctx := gofr.NewContext(nil, nil, app)
    ctx.Context = context.Background()
    s := New()
	err := s.UpdateProduct(ctx, tc.input)

	if !reflect.DeepEqual(err, tc.expErr) {
		t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
	}
}
}

func TestDelete(t *testing.T) {
	app := gofr.New()
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
	err := s.DeleteProduct(ctx, tc.input)

	if !reflect.DeepEqual(err, tc.expErr) {
		t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
	}
}
}
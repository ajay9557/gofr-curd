package products

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func TestCoreLayer(t *testing.T) {
	app := gofr.New()
	testGetProductByID(t, app)
}

func testGetProductByID(t *testing.T, app *gofr.Gofr) {
	tests := []struct {
		desc string
		id   int
		err  error
	}{
		{"Get existent id", 1, nil},
		{"Get non existent id", 100, errors.EntityNotFound{Entity: "products", ID: "100"}},
	}

	for i, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := gofr.NewContext(nil, nil, app)
			ctx.Context = context.Background()

			store := New()

			_, err := store.GetById(ctx, tc.id)
			assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
		})
	}
}

package user

import (
	"errors"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
	"github.com/shaurya-zopsmart/Gopr-devlopment/Stores"
	"github.com/shaurya-zopsmart/Gopr-devlopment/model"
)

func Test_GetId(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	mock := Stores.NewMockStoreint(ctrl)
	s := New(mock)
	testcase := []struct {
		desc   string
		inp    int
		mock   []*gomock.Call
		expout interface{}
		experr error
	}{
		{
			desc: "test case 1",
			inp:  2,
			expout: model.Product{
				Id:   2,
				Name: "gigihadid",
				Type: "libral",
			},
			experr: nil,
			mock: []*gomock.Call{
				mock.EXPECT().GetById(gomock.Any(), 2).Return(model.Product{
					Id:   2,
					Name: "gigihadid",
					Type: "libral",
				}, nil),
			},
		},
		{
			desc:   "test case 2",
			inp:    -1,
			mock:   nil,
			expout: nil,
			experr: errors.New("invalid id"),
		},
	}
	for _, tcs := range testcase {
		ctx := gofr.NewContext(nil, nil, app)
		result, err := s.GetId(ctx, tcs.inp)
		if !errors.Is(err, tcs.experr) {
			t.Errorf("Expected: %s, Output: %s", tcs.experr, err)
		}
		if tcs.experr == nil && !reflect.DeepEqual(result, tcs.expout) {
			t.Errorf("Expected: %v, Output: %v", tcs.expout, result)
		}
	}
}

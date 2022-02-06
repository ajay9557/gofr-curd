package products

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/arohanzst/testapp/models"
	"github.com/arohanzst/testapp/stores"

	gerror "developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
)

func Test_ReadByID(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductStore := stores.NewMockProduct(ctrl)
	mockProductService := New(mockProductStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	testCases := []struct {
		desc     string
		input    int
		mockCall []*gomock.Call
		expOut   *models.Product
		expErr   error
	}{
		{
			desc:  "Case:1",
			input: 1,
			expOut: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().ReadByID(ctx, 1).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),
			},
		},
		{
			desc:   "Case:2",
			input:  -10,
			expErr: errors.New("Invalid Id"),
		},
		{
			desc:  "Case:3",
			input: 10,
			expErr: gerror.EntityNotFound{
				Entity: "Product",
				ID:     "10",
			},
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().ReadByID(ctx, 10).Return(nil, gerror.EntityNotFound{
					Entity: "Product",
					ID:     "10",
				}),
			},
		},
	}

	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		resp, err := mockProductService.ReadByID(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, resp)
		}

	}
}

func Test_Read(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductStore := stores.NewMockProduct(ctrl)
	mockProductService := New(mockProductStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	testCases := []struct {
		desc     string
		mockCall []*gomock.Call
		expOut   []models.Product
		expErr   error
	}{
		{
			desc: "Case:1-Success",
			expOut: []models.Product{{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery"},
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().Read(ctx).Return([]models.Product{{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery"},
				}, nil),
			},
		},
		{
			desc:   "Case:2-Failure",
			expOut: nil,
			expErr: errors.New("Internal Server error: Empty Database"),
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().Read(ctx).Return(nil, errors.New("Internal Server error: Empty Database")),
			},
		},
	}

	for _, tc := range testCases {

		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		resp, err := mockProductService.Read(ctx)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, resp)
		}

	}
}

func Test_Create(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductStore := stores.NewMockProduct(ctrl)
	mockProductService := New(mockProductStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	testCases := []struct {
		desc     string
		input    *models.Product
		mockCall []*gomock.Call
		expOut   *models.Product
		expErr   error
	}{

		{
			desc: "Case:1-Sucess",
			input: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			},
			expOut: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().Create(ctx, gomock.Any()).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),
			},
		},

		{
			desc: "Case:2-Failure, Invalid Entity Name",
			input: &models.Product{
				Id:   1,
				Name: "",
				Type: "Grocery",
			},
			expOut: nil,
			expErr: errors.New("Invalid name or Type"),
		},

		{
			desc:   "Case:3-Failure, Invalid Entity",
			input:  nil,
			expOut: nil,
			expErr: errors.New("Invalid Entity"),
		},

		{
			desc: "Case:4-Failure, Invalid Query",
			input: &models.Product{
				Id:   1,
				Name: "Biscuit",
				Type: "Grocery",
			},
			expOut: nil,
			expErr: errors.New("Invalid Query"),
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().Create(ctx, gomock.Any()).Return(nil, errors.New("Invalid Query")),
			},
		},
	}
	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		resp, err := mockProductService.Create(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, resp)
		}

	}
}

func Test_Update(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductStore := stores.NewMockProduct(ctrl)
	mockProductService := New(mockProductStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	testCases := []struct {
		desc     string
		input    *models.Product
		id       int
		mockCall []*gomock.Call
		expOut   *models.Product
		expErr   error
	}{

		{
			desc: "Case:1-Sucess",
			input: &models.Product{

				Name: "Rice",
			},
			id: 1,
			expOut: &models.Product{
				Id:   1,
				Name: "Rice",
				Type: "Grocery",
			},
			expErr: nil,
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().ReadByID(ctx, 1).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),

				mockProductStore.EXPECT().Update(ctx, &models.Product{

					Name: "Rice",
				}, 1).Return(&models.Product{
					Id:   1,
					Name: "Rice",
					Type: "Grocery",
				}, nil),
			},
		},
		{
			desc: "Case:2-Failure Invalid Id",
			input: &models.Product{

				Name: "Rice",
			},
			id:     10,
			expOut: nil,
			expErr: errors.New("Invalid Id"),
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().ReadByID(ctx, 10).Return(nil, errors.New("Invalid Id")),
			},
		},
		{
			desc:   "Case:3-Failure, Invalid Entity",
			input:  nil,
			expOut: nil,
			expErr: errors.New("Invalid Entity"),
		},

		{
			desc: "Case:4-Failure, Invalid Id",
			input: &models.Product{

				Name: "Rice",
			},
			id:     -1,
			expOut: nil,
			expErr: errors.New("Invalid Id"),
		},

		{
			desc: "Case:5-Failure Invalid Query",
			input: &models.Product{

				Name: "Rice",
			},
			id:     1,
			expOut: nil,
			expErr: errors.New("Invalid Query"),
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().ReadByID(ctx, 1).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),

				mockProductStore.EXPECT().Update(ctx, &models.Product{

					Name: "Rice",
				}, 1).Return(nil, errors.New("Invalid Query")),
			},
		},
	}
	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		resp, err := mockProductService.Update(ctx, tc.input, tc.id)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}
		if tc.expErr == nil && !reflect.DeepEqual(resp, tc.expOut) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expOut, resp)
		}

	}
}

func Test_Delete(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductStore := stores.NewMockProduct(ctrl)
	mockProductService := New(mockProductStore)

	ctx := gofr.NewContext(nil, nil, app)
	ctx.Context = context.Background()

	testCases := []struct {
		desc     string
		id       int
		mockCall []*gomock.Call
		expErr   error
	}{

		{
			desc:   "Case:1-Sucess",
			id:     1,
			expErr: nil,
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().ReadByID(ctx, 1).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),

				mockProductStore.EXPECT().Delete(ctx, 1).Return(nil),
			},
		},
		{
			desc:   "Case:2-Failure Invalid Id",
			id:     10,
			expErr: errors.New("Invalid Id"),
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().ReadByID(ctx, 10).Return(nil, errors.New("Invalid Id")),
			},
		},

		{
			desc:   "Case:3-Failure, Invalid Id",
			id:     -1,
			expErr: errors.New("Invalid Id"),
		},

		{
			desc:   "Case:4-Failure Invalid Query",
			id:     1,
			expErr: errors.New("Invalid Query"),
			mockCall: []*gomock.Call{
				mockProductStore.EXPECT().ReadByID(ctx, 1).Return(&models.Product{
					Id:   1,
					Name: "Biscuit",
					Type: "Grocery",
				}, nil),

				mockProductStore.EXPECT().Delete(ctx, 1).Return(errors.New("Invalid Query")),
			},
		},
	}
	for _, tc := range testCases {
		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		err := mockProductService.Delete(ctx, tc.id)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%s : expected %v, but got %v", tc.desc, tc.expErr, err)
		}

	}
}

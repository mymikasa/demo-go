package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//
//func ErrorNil(err error) assert.ErrorAssertionFunc {
//	return func(t assert.TestingT, erro error, msg ...interface{}) bool {
//		return assert.Equalf(t, err, errors.New("nil"), "nil")
//	}
//}

//(type func(t assert.TestingT, erro error, msg string) bool)
//  type ErrorAssertionFunc func(TestingT, error, ...interface{}) bool

//func WithShutDownTimeOut(shutdownTimeout time.Duration) Option {
//	return func(app *App) {
//		app.shutdownTimeout = shutdownTimeout
//	}
//}

func TestIterateFuncs(t *testing.T) {
	type args struct {
		val any
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*FuncInfo
		wantErr error
	}{
		{
			name:    "nil",
			wantErr: errors.New("nil"),
		},
		{
			name: "basic type",
			args: args{
				val: 123,
			},
			wantErr: errors.New("不支持类型"),
		},
		{
			name: "struct type",
			args: args{
				val: Order{
					buyer:  18,
					seller: 100,
				},
			},
			want: map[string]*FuncInfo{
				"GetBuyer": {
					Name:   "GetBuyer",
					In:     []reflect.Type{reflect.TypeOf(Order{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(18))},
					Result: []any{int64(18)},
				},
				//"getSeller": {
				//	Name:   "getSeller",
				//	In:     []reflect.Type{reflect.TypeOf(Order{})},
				//	Out:    []reflect.Type{reflect.TypeOf(int64(18))},
				//	Result: []any{int64(18)},
				//},
			},
		},
		{
			name: "pointer type",
			args: args{
				val: &OrderV1{
					buyer:  18,
					seller: 100,
				},
			},
			want: map[string]*FuncInfo{
				"GetBuyer": {
					Name:   "GetBuyer",
					In:     []reflect.Type{reflect.TypeOf(&OrderV1{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(18))},
					Result: []any{int64(18)},
				},
			},
		},
		{
			name: "pointer type but input struct",
			args: args{
				val: OrderV1{
					buyer:  18,
					seller: 100,
				},
			},
			want: map[string]*FuncInfo{},
		},
		{
			name: "struct type but input ptr",
			args: args{
				val: &Order{
					buyer:  18,
					seller: 100,
				},
			},
			want: map[string]*FuncInfo{
				"GetBuyer": {
					Name:   "GetBuyer",
					In:     []reflect.Type{reflect.TypeOf(&Order{})},
					Out:    []reflect.Type{reflect.TypeOf(int64(18))},
					Result: []any{int64(18)},
				},
				//"getSeller": {
				//	Name:   "getSeller",
				//	In:     []reflect.Type{reflect.TypeOf(Order{})},
				//	Out:    []reflect.Type{reflect.TypeOf(int64(18))},
				//	Result: []any{int64(18)},
				//},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IterateFuncs(tt.args.val)
			assert.Equal(t, err, tt.wantErr)
			if err != nil {
				return
			}
			assert.Equalf(t, tt.want, got, "IterateFuncs(%v)", tt.args.val)
		})
	}
}

type Order struct {
	buyer  int64
	seller int64
}

func (o Order) GetBuyer(a int) int64 {
	return o.buyer
}

//func (o Order) getSeller() int64 {
//	return o.seller
//}

type OrderV1 struct {
	buyer  int64
	seller int64
}

func (o *OrderV1) GetBuyer() int64 {
	return o.buyer
}

type MyInterface interface {
	Abc()
}

var _ MyInterface = &abcImpl{}

type abcImpl struct {
}

func (a *abcImpl) Abc() {
	//TODO implement me
	panic("implement me")
}

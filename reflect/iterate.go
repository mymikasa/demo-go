package reflect

import (
	"errors"
	"reflect"
)

func Iterate(input any) ([]any, error) {

	val := reflect.ValueOf(input)
	typ := val.Type()
	kind := typ.Kind()

	if kind != reflect.Array && kind != reflect.Slice && kind != reflect.String {
		return nil, errors.New("非法类型")
	}

	res := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		ele := val.Index(i)
		res = append(res, ele.Interface())
	}
	return res, nil
}

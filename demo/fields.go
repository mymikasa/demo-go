package demo

import (
	"errors"
	"fmt"
	"reflect"
)

func IterateFields(val any) {
	res, err := iterateFields(val)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range res {
		fmt.Println(k, v)
	}
}

func iterateFields(val any) (map[string]any, error) {
	if val == nil {
		return nil, errors.New("not nil")
	}

	typ := reflect.TypeOf(val)
	refVal := reflect.ValueOf(val)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		refVal = refVal.Elem()
	}

	numFiled := typ.NumField()
	res := make(map[string]any, numFiled)
	for i := 0; i < numFiled; i++ {
		fdtype := typ.Field(i)
		res[fdtype.Name] = refVal.Field(i).Interface()
	}
	return res, nil
}

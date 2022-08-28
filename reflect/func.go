package reflect

import (
	"errors"
	"reflect"
)

type FuncInfo struct {
	Name string
	In   []reflect.Type
	Out  []reflect.Type

	// 反射调用得到的结果
	Result []any
}

// 输出方法信息并调用
func IterateFuncs(val any) (map[string]*FuncInfo, error) {

	if val == nil {
		return nil, errors.New("nil")
	}
	typ := reflect.TypeOf(val)

	// 只处理一级指针
	//if typ.Kind() == reflect.Ptr {
	//	typ = typ.Elem()
	//}

	if typ.Kind() != reflect.Struct && !(typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct) {
		return nil, errors.New("不支持类型")
	}

	numMethod := typ.NumMethod()
	res := make(map[string]*FuncInfo, numMethod)

	for i := 0; i < numMethod; i++ {
		method := typ.Method(i)
		mt := method.Type
		numIn := mt.NumIn()
		in := make([]reflect.Type, 0, numIn)
		for j := 0; j < numIn; j++ {
			in = append(in, mt.In(j))
		}
		numOut := mt.NumOut()
		out := make([]reflect.Type, 0, numOut)
		for j := 0; j < numOut; j++ {
			out = append(out, mt.Out(j))
		}

		callRes := method.Func.Call([]reflect.Value{
			reflect.ValueOf(val),
		})
		retValues := make([]any, 0, len(callRes))
		for _, cr := range callRes {
			retValues = append(retValues, cr.Interface())
		}
		res[method.Name] = &FuncInfo{
			Name:   method.Name,
			In:     in,
			Out:    out,
			Result: retValues,
		}
	}
	return res, nil
}

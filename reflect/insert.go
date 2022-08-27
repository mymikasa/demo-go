package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var errInvalidEntity = errors.New("invalid entity")

type InsertField struct {
	name        string
	fieldNames  []string
	fieldValues []any
	err         error
}

func NewInsertField() *InsertField {
	return &InsertField{}
}

func (i *InsertField) HandleStruct() {

}

func (i *InsertField) GenerateSql() string {
	res := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s);", i.name, strings.Join(i.fieldNames, ","),
		strings.TrimRight(strings.Repeat("?,", len(i.fieldValues)), ","))
	return res
}

func InsertStmt(entity interface{}) (string, []interface{}, error) {

	insertField := NewInsertField()
	val := reflect.ValueOf(entity)
	typ := val.Type()

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	insertField.name = "`" + typ.Name() + "`"
	fmt.Println(insertField.name)

	if val.IsValid() == false || val.NumField() == 0 {
		return "", nil, errInvalidEntity
	}
	for i := 0; i < val.NumField(); i++ {
		insertField.fieldValues = append(insertField.fieldValues, val.Field(i).Interface())
		insertField.fieldNames = append(insertField.fieldNames, "`"+typ.Field(i).Name+"`")
	}
	return insertField.GenerateSql(), insertField.fieldValues, nil

	//bd := strings.Builder{}
	//bd.Write()

	// val := reflect.ValueOf(entity)
	// typ := val.Type()
	// 检测 entity 是否符合我们的要求
	// 我们只支持有限的几种输入

	// 使用 strings.Builder 来拼接 字符串
	// bd := strings.Builder{}

	// 构造 INSERT INTO XXX，XXX 是你的表名，这里我们直接用结构体名字

	// 遍历所有的字段，构造出来的是 INSERT INTO XXX(col1, col2, col3)
	// 在这个遍历的过程中，你就可以把参数构造出来
	// 如果你打算支持组合，那么这里你要深入解析每一个组合的结构体
	// 并且层层深入进去

	// 拼接 VALUES，达成 INSERT INTO XXX(col1, col2, col3) VALUES

	// 再一次遍历所有的字段，要拼接成 INSERT INTO XXX(col1, col2, col3) VALUES(?,?,?)
	// 注意，在第一次遍历的时候我们就已经拿到了参数的值，所以这里就是简单拼接 ?,?,?

	// return bd.String(), args, nil
	panic("implement me")
}

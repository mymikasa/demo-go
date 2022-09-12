package unsafe

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

func TestNewUnsafeAccessor(t *testing.T) {
	type args struct {
		entity interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *UnsafeAccessor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUnsafeAccessor(tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUnsafeAccessor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnsafeAccessor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnsafeAccessor_Field(t *testing.T) {
	testCases := []struct {
		name    string
		entity  interface{}
		field   string
		wantVal int
		wantErr error
	}{
		{
			name:    "normal case",
			entity:  &User{Age: 18},
			field:   "Age",
			wantVal: 18,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewUnsafeAccessor(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			val, err := accessor.Field(tc.field)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestUnsafeAccessor_SetField(t *testing.T) {
	type fields struct {
		fields     map[string]FieldMeta
		entityAddr unsafe.Pointer
	}
	type args struct {
		field string
		val   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UnsafeAccessor{
				fields:     tt.fields.fields,
				entityAddr: tt.fields.entityAddr,
			}
			if err := u.SetField(tt.args.field, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("SetField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type User struct {
	Age int
}

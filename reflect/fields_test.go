package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User1 struct {
	Name string
}

func TestIterateFields(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantRes map[string]any
		wantErr error
	}{
		{
			name:    "nil",
			val:     nil,
			wantErr: errors.New("not nil"),
		},
		{
			name: "user",
			val: User1{
				Name: "Tom",
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Tom",
			},
		},
		{
			name: "ptr",
			val: &User1{
				Name: "Mikasa",
			},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Mikasa",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := iterateFields(tt.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

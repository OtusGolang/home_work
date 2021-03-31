package hw09_struct_validator //nolint:golint,stylecheck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			User{
				ID:    "d8f4590320e1343a915lb69410650a8f359d",
				Name:  "Positive",
				Age:   50,
				Email: "test@test.ru",
				Role:  "admi",
				Phones: []string{
					"89999999999",
				},
			},
			nil,
		},
		{
			App{
				Version: "12.05",
			},
			nil,
		},
		{
			Response{
				Code: 200,
				Body: "",
			},
			nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("positive case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

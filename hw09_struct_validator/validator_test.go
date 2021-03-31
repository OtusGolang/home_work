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

	UserN struct {
		ID    string `json:"id" validate:"le:36"`
		Name  string
		Age   int    `validate:"min:18|max:50"`
		Email string `validate:"regexp:regexp:^\\w+@\\w+\\.\\w+$"`
	}

	AppN struct {
		Version string `validate:"len:aa"`
	}
)

func TestValidate(t *testing.T) {
	testsP := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			User{
				ID:    "d8f4590320e1343a915lb69410650a8f359d",
				Name:  "Positive",
				Age:   50,
				Email: "test@test.ru",
				Role:  "admin",
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

	testsN := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			UserN{
				ID:    "3131313131",
				Name:  "Negative",
				Age:   100,
				Email: "test2test.ru",
			},
			ValidationErrors{
				{"ID", ErrorUnknownRule},
				{"Age", ErrorValue},
				{"Email", ErrorValue},
			},
		},
		{
			AppN{
				Version: "11111",
			},
			ValidationErrors{
				ValidationError{
					"Version",
					ErrorRuleValueIsNotNumber,
				},
			},
		},
	}

	for i, tt := range testsP {
		t.Run(fmt.Sprintf("positive case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
		})
	}

	for i, tt := range testsN {
		t.Run(fmt.Sprintf("negative case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

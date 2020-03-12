//+build generation

package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Validated interface {
	Validate() ([]ValidationError, error)
}

func TestUserValidation(t *testing.T) {
	var v interface{} = User{}
	_, ok := v.(Validated)
	require.True(t, ok)

	t.Run("ID length", func(t *testing.T) {
		errs, _ := User{ID: "441e0ae8644611eab8a0632c74ca9988"}.Validate()
		require.Len(t, errs, 0)

		errs, _ = User{ID: "123"}.Validate()
		require.Equal(t, len(errs), 1)
		require.Equal(t, errs[0].Field, "ID")
		require.NotNil(t, errs[0].Err)
	})

	t.Run("email regexp", func(t *testing.T) {
		errs, _ := User{Email: "owl@@otus.ru"}.Validate()
		require.Len(t, errs, 0)

		errs, _ = User{Email: "isnotvalid@@email"}.Validate()
		require.Equal(t, len(errs), 1)
		require.Equal(t, errs[0].Field, "Email")
		require.NotNil(t, errs[0].Err)
	})

	t.Run("age borders", func(t *testing.T) {
		errs, _ := User{Age: 17}.Validate()
		require.NotEqual(t, len(errs), 0)

		for _, a := range []int{18, 34, 50} {
			errs, _ := User{Age: a}.Validate()
			require.Len(t, errs, 0)
		}

		errs, _ = User{Age: 51}.Validate()
		require.Equal(t, len(errs), 1)
		require.Equal(t, errs[0].Field, "Age")
		require.NotNil(t, errs[0].Err)
	})

	t.Run("addresses slice", func(t *testing.T) {
		// Write me :)
	})
}

func TestAppValidation(t *testing.T) {
	var v interface{} = App{}
	_, ok := v.(Validated)
	require.True(t, ok)

	t.Run("version length", func(t *testing.T) {
		errs, _ = App{"0.1"}.Validate()
		require.Equal(t, len(errs), 1)
		require.Equal(t, errs[0].Field, "Version")
		require.NotNil(t, errs[0].Err)
	})
}

func TestTokenValidation(t *testing.T) {
	var v interface{} = Token{}
	_, ok := v.(Validated)
	require.False(t, ok)
}

func TestResponseValidation(t *testing.T) {
	var v interface{} = Response{}
	_, ok := v.(Validated)
	require.True(t, ok)

	t.Run("code set", func(t *testing.T) {
		for _, c := range []int{200, 404, 500} {
			errs, _ := Response{Code: c}.Validate()
			require.Len(t, errs, 0)
		}

		errs, _ := Response{Code: 133}.Validate()
		require.Equal(t, len(errs), 1)
		require.Equal(t, errs[0].Field, "Code")
		require.NotNil(t, errs[0].Err)
	})
}

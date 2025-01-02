package configuration

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecretsAreNotExposedAsStrings(t *testing.T) {
	secret := NewSecret[string]("password")
	require.Equal(t, SecretOverwrite, secret.String())
	require.Equal(t, SecretOverwrite, fmt.Sprint(secret))
	require.Equal(t, "password", secret.GetSecretValue())
}

func TestSecretDecodeHook(t *testing.T) {
	t.Run(
		"strings can be convert to Secret[string]",
		func(t *testing.T) {
			input := "password"
			expected := NewSecret[string](input)

			result, err := SecretDecodeHook(reflect.TypeOf(input), reflect.TypeOf(Secret[string]{}), input)
			require.NoError(t, err)
			require.Equal(t, expected, result)
		})

	t.Run(
		"non-string types return the original data as the same type",
		func(t *testing.T) {
			input := 123 // invalid type to be converted to a Secret[string].

			result, err := SecretDecodeHook(reflect.TypeOf(input), reflect.TypeOf(Secret[string]{}), input)
			require.NoError(t, err)
			require.Equal(t, input, result)
		})
}

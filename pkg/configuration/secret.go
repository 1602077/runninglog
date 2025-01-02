package configuration

import (
	"encoding/json"
	"reflect"
)

// SecretOverwrite is the value used to obfuscate secret credentials with inside
// of logging and JSON marshalling etc.
const SecretOverwrite = "***"

// Secret implements the stringer and JSON marshalling interfaces to prevent the
// leaking of secret values in cloud logs.
type Secret[T any] struct {
	value T
}

// GetSecretValue retrieves the underlying private secret.
func (s Secret[T]) GetSecretValue() T { return s.value }

// String overwrites the Stringer interface to obfuscate secrets in logs.
func (s Secret[T]) String() string { return SecretOverwrite }

// MarshalJSON is used to obfuscate secrets when marshalling to JSON.
func (s Secret[T]) MarshalJSON() ([]byte, error) { return json.Marshal(s.String()) }

func NewSecret[T any](value T) Secret[T] { return Secret[T]{value: value} }

// SecretDecodeHook is used by viper to decode a string to a Secret.
var SecretDecodeHook = func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f.Kind() != reflect.String || t != reflect.TypeOf(Secret[string]{}) {
		return data, nil
	}

	return NewSecret[string](data.(string)), nil
}

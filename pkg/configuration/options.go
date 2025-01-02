package configuration

import (
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
)

// Options configures how and where configuration is loaded from.
type Options struct {
	// BaseFile is the path to the default configuration file.
	BaseFile string

	// OverrideFiles are additive to the configuration in the BaseFile and
	// overwrite any existing config where it is defined in multiple files.
	OverrideFiles []string

	//EnvironmentOverridePrefix is the prefix to place on environment variables
	// used to set config.
	//
	// This prevents application specific constants from overwriting any in the
	//runtime container OS.
	EnvironmentOverridePrefix string

	// Validator validates configurations using tags specified on the application
	// config struct.
	Validator *validator.Validate

	// DecodeHooks defines the validation hooks for validating all types.
	// CustomHooks for custom types will need to be registered here.
	DecodeHooks []mapstructure.DecodeHookFunc
}

var DefaultDecodeHooks []mapstructure.DecodeHookFunc = []mapstructure.DecodeHookFunc{
	SecretDecodeHook,
	mapstructure.StringToTimeDurationHookFunc(),
	mapstructure.StringToSliceHookFunc(","),
}

func WithBaseFile(baseFile string) func(*Options) {
	return func(o *Options) {
		o.BaseFile = baseFile
	}
}

func WithOverrideFiles(overrideFiles ...string) func(*Options) {
	return func(o *Options) {
		o.OverrideFiles = append(o.OverrideFiles, overrideFiles...)
	}
}

func WithEnvironmentOverridePrefix(prefix string) func(*Options) {
	return func(o *Options) {
		o.EnvironmentOverridePrefix = prefix
	}
}

func WithValidator(validator *validator.Validate) func(*Options) {
	return func(o *Options) {
		o.Validator = validator
	}
}

func WithDecodeHooks(hooks ...mapstructure.DecodeHookFunc) func(*Options) {
	return func(o *Options) {
		o.DecodeHooks = append(o.DecodeHooks, hooks)
	}
}

package configuration

import (
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Load reads in application configuration as determined by functional options.
func Load[T any](options ...func(*Options)) (*T, error) {
	opts := &Options{ // set sensible defaults.
		BaseFile:                  "",
		OverrideFiles:             []string{},
		EnvironmentOverridePrefix: "APPLICATION",
		Validator:                 nil, // default: no validation.
		DecodeHooks:               DefaultDecodeHooks,
	}
	for _, o := range options {
		o(opts)
	}

	// load base config.
	vv := viper.New()
	vv.SetConfigFile(opts.BaseFile)
	if err := vv.ReadInConfig(); err != nil {
		return nil, err
	}

	// load override config(s).
	for _, override := range opts.OverrideFiles {
		vv.SetConfigFile(override)
		if err := vv.MergeInConfig(); err != nil {
			return nil, err
		}
	}

	// load environment variables.
	vv.SetEnvPrefix(opts.EnvironmentOverridePrefix)
	vv.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vv.AutomaticEnv()

	// marshal config into struct.
	conf := new(T)
	err := vv.Unmarshal(
		conf,
		viper.DecodeHook(
			mapstructure.ComposeDecodeHookFunc(
				opts.DecodeHooks...,
			),
		),
	)
	if err != nil {
		return nil, err
	}

	// validate configuration.
	if opts.Validator != nil {
		if err := opts.Validator.Struct(conf); err != nil {
			return nil, err
		}
	}

	return conf, nil
}

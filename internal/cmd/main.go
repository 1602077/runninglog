package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"

	internal "github.com/1602077/runninglog/internal"
	"github.com/1602077/runninglog/pkg/configuration"
	"github.com/1602077/runninglog/pkg/logger"
)

const ConfigBoostrapEnvVarPrefix = "CONFIG"

func main() {
	config, err := loadConfig()
	if err != nil {
		os.Exit(1)
	}

	logger.NewSlogLogger(&config.Logger)
	logger.Info("starting application")
	logger.Debug("application config loaded", "config", config)

	// TODO: Run application from here.
}

// loadConfig bootstraps the global application configuration.
//
// It uses environment variables prefixed by ConfigBoostrapEnvVarPrefix to
// specific which files and where to pull the config from. As the global logger
// requires config to be provided before it can be used this reverts to using
// slog for logging.
func loadConfig() (*internal.Application, error) {
	validate := validator.New()

	// load details of where config is stored from environment variables.
	basefile := os.Getenv(ConfigBoostrapEnvVarPrefix + "_BASE_FILE")
	overridefilesStr := os.Getenv(ConfigBoostrapEnvVarPrefix + "_OVERRIDE_FILES")
	overridefiles := strings.Split(overridefilesStr, ",")
	envprefix := os.Getenv(ConfigBoostrapEnvVarPrefix + "_ENVIRONMENT_PREFIX")

	config, err := configuration.Load[internal.Application](
		configuration.WithBaseFile(basefile),
		configuration.WithOverrideFiles(overridefiles...),
		configuration.WithEnvironmentOverridePrefix(envprefix),
		configuration.WithValidator(validate),
	)
	if err != nil {
		slog.Error(
			"failed to load application config",
			"error", err,
		)
		return nil, err
	}

	return config, err
}

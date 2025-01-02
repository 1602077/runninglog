package logger

// Config holds the global configuration for logging.
type Config struct {
	Level             Level `mapstructure:"level" validate:"oneof=debug info warn error"`
	StructuredLogging bool  `mapstructure:"structured_logging"`
}

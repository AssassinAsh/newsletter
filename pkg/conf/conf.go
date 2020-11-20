package conf

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// ErrMissingEnvironmentStage missing stage configuration
	ErrMissingEnvironmentStage = errors.New("Missing Stage ENV Variable")

	// ErrMissingEnvironmentBranch missing branch configuration
	ErrMissingEnvironmentBranch = errors.New("Missing Branch ENV Variable")
)

//Config for the environment
type Config struct {
	Debug     bool   `envconfig:"DEBUG"`
	Addr      string `envconfig:"ADDR" default:"8080"`
	Stage     string `envconfig:"STAGE" default:"dev"`
	Branch    string `envconfig:"BRANCH" default:"master"`
	DbSecrets string `envconfig:"DB_SECRET"`
}

//DBSecrets for the database
type DBSecrets struct {
	Secret string `json:"secret,omitempty"`
	DBName string `json:"dbname,omitempty"`
	URL    string `json:"url,omitempty"`
}

//Setting up the Logger for whole project
func (cfg *Config) logging() error {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if cfg.Stage == "local" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return nil
}

//Validating the environment variables
func (cfg *Config) validate() error {
	if cfg.Stage == "" {
		return ErrMissingEnvironmentStage
	}
	if cfg.Branch == "" {
		return ErrMissingEnvironmentBranch
	}

	return nil
}

// NewDefaultConfig reads configuration from environment variables and validates it
func NewDefaultConfig() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, errors.Wrap(err, "Failed to parse Environment config.")
	}

	if err := cfg.validate(); err != nil {
		return nil, errors.Wrap(err, "Failed to validate config.")
	}

	if err := cfg.logging(); err != nil {
		return nil, errors.Wrap(err, "Failed to enable logging based on config.")
	}

	log.Info().Str("stage", cfg.Stage).Bool("debug", cfg.Debug).Msg("Logging Configured")
	log.Info().Str("stage", cfg.Stage).Str("branch", cfg.Branch).Msg("Configurations Loaded")

	return cfg, nil
}

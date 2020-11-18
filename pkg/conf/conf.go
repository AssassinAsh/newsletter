package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var (
	// ErrMissingEnvironmentStage missing stage configuration
	ErrMissingEnvironmentStage = errors.New("Missing Stage ENV Variable")

	// ErrMissingEnvironmentBranch missing branch configuration
	ErrMissingEnvironmentBranch = errors.New("Missing Branch ENV Variable")
)

//Config for the environment
type Config struct {
	Addr      string `envconfig:"ADDR" default:":8080"`
	Stage     string `envconfig:"STAGE" default:"dev"`
	Branch    string `envconfig:"BRANCH"`
	DbSecrets string `envconfig:"DB_SECRET"`
}

//DBSecrets for the database
type DBSecrets struct {
	Secret string `json:"secret,omitempty"`
	DBName string `json:"dbname,omitempty"`
	URL    string `json:"url,omitempty"`
}

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

	err := envconfig.Process("", cfg)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse Environment config.")
	}

	err = cfg.validate()

	if err != nil {
		return nil, errors.Wrap(err, "Failed to validate config.")
	}

	return cfg, nil
}

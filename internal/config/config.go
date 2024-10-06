package config

type Config struct {
	// The port to listen on
	Port int `env:"PORT" envDefault:"8080"`
	// The environment to run in
	Environment string `env:"ENVIRONMENT" envDefault:"development"`

	// The database driver to use
	DatabaseDriver string `env:"DATABASE_DRIVER" envDefault:"postgres"`

	// The database host
	Host string `env:"HOST" envDefault:"localhost"`

	// The database port
	DBPort int `env:"DB_PORT" envDefault:"5432"`

	// The database user
	User string `env:"DB_USER" envDefault:"postgres"`

	// The database password
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`

	// The database name
	DBName string `env:"DB_NAME" envDefault:"postgres"`

	// The Connection Options

	// The maximum number of connections in the idle connection pool
	MaxIdleConns int `env:"DB_MAX_IDLE_CONNS" envDefault:"5"`

	// The maximum number of open connections to the database
	MaxOpenConns int `env:"DB_MAX_OPEN_CONNS" envDefault:"5"`

	// The maximum amount of time a connection may be reused
	ConnMaxLifetime int `env:"DB_CONN_MAX_LIFETIME" envDefault:"300"`
}

// NewConfig returns a new Config struct
func NewConfig() *Config {
	return &Config{}
}

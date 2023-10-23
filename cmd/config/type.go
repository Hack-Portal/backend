package config

var Config *config

type config struct {
	Server struct {
		Addr            string `env:"SERVER_ADDR" envDefault:"localhost:8080"`
		ShutdownTimeout int    `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"10"`
		ContextTimeout  int    `env:"SERVER_CONTEXT_TIMEOUT" envDefault:"5"`
	}

	Postgres struct {
		Host     string `env:"PSQL_HOST" envDefault:"localhost"`
		Port     int    `env:"PSQL_PORT" envDefault:"5432"`
		User     string `env:"PSQL_USER" envDefault:"postgres"`
		Password string `env:"PSQL_PASSWORD" envDefault:"postgres"`
		DBName   string `env:"PSQL_DBNAME" envDefault:"postgres"`
		SSLMode  string `env:"PSQL_SSLMODE" envDefault:"disable"`

		ConnectTimeout  int `env:"PSQL_CONNECT_TIMEOUT" envDefault:"10"`
		ConnectWaitTime int `env:"PSQL_CONNECT_WAIT_TIME" envDefault:"10"`
		ConnectAttempts int `env:"PSQL_CONNECT_ATTEMPTS" envDefault:"3"`
	}

	NewRelic struct {
		LicenceKey string `env:"NEWRELIC_LICENCE_KEY" envDefault:""`
		AppName    string `env:"NEWRELIC_APP_NAME" envDefault:""`
	}

	Firebase struct {
		StorageBucket string `env:"FIREBASE_STORAGE_BUCKET" envDefault:""`
	}
}

package config

var Config *config

type config struct {
	Server struct {
		Addr            string `env:"SERVER_ADDR" envDefault:"localhost:8080"`
		ShutdownTimeout int    `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"10"`
		ContextTimeout  int    `env:"SERVER_CONTEXT_TIMEOUT" envDefault:"10"`
	}

	Database struct {
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		Port     int    `env:"DB_PORT" envDefault:"5432"`
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASSWORD" envDefault:"postgres"`
		DBName   string `env:"DB_NAME" envDefault:"hackportal"`
		SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
		TimeZone string `env:"DB_TIMEZONE" envDefault:"Asia/Tokyo"`

		ConnectTimeout  int  `env:"DB_CONNECT_TIMEOUT" envDefault:"10"`
		ConnectWaitTime int  `env:"DB_CONNECT_WAIT_TIME" envDefault:"10"`
		ConnectAttempts int  `env:"DB_CONNECT_ATTEMPTS" envDefault:"3"`
		ConnectBlocks   bool `env:"DB_CONNECT_BLOCKS" envDefault:"false"`
		CloseTimeout    int  `env:"DB_CLOSE_TIMEOUT" envDefault:"10"`
	}
}

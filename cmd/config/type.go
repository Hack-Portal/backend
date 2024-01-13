package config

import "time"

var Config *config

type config struct {
	Server struct {
		Addr              string `env:"SERVER_ADDR" envDefault:"8080"`
		ShutdownTimeout   int    `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"10"`
		ContextTimeout    int    `env:"SERVER_CONTEXT_TIMEOUT" envDefault:"10"`
		Version           string `env:"SERVER_VERSION" envDefault:"0.1.0"`
		AdminInitPassword string `env:"SERVER_ADMIN_INIT_PASSWORD" envDefault:"ptZmQNRfr8HSBrhGgzYvDSFRfaVktAbfh25KIA2hEywPPIx2hB"`
	}

	Database struct {
		DB       string `env:"DB_DRIVER" envDefault:"postgres"`
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		Port     int    `env:"DB_PORT" envDefault:"5432"`
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASSWORD" envDefault:"postgres"`
		DBName   string `env:"DB_NAME" envDefault:"hack_portal"`
		SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
		TimeZone string `env:"DB_TIMEZONE" envDefault:"Asia/Tokyo"`
	}

	Redis struct {
		Host     string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
		Port     int    `env:"REDIS_PORT" envDefault:"6379"`
		Password string `env:"REDIS_PASSWORD" envDefault:""`
		DB       int    `env:"REDIS_DB" envDefault:"0"`

		ConnectTimeout  int `env:"REDIS_CONNECT_TIMEOUT" envDefault:"10"`
		ConnectWaitTime int `env:"REDIS_CONNECT_WAIT_TIME" envDefault:"10"`
		ConnectAttempts int `env:"REDIS_CONNECT_ATTEMPTS" envDefault:"3"`
	}

	Buckets struct {
		EndPoint        string        `env:"BUCKETS_ENDPOINT" envDefault:""`
		AccountID       string        `env:"BUCKETS_ACCOUNT_ID" envDefault:""`
		AccessKeyID     string        `env:"BUCKETS_ACCESS_KEY_ID" envDefault:""`
		AccessKeySecret string        `env:"BUCKETS_ACCESS_KEY_SECRET" envDefault:""`
		Bucket          string        `env:"CLOUDFLARE_BUCKET" envDefault:""`
		Expired         time.Duration `env:"BUCKETS_EXPIRED" envDefault:"30m"`
	}

	NewRelic struct {
		AppName    string `env:"NEWRELIC_APPLICATION_NAME" envDefault:"hack-portal"`
		LicenseKey string `env:"NEWRELIC_LICENSE_KEY" envDefault:""`
	}

	Discord struct {
		Secret string `env:"DISCORD_SECRET" envDefault:""`
	}
}

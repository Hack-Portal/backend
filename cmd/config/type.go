package config

import "time"

var Config *config

type config struct {
	Server struct {
		Addr              string        `env:"SERVER_ADDR" envDefault:"8080"`
		ShutdownTimeout   time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"10s"`
		Version           string        `env:"SERVER_VERSION" envDefault:"0.1.0"`
		AdminInitPassword string        `env:"SERVER_ADMIN_INIT_PASSWORD" envDefault:"ptZmQNRfr8HSBrhGgzYvDSFRfaVktAbfh25KIA2hEywPPIx2hB"`
	}

	Database struct {
		Driver          string        `env:"POSTGRES_DRIVER" envDefault:"postgres"`
		Host            string        `env:"POSTGRES_HOST" envDefault:"localhost"`
		Port            int           `env:"POSTGRES_PORT" envDefault:"5432"`
		User            string        `env:"POSTGRES_USER" envDefault:"postgres"`
		Password        string        `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
		DBName          string        `env:"POSTGRES_DB_NAME" envDefault:"hack_portal"`
		SSLMode         string        `env:"POSTGRES_SSLMODE" envDefault:"disable"`
		TimeZone        string        `env:"POSTGRES_TIMEZONE" envDefault:"Asia/Tokyo"`
		ConnectTimeout  time.Duration `env:"POSTGRES_CONNECT_TIMEOUT" envDefault:"10s"`
		ConnectWaitTime time.Duration `env:"POSTGRES_CONNECT_WAIT_TIME" envDefault:"10s"`
		ConnectAttempts int           `env:"POSTGRES_CONNECT_ATTEMPTS" envDefault:"3"`
	}

	Redis struct {
		Host            string        `env:"REDIS_HOST" envDefault:"redis"`
		Port            int           `env:"REDIS_PORT" envDefault:"6379"`
		Password        string        `env:"REDIS_PASSWORD" envDefault:""`
		DBName          int           `env:"REDIS_DB_NAME" envDefault:"0"`
		ConnectTimeout  time.Duration `env:"REDIS_CONNECT_TIMEOUT" envDefault:"10s"`
		ConnectWaitTime time.Duration `env:"REDIS_CONNECT_WAIT_TIME" envDefault:"10s"`
		ConnectAttempts int           `env:"REDIS_CONNECT_ATTEMPTS" envDefault:"3"`
	}

	Buckets struct {
		EndPoint        string `env:"BUCKETS_ENDPOINT" envDefault:""`
		AccountID       string `env:"BUCKETS_ACCOUNT_ID" envDefault:""`
		AccessKeyID     string `env:"BUCKETS_ACCESS_KEY_ID" envDefault:""`
		AccessKeySecret string `env:"BUCKETS_ACCESS_KEY_SECRET" envDefault:""`
		Bucket          string `env:"CLOUDFLARE_BUCKET" envDefault:""`
		Expired         int    `env:"BUCKETS_EXPIRED" envDefault:"30"`
	}

	NewRelic struct {
		AppName    string `env:"NEWRELIC_APPLICATION_NAME" envDefault:"hack-portal"`
		LicenseKey string `env:"NEWRELIC_LICENSE_KEY" envDefault:""`
	}
}

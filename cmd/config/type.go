package config

var Config *config

type config struct {
	Server struct {
		Addr            string `env:"SERVER_ADDR" envDefault:"8080"`
		ShutdownTimeout int    `env:"SERVER_SHUTDOWN_TIMEOUT" envDefault:"10"`
		ContextTimeout  int    `env:"SERVER_CONTEXT_TIMEOUT" envDefault:"10"`

		DefaultHackathonImage string `env:"SERVER_DEFAULT_HACKATHON_IMAGE" envDefault:"https://e4fa9209c88aac97b94a1000743846ec.r2.cloudflarestorage.com/hack-portal/hackathon/default.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=b552af962ee8bf7928c3bc83b047d775%2F20231222%2Fauto%2Fs3%2Faws4_request&X-Amz-Date=20231222T063445Z&X-Amz-Expires=900&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=de7e1b3782b4585ba85c521191742a5ab44033a3f308d70d06d60cc6640a1a2a"`
	}

	Database struct {
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		Port     int    `env:"DB_PORT" envDefault:"5432"`
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASSWORD" envDefault:"postgres"`
		DBName   string `env:"DB_NAME" envDefault:"hack_portal"`
		SSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
		TimeZone string `env:"DB_TIMEZONE" envDefault:"Asia/Tokyo"`

		ConnectTimeout  int  `env:"DB_CONNECT_TIMEOUT" envDefault:"10"`
		ConnectWaitTime int  `env:"DB_CONNECT_WAIT_TIME" envDefault:"10"`
		ConnectAttempts int  `env:"DB_CONNECT_ATTEMPTS" envDefault:"3"`
		ConnectBlocks   bool `env:"DB_CONNECT_BLOCKS" envDefault:"false"`
		CloseTimeout    int  `env:"DB_CLOSE_TIMEOUT" envDefault:"10"`
	}

	Buckets struct {
		EndPoint        string `env:"BUCKETS_ENDPOINT" envDefault:""`
		AccountID       string `env:"BUCKETS_ACCOUNT_ID" envDefault:""`
		AccessKeyId     string `env:"BUCKETS_ACCESS_KEY_ID" envDefault:""`
		AccessKeySecret string `env:"BUCKETS_ACCESS_KEY_SECRET" envDefault:""`
		Bucket          string `env:"CLOUDFLARE_BUCKET" envDefault:""`
		Expired         int    `env:"BUCKETS_EXPIRED" envDefault:"720"`
	}
}

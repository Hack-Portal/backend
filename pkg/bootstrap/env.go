package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

// envファイルに記載されている値のモデル
type Env struct {
	// DB情報
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`

	ServerPort     string `mapstructure:"SERVER_PORT"`
	ContextTimeout string `mapstructure:"CONTEXT_TIMEOUT"`
	BasePath       string `mapstructure:"BASE_PATH"`

	TokenDuration     string `mapstructure:"TOKEN_DURATION"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`

	NewRelicAppName string `mapstructure:"NEW_RELIC_APP_NAME"`
	NewRelicLicense string `mapstructure:"NEW_RELIC_LICENSE"`
}

// app.envファイルを読み込む
func LoadEnvConfig(path string) (config Env) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return
}

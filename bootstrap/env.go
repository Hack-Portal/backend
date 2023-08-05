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

	ServerPort string `mapstructure:"SERVER_PORT"`
	BasePath   string `mapstructure:"BASE_PATH"`
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

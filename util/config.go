package util

import (
	"github.com/spf13/viper"
)

// envファイルに記載されている値のモデル
type EnvConfig struct {
	// DB情報
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSouse  string `mapstructure:"DB_SOURSE"`

	ServerPort string `mapstructure:"SERVER_PORT"`

	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClinetSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
}

// app.envファイルを読み込む
func LoadEnvConfig(path string) (config EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

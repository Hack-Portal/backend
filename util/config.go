package util

import "github.com/spf13/viper"

// envファイルに記載されている値のモデル
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSouse  string `mapstructure:"DB_SOURSE"`
}

// app.envファイルを読み込む
func LoadConfig(path string) (config Config, err error) {
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

package util

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

// envファイルに記載されている値のモデル
type EnvConfig struct {
	// DB情報
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSouse  string `mapstructure:"DB_SOURSE"`
	// google Auth情報
	GClientID                string `json:"client_id" mapstructure:"CLIENT_ID"`
	GProjectID               string `json:"project_id" mapstructure:"PROJECT_ID"`
	GAuthUri                 string `json:"auth_uri" mapstructure:"AUTH_URI"`
	GTokenUri                string `json:"token_uri" mapstructure:"TOKEN_URI"`
	GAuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url" mapstructure:"AUTH_PROVIDER_X509_CERT_URL"`
	GClientSecret            string `json:"client_secret" mapstructure:"CLIENT_SECRET"`

	// test用
	credentials string `mapstructure:"CREDENTIALS_JSON"`
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
	fmt.Println(config.credentials)
	return
}

type credentialsJSON struct {
	Client_id                   string `json:"client_id"`
	Project_id                  string `json:"project_id"`
	Auth_uri                    string `json:"auth_uri"`
	Token_uri                   string `json:"token_uri"`
	Auth_provider_x509_cert_url string `json:"auth_provider_x509_cert_url"`
	Client_secret               string `json:"client_secret"`
}

func ConfigToJson(c *EnvConfig) ([]byte, error) {
	credentialsJSON := &credentialsJSON{
		Client_id:                   c.GClientID,
		Project_id:                  c.GProjectID,
		Auth_uri:                    c.GAuthUri,
		Token_uri:                   c.GTokenUri,
		Auth_provider_x509_cert_url: c.GAuthProviderX509CertUrl,
		Client_secret:               c.GClientSecret,
	}
	json, err := json.Marshal(credentialsJSON)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf(`{"web":%s}`, json)), nil
}

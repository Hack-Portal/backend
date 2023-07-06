package token

import (
	"fmt"

	"github.com/hackhack-Geek-v6/backend/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func NewGoogleAuthConf(c *util.EnvConfig) (*oauth2.Config, error) {
	// 実際にはSecretManagerなどに保存して、そこから取得する.
	credentialsJSON, err := util.ConfigToJson(c)
	if err != nil {
		return nil, fmt.Errorf("cannot load config file:%s", err)
	}

	// 第2引数に認証を求めるスコープを設定します.
	// 今回はスプレッドシートのリード権限スコープを指定.
	config, err := google.ConfigFromJSON(credentialsJSON, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %s", err)
	}

	return config, nil
}

func Auth() error {
	config, err := util.LoadEnvConfig("../")
	if err != nil {
		return err
	}

	conf, err := NewGoogleAuthConf(&config)
	if err != nil {
		return err
	}

	state := `CSRF攻撃を防ぐためにstateパラメータをつける.コールバック後のトークン取得時に検証する.`

	// stateをsessionなどに保存.

	// リダイレクトURL作成.
	redirectURL := conf.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	// redirectURLをクライアントに返す.
	fmt.Println(redirectURL)
	return nil
}

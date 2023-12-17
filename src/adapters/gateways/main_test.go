package gateways

import (
	"os"
	"testing"
)

func setup() {
	// ENVを設定する

	// テスト用のDBを作成する

	// テスト用DBにテストデータを投入する

}

func teardown() {
	// CreanUpする

	// テスト用DBを削除する

}

func TestMain(m *testing.M) {
	setup()
	defer teardown()

	os.Exit(m.Run())
}

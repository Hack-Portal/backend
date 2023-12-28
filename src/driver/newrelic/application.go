package newrelic

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Setup() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.Config.NewRelic.AppName),
		newrelic.ConfigLicense(config.Config.NewRelic.LicenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDebugLogger(os.Stdout),

		func(cfg *newrelic.Config) {
			// 無視するステータスコードを設定
			cfg.ErrorCollector.IgnoreStatusCodes = []int{
				http.StatusNotFound,
			}
		},
	)
	if err != nil {
		return nil, fmt.Errorf("newrelic Setup Error: %v\n", err)
	}

	if err := app.WaitForConnection(5 * time.Second); err != nil {
		return nil, fmt.Errorf("newrelic Setup Error: %v\n", err)
	}

	return app, nil
}

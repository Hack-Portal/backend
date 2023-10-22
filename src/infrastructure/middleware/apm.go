package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (m *middleware) Apm() gin.HandlerFunc {
	return nrgin.Middleware(m.newApm())
}

func (m *middleware) newApm() *newrelic.Application {
	apm, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.Config.NewRelic.AppName),
		newrelic.ConfigLicense(config.Config.NewRelic.LicenceKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		m.l.Panic(err)
	}

	return apm
}

package apm

import (
	"log"

	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

func NewApm(env *bootstrap.Env)( *newrelic.Application ){
	apm, err := newrelic.NewApplication(
		newrelic.ConfigAppName(env.NewRelicAppName),
		newrelic.ConfigLicense(env.NewRelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
		func(cofnig *newrelic.Config) {
			logrus.SetLevel(logrus.DebugLevel)
			cofnig.Logger = nrlogrus.StandardLogger()
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return apm
}

package apm

import (
	"log"

	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func ApmSetup(env *bootstrap.Env) *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(env.NewRelicAppName),
		newrelic.ConfigLicense(env.NewRelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatalln("cannnot connect newrelic")
	}
	return app
}

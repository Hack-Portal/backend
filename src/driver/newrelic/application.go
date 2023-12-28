package newrelic

import (
	"log"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func Setup() *newrelic.Application {
	config := newrelic.ConfigFromEnvironment()
	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Println("NewRelic Setup Error: ", err)
	}

	return app
}

package main

import (
	"github.com/cucumber/godog"
	"github.com/mihard/wkda-godog-demo/tests/contexts"
	"os"
)

func InitializeScenario(s *godog.ScenarioContext) {
	wm := contexts.InitWireMockFeature(os.Getenv("WIREMOCK_HOST"), s)
	_ = contexts.InitHTTPFeature(os.Getenv("SERVICE_HOST"), s)
	_ = contexts.InitCarLead(wm, s)
}

package datadog

import (
	"log"
	"os"
	"strings"

	"github.com/zorkian/go-datadog-api"
)

type StageAlert string

const (
	Prod    StageAlert = "prod"
	Preprod StageAlert = "preprod"
	Dev     StageAlert = "dev"
	Test    StageAlert = "test"
	Int     StageAlert = "int"
)

const (
	ddApiKey = "DATADOG_API_KEY"
	ddAppKey = "DATADOG_APP_KEY"
)

func getDatadogClient() *datadog.Client {
	return datadog.NewClient(os.Getenv(ddApiKey), os.Getenv(ddAppKey))
}

func getAllAlertMonitors() ([]datadog.Monitor, error) {
	monitors, err := getDatadogClient().GetMonitors()
	if err != nil {
		return nil, err
	}

	alertedMonitors := []datadog.Monitor{}

	for _, monitor := range monitors {
		if monitor.GetOverallState() == "Alert" {
			alertedMonitors = append(alertedMonitors, monitor)
		}
	}

	return alertedMonitors, nil
}

func GetAllAlertMonitorsByType(stage StageAlert) ([]datadog.Monitor, error) {
	alertedMonitors, err := getAllAlertMonitors()

	if err != nil {
		log.Panic(err)
		return nil, err
	}

	byTypeMonitors := []datadog.Monitor{}

	for _, monitor := range alertedMonitors {
		if strings.Contains(monitor.GetMessage(), string(stage)) {
			byTypeMonitors = append(byTypeMonitors, monitor)
		}
	}

	return byTypeMonitors, nil
}

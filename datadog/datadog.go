package datadog

import (
	"github.com/zorkian/go-datadog-api"
	"os"
	"fmt"
)


type StageAlert string

const (
	Prod StageAlert = "prod"
	Preprod StageAlert = "preprod"
	
)

const (
	ddApiKey = "DATADOG_API_KEY"
	ddAppKey = "DATADOG_APP_KEY"
)


func getDatadogClient() *datadog.Client {
  return datadog.NewClient(os.Getenv(ddApiKey), os.Getenv(ddAppKey))
}


func getAllAlertMonitors() ([]datadog.Monitor, error)  {
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


func GetAllAlertMonitorsByType() {

	getAllAlertMonitors()
}



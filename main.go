package main

import (
	"fmt"
	"log"

	"github.com/tbeaugrand/goDashing/datadog"
)

func main() {
	alertedMonitors, err := datadog.GetAllAlertMonitorsByType(datadog.Int)

	if err != nil {
		log.Panic(err)
		return
	}

	for _, x := range alertedMonitors {

		fmt.Println(x.GetName())
	}
}

package main

import (
	"fmt"
	"github.com/tbeaugrand/goDashing/datadog"
	"time"
)

func main() {
	for {
		widget, _ := datadog.GetAllAlertMonitorsByType(datadog.Int)
		show(widget)

		time.Sleep(5*time.Second)

		widget, _ = datadog.GetAllAlertMonitorsByType(datadog.Dev)
		show(widget)
		time.Sleep(5*time.Second)

		widget, _ = datadog.GetAllAlertMonitorsByType(datadog.Preprod)
		show(widget)

		time.Sleep(5*time.Second)
		widget, _ = datadog.GetAllAlertMonitorsByType(datadog.Prod)
		show(widget)

		time.Sleep(5*time.Second)

	}
}


func show(widget datadog.MonitorWidget) {
	fmt.Sprintf("Nb alert monitors = %d", len(widget.MonitorNames))
	for _, monitor := range widget.MonitorNames {
		fmt.Println(monitor)
	}
	fmt.Println("====================")
}




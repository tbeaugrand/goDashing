package main

import (
	"github.com/tbeaugrand/goDashing/datadog"
	"fmt"
)


func main() {
	alertedMonitors, err := datadog.GetAllAlertMonitors()

	if err != nil {
		fmt.Println("")
	}

	for _, x := range alertedMonitors {

		fmt.Println(x.GetName())
	}
}

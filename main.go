package main

import (
	"github.com/tbeaugrand/goDashing/datadog"
	"fmt"
)


func main() {
	alertedMonitors, err := datadog.GetALertMonitors()

	if err != nil {
		fmt.Println("")
	}

	for _, x := range alertedMonitors {

		fmt.Println(x)
	}
}

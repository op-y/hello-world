package main

import (
	"os"

	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"myscheduler/plugin"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(plugin.Name, plugin.New),
	)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

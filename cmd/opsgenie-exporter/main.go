package main

import (
	"opsgenie-exporter/internal/domain/prometheusDomain"
)

func main() {

	prometheusDomain.PushMetrics()

}

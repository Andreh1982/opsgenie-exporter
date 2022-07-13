package main

import (
	"opsgenie-exporter/internal/metrics"
)

func main() {

	metrics.PushMetrics()

}

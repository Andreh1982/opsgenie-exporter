package main

import (
	"fmt"
	"opsgenie-exporter/internal/domain/opsgenieDomain"
	"opsgenie-exporter/internal/domain/prometheusDomain"
)

func main() {

	opsgenieExporter := opsgenieDomain.New(nil)
	fmt.Println(opsgenieExporter)
	prometheusDomain.PushMetrics()

}

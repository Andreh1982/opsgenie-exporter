package exporter

import (
	"opsgenie-exporter/internal/infrastructure/opsgenie"
)

type IncidentsTotalbyStatus interface {
	IncidentsTotalbyStatus(string) (int, error)
}

func (e *opsgenieExporter) IncidentsTotalbyStatus(status string) (total int, err error) {
	total, err = opsgenie.IncidentsTotalbyStatus(status)
	if err != nil {
		return 0, err
	}
	return total, nil
}

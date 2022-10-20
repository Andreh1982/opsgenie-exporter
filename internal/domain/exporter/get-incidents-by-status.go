package exporter

import (
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/infrastructure/opsgenie"
)

type IncidentsTotalbyStatus interface {
	IncidentsTotalbyStatus(appcontext.Context, string) (int, error)
}

func (e *opsgenieExporter) IncidentsTotalbyStatus(ctx appcontext.Context, status string) (total int, err error) {
	total, err = opsgenie.IncidentsTotalbyStatus(ctx, status)
	if err != nil {
		return 0, err
	}
	return total, nil
}

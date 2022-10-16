package exporter

import (
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/infrastructure/opsgenie"
)

type PostmortemTotalbyIncidentStatus interface {
	PostmortemTotalbyIncidentStatus(ctx appcontext.Context, status string) (counterPostmortemClosed int, counterPostmortemResolved int)
}

func (o *opsgenieExporter) PostmortemTotalbyIncidentStatus(ctx appcontext.Context, status string) (int, int) {

	counterPostmortemClosed, counterPostmortemResolved := opsgenie.CheckPostMortems(ctx, status)

	return counterPostmortemClosed, counterPostmortemResolved
}

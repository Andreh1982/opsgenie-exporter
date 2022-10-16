package exporter

import "opsgenie-exporter/internal/domain/appcontext"

type UseCases interface {
	IncidentsTotalbyStatus
	IncidentsTotalbyTeams
	PostmortemTotalbyIncidentStatus
	PostmortemTotalbyTeams
}

type Input struct{}

type opsgenieExporter struct{}

func New(ctx appcontext.Context, input *Input) UseCases {
	return &opsgenieExporter{}
}

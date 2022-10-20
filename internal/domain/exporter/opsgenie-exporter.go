package exporter

import (
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/infrastructure/logger/logwrapper"
)

type UseCases interface {
	IncidentsTotalbyStatus
	IncidentsTotalbyTeams
	PostmortemTotalbyIncidentStatus
	PostmortemTotalbyTeams
}

type Input struct{}

type opsgenieExporter struct{}

func New(ctx appcontext.Context, input *Input, logger logwrapper.LoggerWrapper) UseCases {
	return &opsgenieExporter{}
}

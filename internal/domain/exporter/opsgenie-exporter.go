package exporter

import "opsgenie-exporter/internal/domain/appcontext"

type UseCases interface {
	IncidentsTotalbyStatus
	IncidentsTotalbyTeams
	PostmortemTotalbyIncidentStatus
	PostmortemTotalbyTeams
}

type Input struct{}

type opsgenieExporter struct {
	CounterPostmortem            int `json:"counterPostmortem"`
	CounterPostmortemClosed      int `json:"counterPostmortemClosed"`
	CounterPostmortemResolved    int `json:"counterPostmortemResolved"`
	CounterTeamIncidentsClosed   int `json:"counterTeamIncidentsClosed"`
	CounterTeamIncidentsResolved int `json:"counterTeamIncidentsResolved"`
}

func New(ctx appcontext.Context, input *Input) UseCases {
	return &opsgenieExporter{}
}

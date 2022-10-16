package exporter

import (
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/infrastructure/opsgenie"
)

var MetricStorage []string

type IncidentsTotalbyTeamsStruct []struct {
	TeamName string `json:"teamname"`
	Total    int    `json:"total"`
}

type IncidentsTotalbyTeams interface {
	IncidentsTotalbyTeams(ctx appcontext.Context) ([]IncidentsTotalbyTeamsStruct, error)
}

func (o *opsgenieExporter) IncidentsTotalbyTeams(ctx appcontext.Context) ([]IncidentsTotalbyTeamsStruct, error) {
	opsgenie.CountTeamsIncident(ctx, "closed")
	opsgenie.CountTeamsIncident(ctx, "resolved")
	return nil, nil
}

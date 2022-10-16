package exporter

import (
	"encoding/json"
	"fmt"
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/infrastructure/api"
	"strings"
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
	countTeamsIncident(ctx, "closed")
	countTeamsIncident(ctx, "resolved")
	return nil, nil
}

func countTeamsIncident(ctx appcontext.Context, status string) (teamsTotal int, err error) {
	var responsePayload TeamsList

	apiUrlString := "https://api.opsgenie.com/v2/teams"

	bodyBytes := api.HandlerSingle("GET", apiUrlString)
	err = json.Unmarshal(bodyBytes, &responsePayload)
	if err != nil {
		return 0, err
	}
	teamsTotal = len(responsePayload.Data)

	listTeamsNames(ctx, responsePayload, teamsTotal, status)

	return teamsTotal, err
}

func listTeamsNames(ctx appcontext.Context, teamList TeamsList, total int, status string) {
	for i := 0; i < total; i++ {
		teamID := teamList.Data[i].ID
		teamName := strings.ReplaceAll(teamList.Data[i].Name, " ", "_")
		teamNameNoHifen := strings.ReplaceAll(teamName, "-", "")
		teamNameNoHifen = strings.ToLower(teamNameNoHifen)
		metricName := "opsgenie_incidents_" + status + "_" + teamNameNoHifen
		counterTeamIncidentsClosed, _ := countTeamsIncidents(ctx, teamID, teamNameNoHifen, status)
		makeMetricVar := "# HELP " + metricName + " TEAM TOTAL number of CLOSED incidents\n" + "# TYPE " + metricName + " gauge\n" + metricName + " " + fmt.Sprintf("%d", counterTeamIncidentsClosed)
		addSlice := makeMetricVar
		MetricStorage = append(MetricStorage, addSlice)
	}
	api.GetIncidentsTeamsMetrics(MetricStorage)
}

func countTeamsIncidents(ctx appcontext.Context, teamID string, teamName string, status string) (int, int) {
	var responseIncidentList IncidentList
	counterTeamIncidentsClosed := 0
	counterTeamIncidentsResolved := 0

	apiIncidentList := "https://api.opsgenie.com/v1/incidents?limit=100"
	bodyBytesTimeline := api.HandlerSingle("GET", apiIncidentList)
	json.Unmarshal(bodyBytesTimeline, &responseIncidentList)

	if status == "closed" {
		for i := 0; i < len(responseIncidentList.Data); i++ {
			checkType := responseIncidentList.Data[i].Responders[0].Type
			if checkType == "team" {
				getTeamID := responseIncidentList.Data[i].Responders[0].ID
				if getTeamID == teamID {
					counterTeamIncidentsClosed++
					ctx.SetTotalTeamIncidentsClosed(counterTeamIncidentsClosed)
				}
			}
		}
	}
	if status == "resolved" {
		for i := 0; i < len(responseIncidentList.Data); i++ {
			checkType := responseIncidentList.Data[i].Responders[0].Type
			if checkType == "team" {
				getTeamID := responseIncidentList.Data[i].Responders[0].ID
				if getTeamID == teamID {
					counterTeamIncidentsResolved++
					ctx.SetTotalTeamIncidentsClosed(counterTeamIncidentsResolved)
				}
			}
		}
	}
	return counterTeamIncidentsClosed, counterTeamIncidentsResolved
}

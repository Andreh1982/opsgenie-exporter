package metrics

import (
	"encoding/json"
	"fmt"
	"opsgenie-exporter/internal/environment"
	"opsgenie-exporter/internal/handlers"
	"opsgenie-exporter/internal/models"
	"strings"
)

var MetricStorage []string

func CountTeams(status string) (teamsTotal int) {

	var responsePayload models.TeamsList

	environment.InitEnv()
	apiUrlString := "https://api.opsgenie.com/v2/teams"

	bodyBytes := handlers.HandlerSingle("GET", apiUrlString)
	json.Unmarshal(bodyBytes, &responsePayload)

	teamsTotal = len(responsePayload.Data)

	ListTeamsNames(responsePayload, teamsTotal, status)

	return teamsTotal
}

func ListTeamsNames(teamList models.TeamsList, total int, status string) {

	for i := 0; i < total; i++ {
		teamID := teamList.Data[i].ID
		teamName := strings.ReplaceAll(teamList.Data[i].Name, " ", "_")
		teamNameNoHifen := strings.ReplaceAll(teamName, "-", "")
		teamNameNoHifen = strings.ToLower(teamNameNoHifen)
		metricName := "opsgenie_incidents_" + status + "_" + teamNameNoHifen
		counterTeamIncidentsClosed = 0
		CountTeamsIncidents(teamID, teamNameNoHifen, status)
		makeMetricVar := "# HELP " + metricName + " TEAM TOTAL number of CLOSED incidents\n" + "# TYPE " + metricName + " gauge\n" + metricName + " " + fmt.Sprintf("%d", counterTeamIncidentsClosed)
		addSlice := makeMetricVar
		MetricStorage = append(MetricStorage, addSlice)
	}
	handlers.GetIncidentsTeamsMetrics(MetricStorage)
}

func CountTeamsIncidents(teamID string, teamName string, status string) {

	var responseIncidentList models.IncidentList
	apiIncidentList := "https://api.opsgenie.com/v1/incidents?limit=100"
	bodyBytesTimeline := handlers.HandlerSingle("GET", apiIncidentList)
	json.Unmarshal(bodyBytesTimeline, &responseIncidentList)

	if status == "closed" {
		for i := 0; i < len(responseIncidentList.Data); i++ {
			checkType := responseIncidentList.Data[i].Responders[0].Type
			if checkType == "team" {
				getTeamID := responseIncidentList.Data[i].Responders[0].ID
				if getTeamID == teamID {
					counterTeamIncidentsClosed++
				}
			}
		}
	}
}

package opsgenie

import (
	"encoding/json"
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/infrastructure/api"
	"opsgenie-exporter/internal/infrastructure/environment"
	"strings"

	"go.uber.org/zap"
)

var counterPostmortemClosed int
var counterPostmortemResolved int

func CheckPostMortems(ctx appcontext.Context, status string) (int, int) {

	logger := ctx.Logger()

	logger.Debug("Checking Postmortems from incidents", zap.String("incident status", status))

	counterPostmortemClosed = 0
	counterPostmortemResolved = 0

	apiUrl := environment.GetInstance().OPSGENIE_API_URL

	var responsePayload IncidentList
	var responsePayloadAdd IncidentList
	var responsePayloadFull IncidentList

	if status == "closed" {
		apiUrlString = apiUrl + "incidents?query=status%3Aclosed&offset=0&limit=200&sort=createdAt&order=desc"
	} else if status == "resolved" {
		apiUrlString = apiUrl + "incidents?query=status%3Aresolved&offset=0&limit=200&sort=createdAt&order=desc"
	}

	bodyBytes := api.HandlerSingle("GET", apiUrlString)
	json.Unmarshal(bodyBytes, &responsePayload)
	total := len(responsePayload.Data)

	apiTotal := responsePayload.TotalCount
	pageUrl := responsePayload.Paging.Last
	if apiTotal > 100 {
		if status == "closed" {
			apiUrlString = apiUrl + "incidents?query=status%3Aclosed&offset=100&limit=300&sort=createdAt&order=desc"
		} else if status == "resolved" {
			apiUrlString = apiUrl + "incidents?query=status%3Aresolved&offset=100&limit=300&sort=createdAt&order=desc"
		}
		addBodyBytes := api.HandlerSingle("GET", pageUrl)
		json.Unmarshal(addBodyBytes, &responsePayloadAdd)
		fullBodyBytes := append(responsePayload.Data, responsePayloadAdd.Data...)
		json.Marshal(fullBodyBytes)
		responsePayloadFull.Data = fullBodyBytes

		for i := 0; i < total; i++ {
			fullID, _ := json.Marshal(responsePayloadFull.Data[i].ID)
			stringID := strings.Replace(string(fullID), "\"", "", -1)
			countPostmortemsFromIncidents(ctx, status, stringID)
		}
	} else {
		for i := 0; i < total; i++ {
			fullID, _ := json.Marshal(responsePayload.Data[i].ID)
			stringID := strings.Replace(string(fullID), "\"", "", -1)
			countPostmortemsFromIncidents(ctx, status, stringID)
		}
	}

	logger.Debug("Checking Postmortems Done", zap.String("incident status", status))
	return counterPostmortemClosed, counterPostmortemResolved
}

func countPostmortemsFromIncidents(ctx appcontext.Context, status string, fullID string) (int, int) {

	logger := ctx.Logger()

	logger.Debug("Counting Postmortems from incidents", zap.String("incident status", status))

	var responseTimeline IncidentTimeline
	apiIncidentTimeLine := "https://api.opsgenie.com/v2/incident-timelines/" + fullID + "/entries"
	bodyBytesTimeline := api.HandlerSingle("GET", apiIncidentTimeLine)
	json.Unmarshal(bodyBytesTimeline, &responseTimeline)

	if status == "closed" {
		for i := 0; i < len(responseTimeline.Data.Entries); i++ {
			checkText := responseTimeline.Data.Entries[i].Description.Content
			allLowerCase := strings.ToLower(checkText)
			if strings.Contains(allLowerCase, "postmortem is published") {
				counterPostmortemClosed++
				ctx.SetTotalPostmortemClosed(counterPostmortemClosed)
			}
		}
	} else if status == "resolved" {
		for i := 0; i < len(responseTimeline.Data.Entries); i++ {
			checkText := responseTimeline.Data.Entries[i].Description.Content
			allLowerCase := strings.ToLower(checkText)
			if strings.Contains(allLowerCase, "postmortem is published") {
				counterPostmortemResolved++
				ctx.SetTotalPostmortemResolved(counterPostmortemResolved)
			}
		}
	}
	sendPostmortemClosed := ctx.GetTotalPostmortemClosed()
	sendPostmortemResolved := ctx.GetTotalPostmortemResolved()
	logger.Debug("Counting Postmortems Done", zap.String("incident status", status))
	return sendPostmortemClosed, sendPostmortemResolved
}

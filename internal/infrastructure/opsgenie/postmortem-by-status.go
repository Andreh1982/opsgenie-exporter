package opsgenie

import (
	"encoding/json"
	"fmt"
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/infrastructure/api"
	"opsgenie-exporter/internal/infrastructure/environment"
	"strings"
)

var responsePayload IncidentList
var responsePayloadAdd IncidentList
var responsePayloadFull IncidentList
var counterPostmortemClosed int
var counterPostmortemResolved int

func CheckPostMortems(ctx appcontext.Context, status string) (int, int) {

	apiUrl := environment.GetInstance().OPSGENIE_API_URL

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

		fmt.Println("# Incidents with status "+status, apiTotal)

		for i := 0; i < total; i++ {
			fullID, _ := json.Marshal(responsePayloadFull.Data[i].ID)
			stringID := strings.Replace(string(fullID), "\"", "", -1)
			countPostmortemsFromIncidents(ctx, status, stringID)
		}
	} else {

		BodyBytes := api.HandlerSingle("GET", apiUrlString)
		json.Unmarshal(BodyBytes, &responsePayload)
		fmt.Println("# Incidents with status "+status, responsePayload.TotalCount)

		for i := 0; i < total; i++ {
			fullID, _ := json.Marshal(responsePayload.Data[i].ID)
			stringID := strings.Replace(string(fullID), "\"", "", -1)
			counterPostmortemClosed, counterPostmortemResolved = countPostmortemsFromIncidents(ctx, status, stringID)
		}
	}
	if status == "closed" {
		fmt.Println("# Postmortem Closed Total " + fmt.Sprint(counterPostmortemClosed))
	} else if status == "resolved" {
		fmt.Println("# Postmortem Resolved Total " + fmt.Sprint(counterPostmortemResolved))
	}
	return counterPostmortemClosed, counterPostmortemResolved
}

func countPostmortemsFromIncidents(ctx appcontext.Context, status string, fullID string) (int, int) {

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
	return counterPostmortemClosed, counterPostmortemResolved
}

package exporter

import (
	"encoding/json"
	"fmt"
	"opsgenie-exporter/internal/infrastructure/api"
	"opsgenie-exporter/internal/infrastructure/environment"
	"os"
	"strings"
)

type PostmortemTotalbyIncidentStatus interface {
	PostmortemTotalbyIncidentStatus(status string) (counterPostmortem int)
}

var ApiUrl string
var responsePayload IncidentList
var responsePayloadAdd IncidentList
var responsePayloadFull IncidentList

func (o *opsgenieExporter) PostmortemTotalbyIncidentStatus(status string) int {

	env := environment.GetInstance()
	ApiUrl = os.Getenv(env.OPSGENIE_API_URL)

	counterPostmortem := checkPostMortems(status)

	return counterPostmortem
}

func checkPostMortems(status string) int {

	var counterPostmortem int

	if status == "closed" {
		apiUrlString = ApiUrl + "incidents?query=status%3Aclosed&offset=0&limit=200&sort=createdAt&order=desc"
	} else if status == "resolved" {
		apiUrlString = ApiUrl + "incidents?query=status%3Aresolved&offset=0&limit=200&sort=createdAt&order=desc"
	}

	bodyBytes := api.HandlerSingle("GET", apiUrlString)
	json.Unmarshal(bodyBytes, &responsePayload)
	total := len(responsePayload.Data)

	apiTotal := responsePayload.TotalCount
	pageUrl := responsePayload.Paging.Last
	if apiTotal > 100 {
		if status == "closed" {
			apiUrlString = ApiUrl + "incidents?query=status%3Aclosed&offset=100&limit=300&sort=createdAt&order=desc"
		} else if status == "resolved" {
			apiUrlString = ApiUrl + "incidents?query=status%3Aresolved&offset=100&limit=300&sort=createdAt&order=desc"
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
			countPostmortemsFromIncidents(status, stringID)
		}
	} else {

		BodyBytes := api.HandlerSingle("GET", apiUrlString)
		json.Unmarshal(BodyBytes, &responsePayload)
		fmt.Println("# Incidents with status "+status, responsePayload.TotalCount)

		for i := 0; i < total; i++ {
			fullID, _ := json.Marshal(responsePayload.Data[i].ID)
			stringID := strings.Replace(string(fullID), "\"", "", -1)
			counterPostmortem = countPostmortemsFromIncidents(status, stringID)
		}
	}
	if status == "closed" {
		fmt.Println("# Postmortem Total " + fmt.Sprint(counterPostmortem))
	} else if status == "resolved" {
		fmt.Println("# Postmortem Total " + fmt.Sprint(counterPostmortem))
	}
	return counterPostmortem
}

func countPostmortemsFromIncidents(status string, fullID string) (counterPostmortem int) {

	var responseTimeline IncidentTimeline
	apiIncidentTimeLine := "https://api.opsgenie.com/v2/incident-timelines/" + fullID + "/entries"
	bodyBytesTimeline := api.HandlerSingle("GET", apiIncidentTimeLine)
	json.Unmarshal(bodyBytesTimeline, &responseTimeline)

	if status == "closed" {
		for i := 0; i < len(responseTimeline.Data.Entries); i++ {
			checkText := responseTimeline.Data.Entries[i].Description.Content
			allLowerCase := strings.ToLower(checkText)
			if strings.Contains(allLowerCase, "postmortem is published") {
				counterPostmortem = counterPostmortem + 1
			}
		}
	} else if status == "resolved" {
		for i := 0; i < len(responseTimeline.Data.Entries); i++ {
			checkText := responseTimeline.Data.Entries[i].Description.Content
			allLowerCase := strings.ToLower(checkText)
			if strings.Contains(allLowerCase, "postmortem is published") {
				counterPostmortem++
			}
		}
	}
	return counterPostmortem
}

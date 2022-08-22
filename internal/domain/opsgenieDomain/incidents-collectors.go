package opsgenieDomain

import (
	"encoding/json"
	"fmt"
	"opsgenie-exporter/internal/infrastructure/api"
	"opsgenie-exporter/internal/infrastructure/environment"
	"os"
)

var apiUrlString string
var counterPostmortemClosed int = 0
var counterPostmortemResolved int = 0
var counterTeamIncidentsClosed int = 0

func TotalIncidentList() (totalClosed int, totalResolved int, totalOpened int) {

	var bodyBytesClosed []byte
	var respPayload IncidentList

	environment.InitEnv()
	apiUrl := os.Getenv("OPSGENIE_API_URL")
	method := "GET"

	apiUrlString = apiUrl + "incidents?query=status%3Aclosed&offset=0&limit=200&sort=createdAt&order=desc"
	bodyBytesClosed = api.HandlerSingle(method, apiUrlString)
	json.Unmarshal(bodyBytesClosed, &respPayload)
	totalClosed = respPayload.TotalCount

	apiUrlString = apiUrl + "incidents?query=status%3Aresolved&offset=0&limit=200&sort=createdAt&order=desc"
	bodyBytesResolved := api.HandlerSingle(method, apiUrlString)
	json.Unmarshal(bodyBytesResolved, &respPayload)
	totalResolved = respPayload.TotalCount

	apiUrlString = apiUrl + "incidents?query=status%3Aopen&offset=0&limit=200&sort=createdAt&order=desc"
	bodyBytesOpened := api.HandlerSingle(method, apiUrlString)
	json.Unmarshal(bodyBytesOpened, &respPayload)
	totalOpened = respPayload.TotalCount

	GetIdFromAll("resolved")

	return totalClosed, totalResolved, totalOpened
}

func GetIdFromAll(status string) {
	environment.InitEnv()
	apiUrl := os.Getenv("OPSGENIE_API_URL")
	var responsePayload IncidentList
	var responsePayloadAdd IncidentList
	var responsePayloadFull IncidentList
	var total int

	method := "GET"
	if status == "closed" {
		apiUrlString = apiUrl + "incidents?limit=100&sort=createdAt&offset=0&order=desc&query=status%3Aclosed"
	} else if status == "resolved" {
		apiUrlString = apiUrl + "incidents?limit=100&sort=createdAt&offset=0&order=desc&query=status%3Aresolved"
	} else if status == "opened" {
		apiUrlString = apiUrl + "incidents?limit=100&sort=createdAt&offset=0&order=desc&query=status%3Aopened"
	}
	bodyBytes := api.HandlerSingle(method, apiUrlString)
	json.Unmarshal(bodyBytes, &responsePayload)

	apiTotal := responsePayload.TotalCount
	pageUrl := responsePayload.Paging.Last
	if apiTotal > 100 {
		addBodyBytes := api.HandlerSingle(method, pageUrl)
		json.Unmarshal(addBodyBytes, &responsePayloadAdd)
		fullBodyBytes := append(responsePayload.Data, responsePayloadAdd.Data...)
		json.Marshal(fullBodyBytes)
		responsePayloadFull.Data = fullBodyBytes

		total = len(responsePayloadFull.Data)
		fmt.Println("## Incidents "+status, total)
		// for i := 0; i < total; i++ {
		// 	idJson := responsePayloadFull.Data[i].ID
		// 	createdAtJson := responsePayloadFull.Data[i].CreatedAt
		// 	messageJson := responsePayloadFull.Data[i].Message

		// 	fmt.Println(string(idJson), createdAtJson, string(messageJson))
		// }
	} else {
		total = len(responsePayload.Data)
		fmt.Println("# Incidents "+status, total)
		// for i := 0; i < total; i++ {
		// 	idJson := responsePayload.Data[i].ID
		// 	createdAtJson := responsePayload.Data[i].CreatedAt
		// 	messageJson := responsePayload.Data[i].Message

		// 	fmt.Println(string(idJson), createdAtJson, string(messageJson))
		// }
	}
}

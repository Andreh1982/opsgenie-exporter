package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"opsgenie-exporter/internal/infrastructure/environment"
	"os"
)

var genieKey string

func HandlerSingle(method string, url string) []byte {
	environment.InitEnv()
	genieKey = os.Getenv("OPSGENIE_API_KEY")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", genieKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}

func GetIncidentsTeamsMetrics(MetricStorage []string) {

	f, _ := os.Create("./swap-incidents.tmp")
	defer f.Close()
	for i := 0; i < len(MetricStorage); i++ {
		fmt.Fprintln(f, MetricStorage[i])
	}
	f.Sync()
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	payload, _ := os.ReadFile("./swap-incidents.tmp")
	fmt.Fprint(w, string(payload))
}

package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/infrastructure/environment"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var genieKey string

func SetupAPI(ctx appcontext.Context) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/incidents-by-teams", IncidentsByTeam)
	http.ListenAndServe(":2112", nil)
}

func HandlerSingle(method string, url string) []byte {
	env := environment.GetInstance()
	genieKey = env.OPSGENIE_API_KEY
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

func IncidentsByTeam(w http.ResponseWriter, _ *http.Request) {
	payload, _ := os.ReadFile("./swap-incidents.tmp")
	fmt.Fprint(w, string(payload))
}

package metrics

import (
	"net/http"
	"opsgenie-exporter/internal/handlers"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PushMetrics() {
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/incidents-by-teams", handlers.HelloHandler)
	http.ListenAndServe(":2112", nil)
}

func recordMetrics() {
	go func() {
		for {
			totalTeams := CountTeams("closed")
			opsTeamsList.Set(float64(totalTeams))

			closed, resolved, opened := TotalIncidentList()
			opsClosed.Set(float64(closed))
			opsResolved.Set(float64(resolved))
			opsOpened.Set(float64(opened))

			CheckPostMortems("closed")
			opsPostmortemClosed.Set(float64(counterPostmortemClosed))
			counterPostmortemClosed = 0

			CheckPostMortems("resolved")
			opsPostmortemResolved.Set(float64(counterPostmortemResolved))
			counterPostmortemResolved = 0

			time.Sleep(15 * time.Second)
		}
	}()
}

var (
	opsClosed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "opsgenie_incidents_closed",
		Help: "TOTAL number of CLOSED incidents",
	})
	opsResolved = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "opsgenie_incidents_resolved",
		Help: "TOTAL number of RESOLVED incidents",
	})
	opsOpened = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "opsgenie_incidents_opened",
		Help: "TOTAL number of OPENED incidents",
	})
	opsPostmortemClosed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "opsgenie_postmortem_incidents_closed",
		Help: "TOTAL number of CLOSED incidents POSTMORTEM",
	})
	opsPostmortemResolved = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "opsgenie_postmortem_incidents_resolved",
		Help: "TOTAL number of RESOLVED incidents POSTMORTEM",
	})
	opsTeamsList = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "opsgenie_teams_total",
		Help: "TOTAL number of TEAMS",
	})
)

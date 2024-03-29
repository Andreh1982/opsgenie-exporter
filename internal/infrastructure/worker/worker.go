package worker

import (
	"fmt"
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/domain/exporter"
	"opsgenie-exporter/internal/infrastructure/logger/logwrapper"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Input struct {
	Logger                   logwrapper.LoggerWrapper
	OpsgenieExporterUseCases exporter.UseCases
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
	//opsTeamsList = promauto.NewGauge(prometheus.GaugeOpts{
	//	Name: "opsgenie_teams_total",
	//	Help: "TOTAL number of TEAMS",
	//})
)

func Start(ctx appcontext.Context, input Input) {
	appctx := appcontext.NewBackground()
	appctx.SetLogger(input.Logger)
	logger := appctx.Logger()
	logger.Info("Starting worker")
	GetMetrics(appctx)
}

func GetMetrics(ctx appcontext.Context) {
	logger := ctx.Logger()
	logger.Info("Getting Metrics from Opsgenie")
	go func() {
		for {

			//	_, err := opsgenieDomain.New().IncidentsTotalbyTeams(ctx appconappcontext.New())
			//	if err != nil {
			//		fmt.Println(err)
			//	}

			_, err := exporter.New(ctx, &exporter.Input{}, logger).PostmortemTotalbyTeams()
			if err != nil {
				fmt.Println(err)
			}

			closed, _ := exporter.New(ctx, &exporter.Input{}, logger).IncidentsTotalbyStatus(ctx, "closed")
			opsClosed.Set(float64(closed))

			resolved, _ := exporter.New(ctx, &exporter.Input{}, logger).IncidentsTotalbyStatus(ctx, "resolved")
			opsResolved.Set(float64(resolved))

			opened, _ := exporter.New(ctx, &exporter.Input{}, logger).IncidentsTotalbyStatus(ctx, "opened")
			opsOpened.Set(float64(opened))

			counterPostmortemClosed, _ := exporter.New(ctx, &exporter.Input{}, logger).PostmortemTotalbyIncidentStatus(ctx, "closed")
			opsPostmortemClosed.Set(float64(counterPostmortemClosed))

			_, counterPostmortemResolved := exporter.New(ctx, &exporter.Input{}, logger).PostmortemTotalbyIncidentStatus(ctx, "resolved")
			opsPostmortemResolved.Set(float64(counterPostmortemResolved))

			time.Sleep(15 * time.Second)
		}
	}()
}

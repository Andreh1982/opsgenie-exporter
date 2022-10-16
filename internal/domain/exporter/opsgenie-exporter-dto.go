package exporter

type OpsgenieExporterDTO struct {
	CounterPostmortem            int `json:"counterPostmortem"`
	CounterPostmortemClosed      int `json:"counterPostmortemClosed"`
	CounterPostmortemResolved    int `json:"counterPostmortemResolved"`
	CounterTeamIncidentsClosed   int `json:"counterTeamIncidentsClosed"`
	CounterTeamIncidentsResolved int `json:"counterTeamIncidentsResolved"`
}

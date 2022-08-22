package opsgenieDomain

type OpsgenieExporter interface {
	IncidentsTotalbyStatus
	IncidentsTotalbyTeams
	PostmortemTotalbyIncidentStatus
	PostmortemTotalbyTeams
}

type opsgenieExporter struct {
	Data []struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	} `json:"data"`
}

func New(exporterData *opsgenieExporter) OpsgenieExporter {
	return &opsgenieExporter{}
}

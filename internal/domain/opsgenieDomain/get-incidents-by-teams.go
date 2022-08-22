package opsgenieDomain

type IncidentsTotalbyTeamsStruct []struct {
	TeamName string `json:"teamname"`
	Total    int    `json:"total"`
}

type IncidentsTotalbyTeams interface {
	IncidentsTotalbyTeams() ([]IncidentsTotalbyTeamsStruct, error)
}

func (o *opsgenieExporter) IncidentsTotalbyTeams() ([]IncidentsTotalbyTeamsStruct, error) {
	return nil, nil
}

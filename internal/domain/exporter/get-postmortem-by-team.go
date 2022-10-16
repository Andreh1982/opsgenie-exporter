package exporter

type PostmortemTotalbyTeamsStruct []struct {
	TeamName string `json:"teamname"`
	Total    int    `json:"total"`
}

type PostmortemTotalbyTeams interface {
	PostmortemTotalbyTeams() ([]PostmortemTotalbyTeamsStruct, error)
}

func (o *opsgenieExporter) PostmortemTotalbyTeams() ([]PostmortemTotalbyTeamsStruct, error) {
	return nil, nil
}

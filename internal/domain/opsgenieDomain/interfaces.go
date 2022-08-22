package opsgenieDomain

type PostmortemTotals interface {
	CountPostmortem() int
	ResolvedPostmortemTotal() int
	ClosedPostmortemTotal() int
}

type IncidentsTotals interface {
	IncidentsResolvedTotal() int
	IncidentsClosedTotal() int
	IncidentsOpenedTotal() int
}

type TeamsIncidentsTotals interface {
	CountTeams() int
	ListTeams(TeamsList, int, string)
	IncidentsTotalTeams(string, string, string)
}

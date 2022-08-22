package opsgenieDomain

type PostmortemTotalbyIncidentStatus interface {
	runPostmortemTotalbyIncidentStatus() (int, int)
}

func (o *opsgenieExporter) runPostmortemTotalbyIncidentStatus() (int, int) {
	return 0, 0
}

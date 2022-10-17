package appcontext

import (
	"context"
	"opsgenie-exporter/internal/infrastructure/logger/logwrapper"
)

type Context interface {
	SetLogger(logger logwrapper.LoggerWrapper)
	Logger() logwrapper.LoggerWrapper
	SetTotalTeamIncidentsClosed(counterTeamIncidentsClosed int) int
	SetTotalTeamIncidentsResolved(counterTeamIncidentsResolved int) int
	SetTotalPostmortemResolved(int)
	GetTotalPostmortemResolved() int
	SetTotalPostmortemClosed(int)
	GetTotalPostmortemClosed() int
}

type appContext struct {
	logger                     logwrapper.LoggerWrapper
	ctx                        context.Context
	totalTeamIncidentsClosed   int
	totalTeamIncidentsResolved int
	totalPostmortemClosed      int
	totalPostmortemResolved    int
}

func New(ctx context.Context) Context {
	return &appContext{
		ctx: ctx,
	}
}

func NewBackground() Context {
	ctx := context.Background()
	return &appContext{
		ctx: ctx,
	}
}

func (appContext *appContext) SetLogger(logger logwrapper.LoggerWrapper) {
	appContext.logger = logger
}

func (appContext *appContext) Logger() logwrapper.LoggerWrapper {
	return appContext.logger
}

func (appContext *appContext) SetTotalTeamIncidentsClosed(counterTeamIncidentsClosed int) int {
	appContext.totalTeamIncidentsClosed = counterTeamIncidentsClosed
	return appContext.totalTeamIncidentsClosed
}

func (appContext *appContext) SetTotalTeamIncidentsResolved(counterTeamIncidentsResolved int) int {
	appContext.totalTeamIncidentsResolved = counterTeamIncidentsResolved
	return appContext.totalTeamIncidentsResolved
}

func (appContext *appContext) SetTotalPostmortemResolved(totalPostmortemResolved int) {
	appContext.totalPostmortemResolved = totalPostmortemResolved
}

func (appContext *appContext) GetTotalPostmortemResolved() int {
	return appContext.totalPostmortemResolved
}

func (appContext *appContext) SetTotalPostmortemClosed(totalPostmortemClosed int) {
	appContext.totalPostmortemClosed = totalPostmortemClosed
}

func (appContext *appContext) GetTotalPostmortemClosed() int {
	return appContext.totalPostmortemClosed
}

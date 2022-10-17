package main

import (
	"opsgenie-exporter/internal/domain/appcontext"
	"opsgenie-exporter/internal/domain/exporter"
	"opsgenie-exporter/internal/infrastructure/api"
	"opsgenie-exporter/internal/infrastructure/environment"
	"opsgenie-exporter/internal/infrastructure/logger"
	"opsgenie-exporter/internal/infrastructure/logger/logwrapper"
	"opsgenie-exporter/internal/infrastructure/worker"

	"go.uber.org/zap"
)

func main() {

	ctx := appcontext.NewBackground()

	env := environment.GetInstance()
	zaplogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zaplogger}).SetVersion("1.0")
	logger.Info("Starting OpsGenie-Prometheus Exporter")

	logger.Info("Initializing OpsGenie Client",
		zap.String("API Endpoint:", env.OPSGENIE_API_URL),
	)

	opsgenieExporterUseCases, err := setupOpsgenieExporter(ctx)

	if err != nil {
		logger.Error("failed to configure opsgenie exporter", zap.Error(err))
	}

	setupWorker(ctx, logger, opsgenieExporterUseCases)
	api.SetupAPI(ctx)
}

func setupOpsgenieExporter(ctx appcontext.Context) (exporter.UseCases, error) {
	exporterInput := &exporter.Input{}
	opsgenieExporterUseCases := exporter.New(ctx, exporterInput)
	return opsgenieExporterUseCases, nil
}

func setupWorker(ctx appcontext.Context, logger logwrapper.LoggerWrapper, opsgenieExporterUseCases exporter.UseCases) {
	logger.Info("Initializing Worker")
	input := worker.Input{
		Logger:                   logger,
		OpsgenieExporterUseCases: opsgenieExporterUseCases,
	}
	worker.Start(ctx, input)
}

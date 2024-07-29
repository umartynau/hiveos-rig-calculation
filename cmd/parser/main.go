package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/zerologWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"path/filepath"
	"rig-calculation/internal/config"
	"rig-calculation/internal/model"
	"rig-calculation/internal/service"
	"strconv"
	"time"
)

func main() {
	configOrigin := flag.String("config", "", "config origin: (file)")
	flag.Parse()

	cfg, err := initConfig(configOrigin)
	if err != nil {
		log.Panic().Err(err).Msg("initConfig")
	}

	var telemetry *newrelic.Application
	if cfg.App.Telemetry.Enabled {
		telemetry, err = newrelic.NewApplication(
			newrelic.ConfigAppName(cfg.App.Telemetry.NewrelicName),
			newrelic.ConfigLicense(cfg.App.Telemetry.NewrelicKey),
			newrelic.ConfigAppLogForwardingEnabled(true),
		)
		if err != nil {
			log.Error().Err(err).Msg("newrelic telemetry")
		}
		defer func() {
			log.Info().Msg("shutting newrelic telemetry...")
			telemetry.Shutdown(time.Second * 10)
		}()
	}

	initGlobalLogger(cfg.App, telemetry)

	hiveService := service.NewHive(cfg.Hive)

	farms, err := hiveService.GetFarms()
	if err != nil {
		log.Panic().Err(err).Msg("hive")
	}

	var report []*model.Report

	for _, farm := range farms.Data {
		workers, err := hiveService.GetWorkers(farm.ID)
		if err != nil {
			log.Error().Err(err).Int("farm_id", farm.ID).Msg("hiveService GetWorkers")
		}

		for _, worker := range workers.Data {
			stats, err := hiveService.GetWorkerStats(farm.ID, worker.ID)
			if err != nil {
				log.Error().Err(err).Int("farm_id", farm.ID).Int("worker_id", worker.ID).Msg("hiveService GetWorkerStats")
			}

			workerHours := 0
			for _, hashrate := range stats.Data {
				if len(hashrate.Hashrates) != 0 {
					workerHours++
				}
			}

			reportEntry := &model.Report{
				FarmName:          farm.Name,
				FarmId:            farm.ID,
				WorkerName:        worker.Name,
				WorkerId:          worker.ID,
				WorkerDescription: worker.Description,
				UptimeHours:       workerHours,
			}

			report = append(report, reportEntry)

			log.Info().Int("farm_id", farm.ID).Int("worker_id", worker.ID).Int("hours", workerHours).Msg("worker uptime")
		}
	}

	err = writeReportsToCSV(report, "uptime_report.csv")
	if err != nil {
		log.Error().Err(err).Msg("Failed to write reports to CSV")
	} else {
		log.Info().Msg("Reports successfully written to CSV")
	}
}

func initConfig(configOrigin *string) (*config.Config, error) {
	if configOrigin == nil || len(*configOrigin) == 0 {
		log.Panic().Msg("no config file, use --config=file|env")
	}

	var cfg *config.Config
	var err error

	switch *configOrigin {
	case "file":
		cfg, err = config.NewConfig("config.yaml")
		if err != nil {
			log.Panic().Err(err).Msg("config error")
		}
	default:
		log.Panic().Msg("unsupported config source, use --config=file")
	}

	return cfg, nil
}

func initGlobalLogger(cfg config.App, telemetry *newrelic.Application) {
	devWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFieldFormat, NoColor: !cfg.LogColorEnabled}

	l := log.With().Timestamp().Str("service", cfg.InstanceName).Str("service_label", cfg.InstanceLabel)
	if cfg.LogTraceEnabled {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.InterfaceMarshalFunc = json.Marshal
		l = l.Caller().Stack()
	}

	if cfg.Telemetry.Enabled && telemetry != nil {
		telemetryWriter := zerologWriter.New(os.Stdout, telemetry)
		telemetryWriter.DebugLogging(cfg.Telemetry.DebugEnabled)
		log.Logger = l.Logger().Output(telemetryWriter)
	} else {
		log.Warn().Msg("telemetry disabled")
		log.Logger = l.Logger().Output(devWriter)
	}

	logLevel := zerolog.Level(cfg.LogLevel)
	log.Info().Msgf("Log level: %s", logLevel)

	zerolog.SetGlobalLevel(logLevel)
}

func writeReportsToCSV(report []*model.Report, filename string) error {
	outputPath := filepath.Join("./output", filename)

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"FarmName", "FarmId", "WorkerName", "WorkerId", "Description", "UptimeHours"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, entry := range report {
		record := []string{
			entry.FarmName,
			strconv.Itoa(entry.FarmId),
			entry.WorkerName,
			strconv.Itoa(entry.WorkerId),
			entry.WorkerDescription,
			strconv.Itoa(entry.UptimeHours),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

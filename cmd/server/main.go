package main

import (
	"balun_homework_1/business/database"
	"balun_homework_1/foundation/compute"
	"balun_homework_1/foundation/logger"
	"balun_homework_1/foundation/network"
	"balun_homework_1/foundation/storage"
	"balun_homework_1/foundation/utils"
	"context"
	"fmt"
	"os"
	"time"
)

const configPath = "config.yaml"

func main() {
	cfg := NewConfigFromFile(configPath)

	log, err := initLogger(cfg.Logging)

	if err != nil {
		fmt.Printf("Logger initiation error: %w", err)
	}

	ctx := context.Background()

	if err := run(ctx, cfg, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
	}
}

func initLogger(cfg LoggingConfig) (*logger.Logger, error) {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT ******")
		},
	}

	logLevel, err := logger.ConvertHumanReadableLevelToLevel(cfg.Level)

	if err != nil {
		return nil, fmt.Errorf("parse log level error: %w", err)
	}

	logOutput, err := os.OpenFile(cfg.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, fmt.Errorf("file open error: %w", err)
	}

	return logger.NewWithEvents(logOutput, logLevel, "Database", nil, events), nil
}

func run(ctx context.Context, cfg ServerConfig, log *logger.Logger) error {
	log.Info(ctx, "[Main::run] Initialize application")

	db, err := initiateDatabase(ctx, cfg.Engine, log)

	if err != nil {
		return fmt.Errorf("initiate db error: %w", err)
	}

	idleTimeout, err := time.ParseDuration(cfg.Network.IdleTimeout)

	if err != nil {
		return fmt.Errorf("parse idle timeout error: %w", err)
	}

	maxMessageSize, err := utils.ParseHumanReadableSize(cfg.Network.MaxMessageSize)

	if err != nil {
		return fmt.Errorf("parse max message size error: %w", err)
	}

	server := network.NewTCPServer(
		ctx,
		db,
		log,
		cfg.Network.Address,
		cfg.Network.MaxConnections,
		idleTimeout,
		maxMessageSize,
	)

	defer server.Stop()

	err = server.Run()

	if err != nil {
		return fmt.Errorf("server sun error: %w", err)
	}

	return nil
}

func initiateDatabase(ctx context.Context, cfg EngineConfig, log *logger.Logger) (*database.Database, error) {
	if cfg.Type != defaultEngineType {
		return nil, fmt.Errorf("invalid engine type")
	}

	e := storage.NewInMemoryEngine()

	st := storage.NewInMemoryStorage(e, log)

	p := compute.NewStdParser()

	cp := compute.NewStdCompute(p, log)

	return database.NewDatabase(ctx, st, cp, log), nil
}

package main

import (
	"balun_homework_1/business/database"
	"balun_homework_1/foundation/compute"
	"balun_homework_1/foundation/logger"
	"balun_homework_1/foundation/storage"
	"bufio"
	"context"
	"fmt"
	"os"
)

func main() {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT ******")
		},
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "Database", nil, events)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "[Main::run] Initialize application")

	db := initiateDatabase(ctx, log)

	reader := bufio.NewReader(os.Stdin)

	for {
		query, err := reader.ReadString('\n')

		if err != nil {
			log.Error(ctx, "[Main::run] Cannot read line from input", err)

			continue
		}

		res, err := db.Execute(query)

		if err != nil {
			log.Error(ctx, "[Main::run] Error occurred while executing query", query, err)

			continue
		}

		fmt.Println(res)
	}

}

func initiateDatabase(ctx context.Context, log *logger.Logger) *database.Database {
	e := storage.NewMapEngine()

	st := storage.NewMapStorage(e, log)

	p := compute.NewStdParser()

	cp := compute.NewStdCompute(p, log)

	return database.NewDatabase(ctx, st, cp, log)
}

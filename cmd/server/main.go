package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/vxdiazdel/rest-api/internal/db"
	"github.com/vxdiazdel/rest-api/internal/db/stores"
	"github.com/vxdiazdel/rest-api/internal/logger"
	"github.com/vxdiazdel/rest-api/internal/router"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("load .env: %w", err))
	}

	// logger
	slg := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(slg)
	lg := logger.NewSLogger(ctx, slg)

	// db
	dbConn := db.NewPostgresConn(ctx, os.Getenv("DB_URL"))
	defer dbConn.Close(ctx)

	// stores
	store := stores.NewPostgresStore(ctx, dbConn, lg)

	// clients

	// router
	r := router.NewRouter(ctx, store, lg)
	r.Run()
}

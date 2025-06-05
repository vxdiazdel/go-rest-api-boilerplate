package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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

	// sessions
	sessionStore, err := redis.NewStore(
		10,
		"tcp",
		os.Getenv("REDIS_URL"),
		"",
		os.Getenv("REDIS_PASSWORD"),
		[]byte(os.Getenv("SESSION_SECRET")),
	)
	if err != nil {
		panic(fmt.Errorf("create session store: %w", err))
	}
	sessionStore.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(7 * 24 * time.Hour.Seconds()),
		Secure:   os.Getenv("GIN_MODE") == "production",
	})

	// clients

	// router
	r := router.NewRouter(ctx, store, sessionStore, lg)
	r.Run()
}

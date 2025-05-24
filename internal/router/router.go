package router

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vxdiazdel/rest-api/internal/db/stores"
	"github.com/vxdiazdel/rest-api/internal/handlers"
	"github.com/vxdiazdel/rest-api/internal/logger"
)

func NewRouter(
	ctx context.Context,
	store stores.IStore,
	lg logger.ILogger,
) *gin.Engine {
	r := gin.Default()

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// load middleware

	// load handler
	h := handlers.NewHandlerContext(ctx, store, lg)

	// base endpoint
	api := r.Group("/v1")
	api.GET("/ping", h.Ping)

	// auth endpoints

	return r
}

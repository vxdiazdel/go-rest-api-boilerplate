package router

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vxdiazdel/rest-api/internal/db/stores"
	"github.com/vxdiazdel/rest-api/internal/handlers"
	"github.com/vxdiazdel/rest-api/internal/logger"
	"github.com/vxdiazdel/rest-api/internal/session"
	"github.com/vxdiazdel/rest-api/middleware"
)

func NewRouter(
	ctx context.Context,
	store stores.IStore,
	sessionStore sessions.Store,
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
	r.Use(sessions.Sessions(session.UserSession, sessionStore))
	amw := middleware.NewAuthMiddleware(ctx, store, lg)

	// load handler
	h := handlers.NewHandlerContext(ctx, store, lg)

	// base endpoint
	api := r.Group("/v1")

	// auth endpoints
	auth := api.Group("/auth")
	auth.POST("/signup", h.SignUp)
	auth.POST("/login", h.Login)
	auth.POST("/logout", h.Logout)

	// user endpoints
	users := api.Group("/users")
	users.Use(amw.RequireAuth())
	users.GET("/self", h.GetSelf)
	users.GET("/:id", h.GetUserByID)

	return r
}

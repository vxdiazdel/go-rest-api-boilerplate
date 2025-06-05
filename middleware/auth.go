package middleware

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vxdiazdel/rest-api/internal/db/stores"
	"github.com/vxdiazdel/rest-api/internal/logger"
	"github.com/vxdiazdel/rest-api/internal/session"
	"github.com/vxdiazdel/rest-api/utils"
)

type AuthMiddleware struct {
	ctx   context.Context
	store stores.IStore
	lg    logger.ILogger
}

func NewAuthMiddleware(
	ctx context.Context,
	store stores.IStore,
	lg logger.ILogger,
) *AuthMiddleware {
	return &AuthMiddleware{
		ctx:   ctx,
		store: store,
		lg:    lg,
	}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		us := s.Get(session.UserSession)

		if us == nil {
			s.Delete(session.UserSession)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// get session token
		token := us.(string)
		if token == "" {
			s.Delete(session.UserSession)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		m.Lg().Log("RequireAuth", "token", token)

		// parse token
		claims, err := utils.VerifyToken(token)
		if err != nil {
			m.Lg().Error("VerifyToken Error", "error", err.Error())
			s.Delete(session.UserSession)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		m.Lg().Log("VerifyToken", "claims", claims)

		// check if user exists
		user, err := m.Store().GetUserByID(m.Ctx(), claims.UserID)
		if err != nil {
			s.Delete(session.UserSession)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// save user session
		c.Set("AuthUser", user)
		c.Set("AuthUserID", user.ID)

		c.Next()
	}
}

func (m *AuthMiddleware) Ctx() context.Context {
	return m.ctx
}

func (m *AuthMiddleware) Store() stores.IStore {
	return m.store
}

func (m *AuthMiddleware) Lg() logger.ILogger {
	return m.lg
}

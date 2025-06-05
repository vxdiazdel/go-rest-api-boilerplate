package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vxdiazdel/rest-api/internal/session"
	"github.com/vxdiazdel/rest-api/models"
	"github.com/vxdiazdel/rest-api/utils"
)

func (h *HandlerContext) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	// read request data
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// create user
	user, err := h.Store().CreateUser(h.Ctx(), req.Email, string(hashedPassword))
	if err != nil {
		h.Lg().Error("CreateUser", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	// create session
	err = session.CreateUserSession(sessions.Default(c), session.UserSession, user.ID.String())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create session"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (h *HandlerContext) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.Store().GetUserByID(h.Ctx(), userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *HandlerContext) GetSelf(c *gin.Context) {
	// read userID from context
	userID, ok := c.Get("AuthUserID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	authUserID := userID.(uuid.UUID)

	user, err := h.Store().GetUserByID(h.Ctx(), authUserID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

package handlers

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/vxdiazdel/rest-api/internal/session"
	"github.com/vxdiazdel/rest-api/models"
	"github.com/vxdiazdel/rest-api/utils"
)

func (h *HandlerContext) SignUp(c *gin.Context) {
	var req models.SignUpRequest

	// read request data
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "hash password failed"})
		return
	}

	// create user
	user, err := h.Store().CreateUser(h.Ctx(), req.Email, string(hashedPassword))
	if err != nil {
		h.Lg().Error("CreateUser", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "create user failed"})
		return
	}

	// create auth token
	token, err := utils.SignToken(user.ID)
	if err != nil {
		h.Lg().Error("SignToken", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "sign token failed"})
		return
	}

	// create session
	err = session.CreateUserSession(sessions.Default(c), session.UserSession, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "create session failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (h *HandlerContext) Login(c *gin.Context) {
	var req models.LoginRequest

	// read request data
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// get user by email
	user, err := h.Store().GetUserByEmail(h.Ctx(), req.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// verify password
	err = utils.VerifyPassword(user.Password, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// check if user has an existing session
	s := sessions.Default(c)
	jwt := s.Get(session.UserSession)
	if jwt != nil {
		if jwtStr, ok := jwt.(string); ok {
			_, err := utils.VerifyToken(jwtStr)
			if err == nil {
				// token is valid and user is already logged in
				// refresh session expiry
				s.Options(sessions.Options{
					MaxAge: int(7 * 24 * time.Hour.Seconds()),
				})
				if saveErr := s.Save(); saveErr != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "refresh user session failed"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": user})
				return
			}
		}
	}

	// create auth token
	token, err := utils.SignToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "sign token failed"})
		return
	}

	// create session
	err = session.CreateUserSession(sessions.Default(c), session.UserSession, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "create session failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *HandlerContext) Logout(c *gin.Context) {
	// delete user's session
	s := sessions.Default(c)
	s.Clear()
	s.Options(sessions.Options{MaxAge: -1})

	if err := s.Save(); err != nil {
		h.Lg().Error("Logout", "save session", err.Error())
	}

	c.Status(http.StatusOK)
}

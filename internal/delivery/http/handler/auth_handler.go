package handler

import (
	"clean-go/internal/entity"
	"clean-go/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc usecase.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, uc usecase.AuthUsecase) {
	handler := &AuthHandler{
		uc: uc,
	}

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
}

// Register User
func (h *AuthHandler) Register(c *gin.Context) {
	var params entity.Register

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.uc.Register(params.Username, params.Email, params.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login User
func (h *AuthHandler) Login(c *gin.Context) {
	var params entity.Login

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := h.uc.Login(params.Email, params.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

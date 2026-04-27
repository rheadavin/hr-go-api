package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rheadavin/hr-go-api/internal/dto"
	"github.com/rheadavin/hr-go-api/internal/service"
	response "github.com/rheadavin/hr-go-api/pkg/response"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// @Summary Register user
// @Description Registrasi user baru
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register credentials"
// @Success 201 {object} response.Response{data=dto.RegisterResponse}
// @Failure 400 {object} response.Response
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userRegister, err := h.authService.Register(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusCreated, "User registered successfully", userRegister)
}

// @Summary Login user
// @Description Autentikasi user dengan email dan password, return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} response.Response{data=dto.LoginResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.authService.Login(req)
	if err != nil {
		response.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, "Login successfully", result)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("user_email")

	response.SuccessResponse(c, http.StatusOK, "OK", gin.H{
		"user_id": userID,
		"email":   email,
	})
}

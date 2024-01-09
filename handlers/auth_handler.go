package handlers

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input *dto.SignInRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "FIELD_REQUIRED_ERROR",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}
	response, err := h.authService.SignIn(&dto.SignInRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) SignUp(c *gin.Context) {
	var input *dto.SignUpRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "FIELD_REQUIRED_ERROR",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}
	response, err := h.authService.SignUp(&dto.SignUpRequest{
		Email:    input.Email,
		FullName: input.FullName,
		Password: input.Password,
		Address:  input.Address,
		CityID:   input.CityID,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) GetUserDetail(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer ")
	encodedToken := s[1]

	signedOut, _ := h.authService.CheckToken(encodedToken)
	if signedOut {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Token invalid, user signed out",
		})
		return
	}

	payload, _ := c.Get("user")
	userFromPayload, _ := payload.(models.User)
	userId := userFromPayload.ID

	user, err := h.authService.GetUserDetail(userId)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":       "NOT FOUND ERROR",
			"statusCode": http.StatusNotFound,
			"message":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer ")
	encodedToken := s[1]

	signedOut, _ := h.authService.CheckToken(encodedToken)
	if signedOut {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Token invalid, user signed out",
		})
		return
	}

	var input *dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "FIELD_REQUIRED_ERROR",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}
	payload, _ := c.Get("user")
	userFromPayload, _ := payload.(models.User)
	userId := userFromPayload.ID

	user, err := h.authService.UpdateProfile(&dto.UpdateProfileRequest{
		ID:       userId,
		FullName: input.FullName,
		Password: input.Password,
		Address:  input.Address,
		CityID:   input.CityID,
		Role:     input.Role,
	})
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) SignOut(c *gin.Context) {
	payload, _ := c.Get("user")

	fmt.Println(payload)

	if payload == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Unauthorized error, not signed in",
		})
		return
	}

	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer ")
	encodedToken := s[1]

	signedOut, _ := h.authService.CheckToken(encodedToken)
	if signedOut {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Token invalid, user signed out",
		})
		return
	}

	response, err := h.authService.SignOut(encodedToken)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) CleanBlacklists(c *gin.Context) {
	payload, _ := c.Get("user")
	userFromPayload, _ := payload.(models.User)
	role := userFromPayload.Role

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED_ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Unauthorized Error: User Not Admin",
		})
		return
	}

	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, "Bearer ")
	encodedToken := s[1]

	signedOut, _ := h.authService.CheckToken(encodedToken)
	if signedOut {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Token invalid, user signed out",
		})
		return
	}

	response, err := h.authService.CleanBlacklists()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

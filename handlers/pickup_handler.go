package handlers

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) PickUp(c *gin.Context) {
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

	var input *dto.PickUpRequest
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

	pickUp, err := h.pickUpService.PickUp(userId, &dto.PickUpRequest{ReservationID: input.ReservationID})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pickUp})
}

func (h *Handler) UpdatePickUpStatus(c *gin.Context) {
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

	var input *dto.UpdatePickUpStatusRequest
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
	role := userFromPayload.Role

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED_ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Unauthorized Error: User Not Admin",
		})
		return
	}

	response, err := h.pickUpService.UpdatePUStatus(input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) GetPickUps(c *gin.Context) {
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

	pickups, err := h.pickUpService.GetPickUps(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":       "NOT_FOUND_ERROR",
			"statusCode": http.StatusNotFound,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pickups})
}

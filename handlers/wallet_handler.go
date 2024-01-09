package handlers

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) TopUp(c *gin.Context) {
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

	var input *dto.TopUpRequest
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

	wallet, _ := h.walletService.GetWalletFromUserId(userId)
	if wallet == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Unauthorized error, please sign in first.",
		})
		return
	}

	user, _ := h.walletService.TopUp(&dto.TopUpRequest{
		UserID:   userId,
		Wallet:   wallet,
		Amount:   input.Amount,
		SourceID: input.SourceID,
	})

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) Pay(c *gin.Context) {
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

	var input *dto.BookingPaymentRequest
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

	if userId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "Unauthorized error, please sign in first",
		})
		return
	}

	paymentResponse, err := h.walletService.Pay(userId, input)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentResponse})
}

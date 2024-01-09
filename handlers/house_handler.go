package handlers

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/dto"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) ShowHouses(c *gin.Context) {
	searchType := c.Query("searchType")
	search := c.Query("search")
	sortBy := c.Query("sortBy")
	sortType := c.Query("sort")
	start := c.Query("start")
	end := c.Query("end")
	limit := c.Query("limit")
	page := c.Query("page")

	houses, err := h.houseService.GetHouses(searchType, search, sortBy, sortType, start, end, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":       "INTERNAL_STATUS_ERROR",
			"statusCode": http.StatusInternalServerError,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": houses})
}

func (h *Handler) ShowHouse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	house, err := h.houseService.GetHouse(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":       "NOT_FOUND_ERROR",
			"statusCode": http.StatusNotFound,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": house})
}

func (h *Handler) BookHouse(c *gin.Context) {
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

	var input *dto.BookHouseRequest
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

	houseId, _ := strconv.Atoi(c.Param("id"))
	if houseId == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":       "NOT FOUND ERROR",
			"statusCode": http.StatusNotFound,
			"message":    "House not found.",
		})
		return
	}

	reservation, err := h.houseService.CreateReservation(uint(houseId), userId, input)
	if reservation == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "FIELD_REQUIRED_ERROR",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservation})
}

func (h *Handler) ShowBookings(c *gin.Context) {
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

	bookings, err := h.houseService.GetBookings(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":       "NOT_FOUND_ERROR",
			"statusCode": http.StatusNotFound,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

func (h *Handler) GetBookingById(c *gin.Context) {
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

	bookingId, _ := strconv.Atoi(c.Param("bookingId"))

	booking, err := h.houseService.GetBookingById(bookingId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":       "NOT_FOUND_ERROR",
			"statusCode": http.StatusNotFound,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

func (h *Handler) AddHouse(c *gin.Context) {
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

	role, err := h.houseService.CheckRole(userId)
	if err != nil {
		_ = c.Error(err)
		return
	}
	if role != "host" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "User not a host",
		})
		return
	}

	var input *dto.AddHouseRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "FIELD_REQUIRED_ERROR",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}
	response, err := h.houseService.AddHouse(userId, input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) UpdateHouse(c *gin.Context) {
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
	role := userFromPayload.Role
	if role != "host" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "User not a host.",
		})
		return
	}
	var input *dto.UpdateHouseRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "FIELD_REQUIRED_ERROR",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}
	houseId, _ := strconv.Atoi(c.Param("houseId"))
	response, err := h.houseService.UpdateHouse(userId, uint(houseId), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) DeleteHouse(c *gin.Context) {
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
	role := userFromPayload.Role
	if role != "host" && role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "User not a host or an admin",
		})
		return
	}
	houseId, _ := strconv.Atoi(c.Param("houseId"))
	response, err := h.houseService.DeleteHouse(userId, uint(houseId))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) GetPhotosByHouseId(c *gin.Context) {
	houseId, _ := strconv.Atoi(c.Param("houseId"))
	response, err := h.houseService.GetPhotosByHouseId(uint(houseId))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) GetFirstPhoto(c *gin.Context) {
	houseId, _ := strconv.Atoi(c.Param("houseId"))
	response, err := h.houseService.GetFirstPhoto(uint(houseId))
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) UploadPhotos(c *gin.Context) {
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
	role := userFromPayload.Role
	if role != "host" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":       "UNAUTHORIZED ERROR",
			"statusCode": http.StatusUnauthorized,
			"message":    "User not a host",
		})
		return
	}
	var input *dto.UploadPhotosRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "FIELD_REQUIRED_ERROR",
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		return
	}

	houseId, _ := strconv.Atoi(c.Param("houseId"))
	response, err := h.houseService.UploadPhotos(uint(houseId), input)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) GetCities(c *gin.Context) {
	response, err := h.houseService.GetCities()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *Handler) ShowTransactions(c *gin.Context) {
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

	transactions, err := h.houseService.GetTransactions(userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":       "NOT_FOUND_ERROR",
			"statusCode": http.StatusNotFound,
			"message":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

package server

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/handlers"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/middlewares"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type RouterConfig struct {
	AuthService   services.AuthService
	HouseService  services.HouseService
	WalletService services.WalletService
	PickUpService services.PickUpService
	GamesService  services.GamesService
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewRouter(c *RouterConfig) *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "POST", "PATCH", "GET", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization"},
		AllowAllOrigins:  true,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	h := handlers.New(&handlers.HandlerConfig{
		AuthService:   c.AuthService,
		HouseService:  c.HouseService,
		WalletService: c.WalletService,
		PickUpService: c.PickUpService,
		GamesService:  c.GamesService,
	})

	router.Static("/docs", "./swaggerui")

	// USER AND HOST
	router.GET("/", h.ShowHouses, middlewares.ErrorHandler)
	router.GET("/users/details", middlewares.AuthorizeJWT, h.GetUserDetail, middlewares.ErrorHandler)
	router.PATCH("/users/details", middlewares.AuthorizeJWT, CORSMiddleware(), h.UpdateProfile, middlewares.ErrorHandler)
	router.POST("/signin", h.SignIn, middlewares.ErrorHandler)
	router.POST("/signup", h.SignUp, middlewares.ErrorHandler)
	router.POST("/topups", middlewares.AuthorizeJWT, h.TopUp, middlewares.ErrorHandler)
	router.GET("/houses/:id", h.ShowHouse, middlewares.ErrorHandler)
	router.GET("/transactions", middlewares.AuthorizeJWT, h.ShowTransactions, middlewares.ErrorHandler)

	router.GET("/photos/firsts/:houseId", h.GetFirstPhoto, middlewares.ErrorHandler)
	router.GET("/photos/:houseId", h.GetPhotosByHouseId, middlewares.ErrorHandler)
	router.POST("/photos/:houseId", middlewares.AuthorizeJWT, h.UploadPhotos, middlewares.ErrorHandler)

	router.GET("/houses/bookings/:bookingId", middlewares.AuthorizeJWT, h.GetBookingById, middlewares.ErrorHandler)
	router.GET("/houses/bookings", middlewares.AuthorizeJWT, h.ShowBookings, middlewares.ErrorHandler)
	router.POST("/houses/bookings/:id", middlewares.AuthorizeJWT, h.BookHouse, middlewares.ErrorHandler) // House id
	router.POST("/payment", middlewares.AuthorizeJWT, h.Pay, middlewares.ErrorHandler)
	router.GET("/pickups", middlewares.AuthorizeJWT, h.GetPickUps, middlewares.ErrorHandler)
	router.POST("/pickups", middlewares.AuthorizeJWT, h.PickUp, middlewares.ErrorHandler)
	router.POST("/games/coins", middlewares.AuthorizeJWT, h.FlipCoin, middlewares.ErrorHandler)

	// HOST ONLY
	router.POST("/houses", middlewares.AuthorizeJWT, h.AddHouse, middlewares.ErrorHandler)
	router.PATCH("/houses/:houseId", middlewares.AuthorizeJWT, h.UpdateHouse, middlewares.ErrorHandler)

	// HOST AND ADMIN ONLY
	router.DELETE("/houses/:houseId", middlewares.AuthorizeJWT, h.DeleteHouse, middlewares.ErrorHandler)

	// ADMIN ONLY
	router.PATCH("/pickups", middlewares.AuthorizeJWT, h.UpdatePickUpStatus, middlewares.ErrorHandler)

	router.POST("/signout", middlewares.AuthorizeJWT, h.SignOut, middlewares.ErrorHandler)
	router.POST("/signout/clean", middlewares.AuthorizeJWT, h.CleanBlacklists, middlewares.ErrorHandler)
	//router.POST("/tables/reset", middlewares.ErrorHandler)

	router.GET("/cities", h.GetCities, middlewares.ErrorHandler)
	return router
}

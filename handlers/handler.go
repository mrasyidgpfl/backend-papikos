package handlers

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/services"
)

type Handler struct {
	authService   services.AuthService
	houseService  services.HouseService
	walletService services.WalletService
	pickUpService services.PickUpService
	gamesService  services.GamesService
}

type HandlerConfig struct {
	AuthService   services.AuthService
	HouseService  services.HouseService
	WalletService services.WalletService
	PickUpService services.PickUpService
	GamesService  services.GamesService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		authService:   c.AuthService,
		houseService:  c.HouseService,
		walletService: c.WalletService,
		pickUpService: c.PickUpService,
		gamesService:  c.GamesService,
	}
}

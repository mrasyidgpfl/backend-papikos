package server

import (
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/config"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/db"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/repositories"
	"final-project/git.garena.com/sea-labs-id/batch-01/rasyid-wijaya/final-project-backend-papikos/services"
	"fmt"
)

func Init() {
	authRepositories := repositories.NewAuthRepository(&repositories.AuthConfig{DB: db.Get()})
	authService := services.NewAuthService(&services.AuthConfig{
		AuthRepository: authRepositories,
		AppConfig:      config.Config,
	})

	houseRepositories := repositories.NewHouseRepository(&repositories.HouseConfig{DB: db.Get()})
	houseService := services.NewHouseService(&services.HouseConfig{
		HouseRepository: houseRepositories,
		AppConfig:       config.Config,
	})

	walletRepositories := repositories.NewWalletRepository(&repositories.WalletConfig{DB: db.Get()})
	walletService := services.NewWalletService(&services.WalletConfig{
		WalletRepository: walletRepositories,
		AppConfig:        config.Config,
	})

	pickUpRepositories := repositories.NewPickUpRepository(&repositories.PickUpConfig{DB: db.Get()})
	pickUpService := services.NewPickUpService(&services.PickUpConfig{
		PickUpRepository: pickUpRepositories,
		AppConfig:        config.Config,
	})

	gamesRepositories := repositories.NewGamesRepository(&repositories.GamesConfig{DB: db.Get()})
	gamesService := services.NewGamesService(&services.GamesConfig{
		GamesRepository: gamesRepositories,
		AppConfig:       config.Config,
	})

	router := NewRouter(&RouterConfig{authService, houseService, walletService, pickUpService, gamesService})
	err := router.Run()
	if err != nil {
		fmt.Println("server error: ", err)
	}
}
